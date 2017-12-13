# How to use


```

package main

import (
	ideamart "github.com/joomtriggers/ideago"
)

func main() {

	sms := ideamart.SMS();
	defer sms.Send();
	sms.SetApplication("APP_000001").SetPassword("password").SetServer("http://127.0.0.1:7000/sms/send")
	sms.SetMessage("Message").AddReceiver("tel:94771231232");
}


```
