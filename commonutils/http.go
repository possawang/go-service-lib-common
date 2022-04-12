package commonutils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func RestConsuming[V any, P any](method string, url string, headers map[string]string, body P) (V, error) {
	var response V
	client := &http.Client{}
	jsonReq, err := json.Marshal(body)
	if err != nil {
		return response, err
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		return response, err
	}
	for key, val := range headers {
		req.Header.Add(key, val)
	}
	resp, err := client.Do(req)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}
	err = json.Unmarshal(bodyBytes, &response)
	return response, err
}
