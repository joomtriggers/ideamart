package sms;

import "github.com/joomtriggers/ideago"

type SenderInterface interface {
	Send() *SendResponse
	SMS() (*Sender, *SendRequest)
}

type Sender struct {
	request *SendRequest
}

func SMS() (*Sender, *SendRequest) {
	sender := &Sender{};
	request := SendRequest{};
	sender.request = &request;
	return sender, sender.request;
}

func (sender *Sender) Send() *SendResponse {
	ideago.SendRequest(sender.request);
	return &SendResponse{};
}
