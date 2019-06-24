package services

import (
	"encoding/json"
)

type TwitchClient struct {
	*service
	ClientID string
}

type TwitchAPIResponse struct {
	Stream struct {
		Viewers    int    `json:"viewers"`
		StreamType string `json:"stream_type"`
		Preview    struct {
			Large string `json:"large"`
		} `json:"preview"`
		Channel struct {
			Status      string `json:"status"`
			DisplayName string `json:"display_name"`
		} `json:"channel"`
	} `json:"stream"`
}

var _ client = (*TwitchClient)(nil)
var _ response = (*TwitchAPIResponse)(nil)

func (t *TwitchClient) GetChannelByName(name string) (response, error) {

	res, err := t.Get(serviceURLs["twitch"]+name, map[string]string{
		"Client-ID": "jzkbprff40iqj646a697cyrvl0zt2m6", // t.ClientID
	})
	if err != nil {
		return nil, err
	}

	var r TwitchAPIResponse

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (r *TwitchAPIResponse) GetLive() bool {
	if r.Stream.StreamType == "live" {
		return true
	}
	return false
}

func (r *TwitchAPIResponse) GetTitle() string {
	return r.Stream.Channel.DisplayName
}

func (r *TwitchAPIResponse) GetViewers() int {
	return r.Stream.Viewers
}

func (r *TwitchAPIResponse) GetThumbnail() string {
	return r.Stream.Preview.Large
}
