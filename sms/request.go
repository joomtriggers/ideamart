package sms

type SendRequestInterface interface {
	SetMessage(string) SendRequestInterface
	GetMessage() string
	AddReceiver(string) SendRequestInterface
}

type SendResponseInterface interface {
	GetStatusCode() string
	GetStatusDetail() string
	GetMessageId() string
	GetVersion() string
}

type SendResponse struct {
	StatusCode   string `json:"statusCode"`
	StatusDetail string `json:"statusDetail"`
	MessageId    string `json:"messageId"`
	Version      string `json:"version"`
}
