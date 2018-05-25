package sms

type SendRequestInterface interface {
	SetMessage(string) SendRequestInterface
	GetMessage() string
	AddReceiver(string) SendRequestInterface
	//AddReceivers([]string) SendRequestInterface
	//setPassword(string) SendRequestInterface
	//setApplication(string) SendRequestInterface
	//setSourceAddress(string) SendRequestInterface
	//setVersion(string) SendRequestInterface
	//Send()
}
type SendResponseInterface interface {
	GetStatusCode() string
	GetStatusDetail() string
	GetMessageId() string
	GetVersion() string
}

type SendResponse struct {
	StatusCode   string `json:statusCode`
	StatusDetail string `json:statusDetail`
	MessageId    string `json:messageId`
	Version      string `json:version`
}
