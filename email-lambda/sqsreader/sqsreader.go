package sqsreader

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

// SQSReader... An adapter for reading SQS messages from a queue. This abstracts the error handling and retry mechanism for failed messages
func SQSReader(handler func(string) error) func(context.Context, events.SQSEvent) (events.SQSEventResponse, error) {

	return func(ctx context.Context, event events.SQSEvent) (events.SQSEventResponse, error) {

		result := events.SQSEventResponse{BatchItemFailures: make([]events.SQSBatchItemFailure, 0)}

		for i := 0; i < len(event.Records); i++ {
			record := event.Records[i]
			err := handler(record.Body)
			if err != nil {
				addFailure(&result, record)
			}
		}

		return result, nil
	}
}

func addFailure(response *events.SQSEventResponse, failingMessage events.SQSMessage) {
	response.BatchItemFailures = append(response.BatchItemFailures, events.SQSBatchItemFailure{ItemIdentifier: failingMessage.MessageId})
}
