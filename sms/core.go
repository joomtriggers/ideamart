package sms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func SendSMSRequest(sender *Sender) SendResponse {
	b := new(bytes.Buffer)
	//json.NewEncoder(b).Encode(sender.SendRequest)
	c, _ := json.Marshal(sender.SendRequest)
	d := string(c)
	b.WriteString(d)
	w, _ := http.NewRequest("POST", sender.Server, b)
	w.Header.Add("content-type", "application/json")
	client := &http.Client{}
	res, _ := client.Do(w)
	fmt.Printf("%+v\n", string(c))
	io.Copy(os.Stdout, res.Body)
	return sender.SendResponse
}
