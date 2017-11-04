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

type SendRequest struct {
	Message               string   `json:"message"`
	DestinationAddresses  []string `json:"destinationAddresses"`
	Password              string   `json:"password"`
	ApplicationId         string   `json:"applicationId"`
	SourceAddress         string   `json:"sourceAddress"`
	DeliveryStatusRequest bool     `json:"deliveryStatusRequest"`
	ChargingAmount        string   `json:"chargingAmount"`
	Version               string   `json:"version"`
	BinaryHeader          string   `json:"binaryHeader"`
}
type SendResponse struct {
	StatusCode   string `json:statusCode`
	StatusDetail string `json:statusDetail`
	MessageId    string `json:messageId`
	Version      string `json:version`
}

func (request *SendRequest) SetMessage(message string) *SendRequest {
	request.Message = message
	return request
}
func (request *SendRequest) GetMessage() string {
	return request.Message;
}

func (request *SendRequest) AddReceiver(receiver string) *SendRequest {
	request.DestinationAddresses = append(request.DestinationAddresses, receiver)
	return request;
}

func (request *SendRequest) Configure(c *Configuration) *SendRequest{
	request.Password = (c.Password);
	request.ApplicationId = c.ApplicationId;
	return request
}
