package request



type ReceiveRequest struct {
	Request
	SourceAddress string `json:"sourceAddress"`
	RequestId     string `json:"requestId"`
}
