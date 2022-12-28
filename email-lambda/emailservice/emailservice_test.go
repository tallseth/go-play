package emailservice

import (
	//"fmt"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSendsToSES(t *testing.T) {
	s := Null(nil)
	s.OutputTrackingEnabled = true

	m := ses.SendTemplatedEmailInput{
		Source: aws.String(uuid.New().String()),
	}

	err := s.Send(&m)

	assert.Nil(t, err)
	assert.Equal(t, &m, s.LastOutput)
}

func TestDoesNotTrackOutputIfNotEnabled(t *testing.T) {
	s := Null(nil)
	s.OutputTrackingEnabled = false

	m := ses.SendTemplatedEmailInput{
		Source: aws.String(uuid.New().String()),
	}

	err := s.Send(&m)

	assert.Nil(t, err)
	assert.Nil(t, s.LastOutput)
}

func TestReturnsErrorIfPassedIn(t *testing.T) {
	expected := errors.New(uuid.New().String())
	s := Null(expected)

	m := ses.SendTemplatedEmailInput{
		Source: aws.String(uuid.New().String()),
	}

	err := s.Send(&m)

	assert.Equal(t, expected, err)
}
