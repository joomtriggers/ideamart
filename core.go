package ideamart;

import (
	"bytes"
	"net/http"
	"io"
	"os"
	"encoding/json"
)


func SendRequest(jsonResponse interface{}) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(jsonResponse)
	w, _ := http.NewRequest("POST", "https://httpbin.org/post", b);
	w.Header.Add("content-type", "application/json");
	client := &http.Client{};
	res, _ := client.Do(w);
	io.Copy(os.Stdout, res.Body)
}