package sms

import (
	"encoding/json"
	"errors"
)

type ReceiverInterface interface {
	Receive() (string, error)
}
type Receiver struct {
	Sender      string `json:"sourceAddress"`
	Application string `json:"applicationId"`
	Message     string `json:"message"`
	Encoding    string `json:"encoding"`
	Version     string `json:"version"`
}

func (receiver *Receiver) Receive(requestString []byte) (*Receiver, error) {
	if err := json.Unmarshal(requestString, receiver); err == nil {
		return receiver, nil
	} else {
		return receiver, errors.New("unable to parse input")
	}
}
