package request

import "github.com/joomtriggers/ideamart/config"

type SendRequest struct {
	Request
	DestinationAddresses []string `json:"destinationAddresses"`
	SourceAddress        string   `json:"sourceAddress"`
	Password             string   `json:"password"`
	//DeliveryStatusRequest bool     `json:"deliveryStatusRequest"`
	//ChargingAmount        string   `json:"chargingAmount"`
	//BinaryHeader          string   `json:"binaryHeader"`
}

type SendRequestInterface interface {
	SetMessage(string) SendRequestInterface
	GetMessage() string
	AddReceiver(string) SendRequestInterface
	//AddReceivers([]string) SendRequestInterface
	//setPassword(string) SendRequestInterface
	setApplication(string) SendRequestInterface
	setSourceAddress(string) SendRequestInterface
	//setVersion(string) SendRequestInterface
	//Send()
}

func (request *SendRequest) AddReceiver(receiver string) *SendRequest {
	request.DestinationAddresses = append(request.DestinationAddresses, receiver)
	return request
}
func (request *SendRequest) SetSourceAddress(sourceAddress string) *SendRequest {
	request.SourceAddress = sourceAddress
	return request
}
func (request *SendRequest) SetApplicationId(appId string) *SendRequest {
	request.ApplicationId = appId
	return request
}

func (request *SendRequest) Configure(c *config.Configuration) *SendRequest {
	request.Password = (c.Password)
	request.ApplicationId = c.ApplicationId
	return request
}
