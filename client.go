package services

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var serviceURLs = map[string]string{
	"angelthump": "https://api.angelthump.com/v1/",
	"mixer":      "https://mixer.com/api/v1/channels/",
	"smashcast":  "https://api.smashcast.tv/media/live/",
	"twitch":     "https://api.twitch.tv/kraken/streams/",
	"youtube":    "https://www.googleapis.com/youtube/v3/videos",
}

type client interface {
	GetChannelByName(string) (response, error)
}

type response interface {
	GetLive() bool
	GetTitle() string
	GetThumbnail() string
	GetViewers() int
}

type service struct {
}

func (s *service) Get(url string, headers map[string]string) (*http.Response, error) {
	req, err := buildRequest("GET", url, headers)
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

	if res.StatusCode == 400 {
		return nil, errors.New("400: bad request")
	}

	if res.StatusCode == 401 {
		return nil, errors.New("401: access denied")
	}

	if res.StatusCode == 403 {
		return nil, errors.New("403: forbidden")
	}

	if res.StatusCode == 404 {
		return nil, errors.New("404: not found")
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
