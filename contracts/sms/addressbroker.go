package sms;

type AddressBroker interface {
    AddSubscriber(number string) AddressBroker
    RemoveSubscriber(number string) AddressBroker
    GetSubscribers() []string
}
