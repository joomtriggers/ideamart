package sms

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func sendRequest(sender *Sender) SendResponse {
	b := new(bytes.Buffer)
	c, _ := json.Marshal(sender.SendRequest)
	d := string(c)
	b.WriteString(d)
	w, _ := http.NewRequest("POST", sender.Server, b)
	w.Header.Add("content-type", "application/json")
	client := &http.Client{}
	res, _ := client.Do(w)
	io.Copy(os.Stdout, res.Body)
	return *sender.SendResponse
}
