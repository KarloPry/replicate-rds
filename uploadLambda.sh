#!/bin/sh
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go
zip myFunction.zip bootstrap
aws lambda create-function --function-name myFunction \
--runtime provided.al2023 --handler bootstrap \
--role arn:aws:iam::438785512376:role/lambda-create-replica \
--zip-file fileb://myFunction.zip
