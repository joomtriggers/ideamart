# How to use


```


	sender,request := ideago.SMS();
	defer sender.Send();
	request.SetMessage("Message").AddReceiver("TEsting");

```
