package sms;

type SenderInterface interface {
	Send() *SendResponse
	SMS() (*Sender, *SendRequest)
}

type Sender struct {
	SendRequest
	SendResponse
	Configuration
}

func (sender *Sender) Send() SendResponse {
	sender.SendRequest.Configure(&sender.Configuration);
	return SendSMSRequest(sender);
}

func (sender *Sender) Configure(c *Configuration) *Sender {
	sender.Configuration = *c
	return sender
}
