package emailhandler

import (
	//"fmt"
	"encoding/json"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGoodMessageSendsSuccessfully(t *testing.T) {
	s := Null(nil)
	s.Email.OutputTrackingEnabled = true

	m := ses.SendTemplatedEmailInput{
		Source: aws.String(uuid.New().String()),
	}

	bytes, _ := json.Marshal(m)
	jsonForm := string(bytes)

	err := s.SendEmail(jsonForm)

	assert.Nil(t, err)
	assert.Equal(t, &m, s.Email.LastOutput)
}

func TestBadMessageReportsError(t *testing.T) {
	s := Null(nil)

	err := s.SendEmail(uuid.New().String())

	assert.NotNil(t, err)
}

func TestErrorFromEmailServiceReported(t *testing.T) {
	expected := errors.New(uuid.New().String())
	s := Null(expected)

	m := ses.SendTemplatedEmailInput{
		Source: aws.String(uuid.New().String()),
	}

	bytes, _ := json.Marshal(m)
	jsonForm := string(bytes)

	err := s.SendEmail(jsonForm)

	assert.Equal(t, expected, err)
}
