# How to use


This works with simulator straightforward without any other configurations.
```go

package main

import (
	"github.com/joomtriggers/ideamart"
)

func main() {
	simple := ideamart.SMS()
	defer simple.Send()
	simple.SetMessage("Message")
	simple.AddReceiver("tel:94771231232")
}


```
