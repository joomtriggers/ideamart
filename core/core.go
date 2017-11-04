package core;

import (
	"bytes"
	"net/http"
	"io"
	"os"
	"encoding/json"
)

type Request struct {
	Name string
}

func sendRequest(jsonResponse Request, url string) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(jsonResponse)
	w, _ := http.NewRequest("POST", "https://httpbin.org/post", b);
	w.Header.Add("content-type", "application/json");
	client := &http.Client{};
	res, _ := client.Do(w);
	io.Copy(os.Stdout, res.Body)
	//io.Copy(os.Stdout, p)

}

func main() {
	ab := Request{Name: "TEsting"};
	sendRequest(ab, "https://localhost:8080/");
}
