package sms;

import (
	"bytes"
	"net/http"
	"io"
	"os"
	"encoding/json"
)

func SendSMSRequest(sender *Sender) SendResponse {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(sender.SendRequest)
	w, _ := http.NewRequest("POST", sender.Server, b);
	w.Header.Add("content-type", "application/json");
	client := &http.Client{};
	res, _ := client.Do(w);
	io.Copy(os.Stdout, b)
	io.Copy(os.Stdout, res.Body)
	return sender.SendResponse;
}
