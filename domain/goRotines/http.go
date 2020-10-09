package goRotines

import (
	"bytes"
	"net/http"
)

type Http struct {
	url        string
	clientHttp *http.Client
}

func (h *Http) Post(payload string) {
	var jsonStr = []byte(payload)
	req, err := http.NewRequest("POST", h.url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, err := h.clientHttp.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func NewHttp(_url string) *Http {
	return &Http{
		url:        _url,
		clientHttp: &http.Client{},
	}
}
