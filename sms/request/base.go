package request

type Request struct {
	Message       string `json:"message"`
	ApplicationId string `json:"applicationId"`
	// Encoding      string `json:"encoding"`
	// Version       string `json:"version"`
}

func (request *Request) SetMessage(message string) (*Request) {
	request.Message = message
	return request
}
func (request *Request) GetMessage() string {
	return request.Message
}
func (request *Request) SetApplicationId(appId string) {
	request.ApplicationId = appId
}
