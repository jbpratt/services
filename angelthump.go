package services

import "encoding/json"

type AngelThumpClient struct {
	service
}

type AngelThumpResponse struct {
	Username  string `json:"username"`
	Live      bool   `json:"live"`
	Title     string `json:"title"`
	Viewers   int    `json:"viewers"`
	Thumbnail string `json:"thumbnail"`
}

var _ client = (*AngelThumpClient)(nil)
var _ response = (*AngelThumpResponse)(nil)

func (at *AngelThumpClient) GetChannelByName(name string) (response, error) {
	res, err := at.Get(serviceURLs["angelthump"]+name, nil)
	if err != nil {
		return nil, err
	}
	var r AngelThumpResponse
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (r *AngelThumpResponse) GetLive() bool {
	return r.Live
}

func (r *AngelThumpResponse) GetTitle() string {
	return r.Title
}

func (r *AngelThumpResponse) GetViewers() int {
	return r.Viewers
}

func (r *AngelThumpResponse) GetThumbnail() string {
	return r.Thumbnail
}
