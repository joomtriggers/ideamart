package sms;

import "github.com/joomtriggers/ideago"

type SenderInterface interface {
	Send() *SendResponse
	SMS() (*Sender, *SendRequest)
}

type Sender struct {
	request       *SendRequest
	configuration *Configuration
}

func SMS() (*Sender, *SendRequest) {
	sender := &Sender{};
	request := SendRequest{};
	sender.request = &request;
	return sender, sender.request;
}
func (sender *Sender) LoadConfiguration(c *Configuration) *Sender {
	sender.configuration = c;
	return sender;
}

func (sender *Sender) Send() *SendResponse {
	ideago.SendRequest(sender.request, sender.configuration);
	return &SendResponse{};
}
