package main

import (
	b64 "encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://localhost:8888/videos"

	auth := b64.StdEncoding.EncodeToString([]byte("user:pass"))
	fmt.Println(auth)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("authorization", "Basic "+auth)
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(res)
	fmt.Println(string(body))
}
