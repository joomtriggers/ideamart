package ideago;

import "github.com/joomtriggers/ideago/sms"

func SMS() (*sms.Sender,*sms.SendRequest) {
	return sms.SMS();
}
