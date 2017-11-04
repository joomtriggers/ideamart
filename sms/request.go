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

type ConfigurationInterface interface {
	SetServer(string)
	SetApplication(string)
	SetPassword(string)
	GetPassword() string
	GetApplication() string
	GetServer() string
}
type Configuration struct {
	server        string
	applicationId string
	password      string
}

func (c *Configuration) SetServer(server string) *Configuration {
	c.server = server;
	return c
}
func (c *Configuration) SetApplication(app string) *Configuration {
	c.applicationId = app;
	return c
}
func (c *Configuration) SetPassword(pwd string) *Configuration {
	c.password = pwd;
	return c;
}
func (c *Configuration) GetServer() string {
	return c.server
}
func (c *Configuration) GetApplication() string {
	return c.applicationId
}
func (c *Configuration) GetPassword() string {
	return c.password
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

func (request *SendRequest) LoadConfiguration(c ConfigurationInterface) *SendRequest{
	request.Password = c.GetPassword();
	request.ApplicationId = c.GetApplication();
	return request
}
