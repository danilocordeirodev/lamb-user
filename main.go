package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/danilocordeirodev/lamb-user/awsgo"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, events events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error){
	awsgo.InitializeAWS()
}