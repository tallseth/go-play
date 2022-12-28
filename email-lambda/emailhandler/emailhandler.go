package emailhandler

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/tallseth/go-play/email-lambda/emailservice"
)

type EmailHandler struct {
	//todo: plumb output tracking thru and make this private
	Email emailservice.EmailService
}

func New() EmailHandler {
	return EmailHandler{Email: emailservice.New()}
}

func Null(err error) EmailHandler {
	return EmailHandler{Email: emailservice.Null(err)}
}

func (h *EmailHandler) SendEmail(serializedMessage string) error {
	toSend := ses.SendTemplatedEmailInput{}

	err := json.Unmarshal([]byte(serializedMessage), &toSend)
	if err != nil {
		return err
	}

	err = h.Email.Send(&toSend)
	if err != nil {
		return err
	}

	return nil
}
