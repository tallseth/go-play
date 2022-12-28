package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/tallseth/go-play/email-lambda/emailhandler"
	"github.com/tallseth/go-play/email-lambda/sqsreader"
)

func main() {
	handler := emailhandler.New()
	lambda.Start(sqsreader.SQSReader(handler.SendEmail))
}
