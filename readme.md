# How to use


```

	smss := ideago.SMS();
	defer smss.Send();
	smss.SetApplication("APP_000001").SetPassword("password").SetServer("http://127.0.0.1:7000/sms/send")
	smss.SetMessage("Message").AddReceiver("tel:94771231232");


```
