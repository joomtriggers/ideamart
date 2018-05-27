package ideamart

import "github.com/joomtriggers/ideamart/sms"

func SMS() *sms.Handler {
	sender := &sms.Sender{}
	receiver := &sms.Receiver{}
	handler := &sms.Handler{Sender: sender, Receiver: receiver}
	handler.SetApplication("APP_000001")
	handler.SetServer("http://127.0.0.1:7000/sms/send")
	handler.SetPassword("password")
	return handler
}
