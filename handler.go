package ideamart

import "github.com/joomtriggers/ideamart/sms"

func SMS() *sms.Handler {
	sender := &sms.Sender{}
	receiver := &sms.Receiver{}
	handler := &sms.Handler{Sender: sender, Receiver: receiver}
	return handler
}
