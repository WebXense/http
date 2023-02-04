package http

import (
	"net/http"
)

func Get[T any](url string, headers map[string]string, params map[string]string, requestBody interface{}) (status int, body *T, err error) {
	url += encodeParams(params)
	req, err := http.NewRequest("GET", url, handleRequestBody(requestBody))
	if err != nil {
		return 0, nil, err
	}
	setHeaders(req, headers)
	return doRequest[T](req)
}

func Post[T any](url string, headers map[string]string, params map[string]string, requestBody interface{}) (status int, body *T, err error) {
	url += encodeParams(params)
	req, err := http.NewRequest("POST", url, handleRequestBody(requestBody))
	if err != nil {
		return 0, nil, err
	}
	setHeaders(req, headers)
	return doRequest[T](req)
}

func Put[T any](url string, headers map[string]string, params map[string]string, requestBody interface{}) (status int, body *T, err error) {
	url += encodeParams(params)
	req, err := http.NewRequest("PUT", url, handleRequestBody(requestBody))
	if err != nil {
		return 0, nil, err
	}
	setHeaders(req, headers)
	return doRequest[T](req)
}

func Delete[T any](url string, headers map[string]string, params map[string]string, requestBody interface{}) (status int, body *T, err error) {
	url += encodeParams(params)
	req, err := http.NewRequest("DELETE", url, handleRequestBody(requestBody))
	if err != nil {
		return 0, nil, err
	}
	setHeaders(req, headers)
	return doRequest[T](req)
}
