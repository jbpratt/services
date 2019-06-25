package services

import "encoding/json"

type AngelThumpClient struct {
	service
}

type AngelThumpAPIResponse struct {
	Live      bool   `json:"live"`
	Title     string `json:"title"`
	Viewers   int    `json:"viewers"`
	Thumbnail string `json:"thumbnail"`
}

var _ client = (*AngelThumpClient)(nil)
var _ response = (*AngelThumpAPIResponse)(nil)

func (at *AngelThumpClient) GetChannelByName(name string) (response, error) {

	res, err := at.Get(serviceURLs["angelthump"]+name, nil)
	if err != nil {
		return nil, err
	}

	var r AngelThumpAPIResponse

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (r *AngelThumpAPIResponse) GetLive() bool {
	return r.Live
}

func (r *AngelThumpAPIResponse) GetTitle() string {
	return r.Title
}

func (r *AngelThumpAPIResponse) GetViewers() int {
	return r.Viewers
}

func (r *AngelThumpAPIResponse) GetThumbnail() string {
	return r.Thumbnail
}
