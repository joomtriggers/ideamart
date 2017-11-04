package ideamart;

import "github.com/joomtriggers/ideamart/sms"

func SMS() (*sms.Sender,*sms.SendRequest) {
	return sms.SMS();
}
