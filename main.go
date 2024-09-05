package main

import (
    "context"
    "fmt"
    "time"

    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/rds"
)

// SnapshotInput represents the input to the Lambda function
type SnapshotInput struct {
    DBInstanceIdentifier string `json:"dbInstanceIdentifier"`
    AwsRegion string `json:"region"`
}

// Handler handles Lambda invocations
func Handler(ctx context.Context, input SnapshotInput) (string, error) {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String(input.AwsRegion),
    })
    if err != nil {
        return "", fmt.Errorf("failed to create session: %v", err)
    }

    svc := rds.New(sess)

    snapshotIdentifier := fmt.Sprintf("snapshot-%s", time.Now().Format("20060102150405"))
    fmt.Printf("db identifier: %s", *aws.String(input.DBInstanceIdentifier))
    inputParam := &rds.CreateDBSnapshotInput{
        DBInstanceIdentifier: aws.String(input.DBInstanceIdentifier),
        DBSnapshotIdentifier: aws.String(snapshotIdentifier),
    }
    _, err = svc.CreateDBSnapshotWithContext(ctx, inputParam)
    if err != nil {
        return "", fmt.Errorf("failed to create snapshot: %v", err)
    }

    return fmt.Sprintf("Snapshot %s created successfully!", snapshotIdentifier), nil
}

func main() {
    lambda.Start(Handler)
}
