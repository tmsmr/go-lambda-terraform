### GENERAL ###

locals {
  # 'one-shot' invocations based on input value changes
  invoke = true
  values = ["a", "b", "c"]
}

resource "random_pet" "deployment_id" {
  length = 2
}

### ROLE/POLICY ###

data "aws_iam_policy_document" "assume_role_doc" {
  statement {
    effect = "Allow"

    principals {
      type = "Service"
      identifiers = ["lambda.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "lambda_role" {
  name               = "${random_pet.deployment_id.id}-lambda-role"
  assume_role_policy = data.aws_iam_policy_document.assume_role_doc.json
}

### FUNCTION ARCHIVE ###

data "archive_file" "fn_zip" {
  output_path = "${path.module}/dist/fn.zip"
  source_file = "${path.module}/fn/bin/bootstrap"
  type        = "zip"
}

### LAMBDA ###

resource "aws_lambda_function" "lambda_func" {
  filename      = "${path.module}/dist/fn.zip"
  function_name = "${random_pet.deployment_id.id}-lambda"
  role          = aws_iam_role.lambda_role.arn
  handler       = "unused"

  source_code_hash = data.archive_file.fn_zip.output_base64sha256

  runtime = "provided.al2023"
  architectures = ["arm64"]

  environment {
    variables = {
      FOO = "bar"
    }
  }

  depends_on = [data.archive_file.fn_zip]
}

### INVOCAIONS ###

resource "aws_lambda_invocation" "lambda_invocations" {
  count = local.invoke ? length(local.values) : 0
  function_name = aws_lambda_function.lambda_func.function_name

  input = jsonencode({
    input = local.values[count.index]
  })
}

output "invocation_results" {
  value = aws_lambda_invocation.lambda_invocations[*].result
}
