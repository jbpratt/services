package services

import (
	"io/ioutil"
	"net/http"
	"time"
)

var serviceURLs = map[string]string{
	"twitch":     "",
	"angelthump": "https://api.angelthump.com/v1/",
}

type client interface {
	GetChannelByName(string) error
}

type service struct{}

func (s *service) Get(url string) (*http.Response, error) {
	req, err := buildRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	return s.do(req)
}

func (s *service) do(req *http.Request) (*http.Response, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func ReadAll(res *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func buildRequest(method, url string, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req, nil
}
