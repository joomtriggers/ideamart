package sms;

import "github.com/joomtriggers/ideamart"

type SenderInterface interface {
	Send() *SendResponse
	SMS() (*Sender,*SendRequest)
}

type Sender struct {
	request *SendRequest
}

func SMS() (*Sender,*SendRequest) {
	sender := &Sender{};
	request := SendRequest{};
	sender.request = &request;
	return sender,sender.request;
}

func (sender *Sender) Send() *SendResponse {
	ideamart.SendRequest(sender.request);
	return &SendResponse{};
}
