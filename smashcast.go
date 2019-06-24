package services

import (
	"encoding/json"
	"strconv"
)

type SmashcastClient struct {
	*service
}

type SmashcastAPIResponse struct {
	Livestream []struct {
		MediaIsLive    string `json:"media_is_live"`
		MediaStatus    string `json:"media_status"`
		MediaViews     string `json:"media_views"`
		MediaThumbnail string `json:"media_thumbnail"`
	} `json:"livestream"`
}

var _ client = (*SmashcastClient)(nil)
var _ response = (*SmashcastAPIResponse)(nil)

func (s *SmashcastClient) GetChannelByName(name string) (response, error) {

	res, err := s.Get(serviceURLs["smashcast"]+name, nil)
	if err != nil {
		return nil, err
	}

	var r SmashcastAPIResponse

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (r *SmashcastAPIResponse) GetLive() bool {
	if r.Livestream[0].MediaIsLive == "1" {
		return true
	}
	return false
}

func (r *SmashcastAPIResponse) GetTitle() string {
	return r.Livestream[0].MediaStatus
}

func (r *SmashcastAPIResponse) GetViewers() int {
	v, _ := strconv.Atoi(r.Livestream[0].MediaViews)
	return v
}

func (r *SmashcastAPIResponse) GetThumbnail() string {
	return r.Livestream[0].MediaThumbnail
}
