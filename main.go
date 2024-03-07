package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/danilocordeirodev/lamb-user/awsgo"
	"github.com/danilocordeirodev/lamb-user/db"
	"github.com/danilocordeirodev/lamb-user/models"
)

func main() {
	lambda.Start(ExecuteLambda)
}

func ExecuteLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error){
	awsgo.InitializeAWS()

	if !ValidateParameters() {
		fmt.Println("Error in parameters. Should send 'SecretManager'")
		err := errors.New("Error in parameters. Should send 'SecretManager'")
		return event, err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			data.UserEmail = att
			fmt.Println("Email = " + data.UserEmail)
		case "sub":
			data.UserUUID = att
			fmt.Println("Sub = " + data.UserUUID)
		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error to read secret " + err.Error())
		return event, err
	}

	err = db.SignUp(data)
	return event, err
}

func ValidateParameters() bool {
	var loadParameter bool
	_, loadParameter = os.LookupEnv("SecretName")
	return loadParameter
}