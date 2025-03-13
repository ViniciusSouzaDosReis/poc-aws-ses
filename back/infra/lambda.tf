resource "aws_lambda_function" "email_lambda" {
  function_name = "email-lambda-function"
  role          = aws_iam_role.lambda_exec.arn
  handler       = "bootstrap"

  filename         = "../bin/api.zip"
  source_code_hash = filebase64sha256("../bin/api.zip")

  environment {
    variables = {
      SENDER_EMAIL = var.sender_email
    }
  }

  runtime       = "provided.al2"
  architectures = ["x86_64"]
}

resource "aws_lambda_permission" "apigw" {
  statement_id  = "AllowAPIGatewayInvoke"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.email_lambda.function_name
  principal     = "apigateway.amazonaws.com"

  source_arn = "${aws_api_gateway_rest_api.example.execution_arn}/*/*"
}
