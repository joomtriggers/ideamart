# How to use


This works with simulator straightforward without any other configurations.
```go

package main

import (
	ideamart "github.com/joomtriggers/ideamart"
)

func main() {

	simple := ideamart.SMS()
	defer simple.Send()
	simple.SetApplication("APP_000001")
	simple.SetServer("http://127.0.0.1:7000/sms/send")
	simple.SetPassword("password")
	simple.SetMessage("Message")
	simple.AddReceiver("tel:94771231232")
}


```
