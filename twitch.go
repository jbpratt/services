package services

import (
	"encoding/json"
)

type TwitchClient struct {
	*service
	ClientID string
}

type TwitchResponse struct {
	Username  string `json:"username"`
	Live      bool   `json:"live"`
	Title     string `json:"title"`
	Viewers   int    `json:"viewers"`
	Thumbnail string `json:"thumbnail"`
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
var _ response = (*TwitchResponse)(nil)

func (t *TwitchClient) GetChannelByName(name string) (response, error) {

	res, err := t.Get(serviceURLs["twitch"]+name, map[string]string{
		"Client-ID": "jzkbprff40iqj646a697cyrvl0zt2m6", // t.ClientID
	})
	if err != nil {
		return nil, err
	}

	var intermResp TwitchAPIResponse

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&intermResp)
	if err != nil {
		return nil, err
	}

	r := &TwitchResponse{
		Username:  intermResp.Stream.Channel.DisplayName,
		Title:     intermResp.Stream.Channel.Status,
		Viewers:   intermResp.Stream.Viewers,
		Thumbnail: intermResp.Stream.Preview.Large,
	}

	r.Live = determineLiveStatus(&intermResp)

	return r, nil
}

func (r *TwitchResponse) GetLive() bool {
	return r.Live
}

func (r *TwitchResponse) GetTitle() string {
	return r.Title
}

func (r *TwitchResponse) GetViewers() int {
	return r.Viewers
}

func (r *TwitchResponse) GetThumbnail() string {
	return r.Thumbnail
}

func determineLiveStatus(res *TwitchAPIResponse) bool {
	if res.Stream.StreamType == "live" {
		return true
	} else {
		return false
	}
}
