package common

import (
	"bytes"
	"net/http"
)

type ErrResponse struct {
	Message string
	Error   string
}

func MakeRequest(method, url, stringBody string, headers map[string]string) *http.Request {
	jsonStr := []byte(stringBody)
	r, _ := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))
	for k, v := range headers {
		r.Header.Set(k, v)
	}
	return r
}
