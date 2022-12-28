package emailservice

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type sesWrapper interface {
	SendTemplatedEmail(input *ses.SendTemplatedEmailInput) (*ses.SendTemplatedEmailOutput, error)
}

type EmailService struct {
	ses                   sesWrapper
	OutputTrackingEnabled bool
	LastOutput            *ses.SendTemplatedEmailInput
}

func (e *EmailService) Send(input *ses.SendTemplatedEmailInput) error {
	if e.OutputTrackingEnabled {
		e.LastOutput = input
	}

	_, err := e.ses.SendTemplatedEmail(input)
	return err
}

func new(core sesWrapper) EmailService {
	return EmailService{ses: core}
}

func New() EmailService {
	cfg := aws.NewConfig().WithRegion("us-east-1")
	return new(ses.New(session.Must(session.NewSession(cfg))))
}

type nullableEmailService struct {
	err error
}

func (n *nullableEmailService) SendTemplatedEmail(input *ses.SendTemplatedEmailInput) (*ses.SendTemplatedEmailOutput, error) {
	return nil, n.err
}

func Null(err error) EmailService {
	return new(&nullableEmailService{err: err})
}
