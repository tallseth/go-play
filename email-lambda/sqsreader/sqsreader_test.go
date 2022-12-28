package sqsreader

import (
	"fmt"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestReportsErrorsButRunsEverything(t *testing.T) {
	fail := true
	callCount := 0
	failEveryOtherTimeHandler := func(body string) error {
		callCount++
		fail = !fail
		if fail {
			return fmt.Errorf("failed")
		}
		return nil
	}
	lambda := SQSReader(failEveryOtherTimeHandler)
	incomingEvent := events.SQSEvent{Records: []events.SQSMessage{randomMessage(), randomMessage(), randomMessage()}}

	batchResponse, err := lambda(nil, incomingEvent)

	assert.NotNil(t, batchResponse.BatchItemFailures)
	assert.Nil(t, err)
	assert.Equal(t, len(batchResponse.BatchItemFailures), 1)

	assert.Contains(t, batchResponse.BatchItemFailures[0].ItemIdentifier, incomingEvent.Records[1].MessageId)
}

func TestNoErrors(t *testing.T) {
	alwaysSucceed := func(body string) error {
		return nil
	}
	lambda := SQSReader(alwaysSucceed)
	incomingEvent := events.SQSEvent{Records: []events.SQSMessage{randomMessage(), randomMessage(), randomMessage()}}

	batchResponse, err := lambda(nil, incomingEvent)

	assert.Nil(t, err)
	assert.NotNil(t, batchResponse.BatchItemFailures)
	assert.Equal(t, len(batchResponse.BatchItemFailures), 0)
}

func TestPassesMessageBody(t *testing.T) {
	message := randomMessage()
	incomingEvent := events.SQSEvent{Records: []events.SQSMessage{message}}
	assertCorrectBodyPassed := func(body string) error {
		if body != message.Body {
			panic(fmt.Errorf("bad message"))
		}
		return nil
	}
	lambda := SQSReader(assertCorrectBodyPassed)

	batchResponse, err := lambda(nil, incomingEvent)

	assert.Nil(t, err)
	assert.NotNil(t, batchResponse.BatchItemFailures)
	assert.Equal(t, len(batchResponse.BatchItemFailures), 0)
}

func randomMessage() events.SQSMessage {
	return events.SQSMessage{
		MessageId: uuid.New().String(),
		Body:      uuid.New().String(),
	}
}

func getFailureIDs(batchResponse events.SQSEventResponse) []string {
	failureIDs := []string{}
	for _, batchItemFailure := range batchResponse.BatchItemFailures {
		failureIDs = append(failureIDs, batchItemFailure.ItemIdentifier)
	}
	return failureIDs
}
