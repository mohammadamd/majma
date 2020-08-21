package resource

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type HttpClient struct {
	Translator Translator
	client     http.Client
	Method     string
	URL        string
	Key        string
}

type HttpRequest struct {
	Body    interface{}
	Headers map[string]string
}

func (hc HttpClient) initialize() error {
	transport := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}

	hc.client = http.Client{
		Transport: transport,
		Timeout:   1 * time.Second,
	}

	return nil
}

func (hc HttpClient) GetData(cache Cache, Request interface{}) (interface{}, error) {
	req := Request.(HttpRequest)

	body, err := json.Marshal(req.Body)
	if err != nil {
		return nil, err
	}

	hReq, err := http.NewRequest(hc.Method, hc.URL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	hReq.Header.Add("Content-Type", " application/json")
	for key, value := range req.Headers {
		hReq.Header.Add(key, value)
	}

	res, err := hc.client.Do(hReq)
	if err != nil {
		return nil, err
	}

	err = res.Body.Close()
	if err != nil {

		return nil, err
	}

	return ioutil.ReadAll(res.Body)
}

func (hc HttpClient) GetTranslator() Translator {
	return hc.Translator
}

func (hc HttpClient) GetKey() string {
	return hc.Key
}
