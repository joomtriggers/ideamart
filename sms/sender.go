package sms;

import (
	"github.com/joomtriggers/ideago/sms/request"
	"github.com/joomtriggers/ideago/sms/config"
)

type SenderInterface interface {
	Send() *SendResponse
	SMS() (*Sender, *request.SendRequest)
}

type Sender struct {
	request.SendRequest
	SendResponse
	config.Configuration
}

func (sender *Sender) Send() SendResponse {
	sender.SendRequest.Configure(&sender.Configuration);
	return SendSMSRequest(sender);
}

func (sender *Sender) Configure(c *config.Configuration) *Sender {
	sender.Configuration = *c
	return sender
}
