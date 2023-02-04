package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func encodeParams(params map[string]string) (encoded string) {
	if len(params) == 0 {
		return ""
	}
	encoded = "?"
	for key, value := range params {
		encoded += key + "=" + value + "&"
	}
	return encoded[:len(encoded)-1]
}

func setHeaders(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		req.Header.Set(key, value)
	}
}

func handleRequestBody(requestBody interface{}) *bytes.Buffer {
	if requestBody != nil {
		reqBody, err := json.Marshal(requestBody)
		if err != nil {
			return nil
		}
		return bytes.NewBuffer(reqBody)
	}
	return nil
}

func doRequest[T any](req *http.Request) (status int, body *T, err error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return resp.StatusCode, nil, nil
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, body, nil
}
