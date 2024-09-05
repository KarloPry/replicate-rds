#!/bin/sh
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go
zip myFunction.zip bootstrap
clear
read -p "Function name for this lambda " functionName
read -p "Does this lambda already exists? (y/n) " alreadyExists
# Create a new version of the lambda function

if [ $alreadyExists = "y" ]; then
    echo "Updating lambda function"
    aws lambda update-function-code --function-name $functionName --zip-file fileb://myFunction.zip
else
    echo "Creating a new lambda function"
    aws lambda create-function --function-name $functionName \
    --runtime provided.al2023 --handler bootstrap \
    --role arn:aws:iam::438785512376:role/lambda-create-replica \
    --zip-file fileb://myFunction.zip
fi
