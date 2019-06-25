package services

import (
	"encoding/json"
)

type MixerClient struct {
	service
}

type MixerAPIResponse struct {
	Online         bool   `json:"online"`
	ViewersCurrent int    `json:"viewersCurrent"`
	Name           string `json:"name"`
	Thumbnail      struct {
		URL string `json:"url"`
	} `json:"thumbnail"`
}

var _ client = (*MixerClient)(nil)
var _ response = (*MixerAPIResponse)(nil)

func (at *MixerClient) GetChannelByName(name string) (response, error) {

	res, err := at.Get(serviceURLs["mixer"]+name, nil)
	if err != nil {
		return nil, err
	}

	var r MixerAPIResponse

	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (r *MixerAPIResponse) GetLive() bool {
	return r.Online
}

func (r *MixerAPIResponse) GetTitle() string {
	return r.Name
}

func (r *MixerAPIResponse) GetViewers() int {
	return r.ViewersCurrent
}

func (r *MixerAPIResponse) GetThumbnail() string {
	return r.Thumbnail.URL
}
