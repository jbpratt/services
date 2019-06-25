package services

import (
	"encoding/json"
	"fmt"
)

type YoutubeClient struct {
	service
	APIKey string
}

type YoutubeAPIResponse struct {
	Username  string `json:"username"`
	Live      bool   `json:"live"`
	Title     string `json:"title"`
	Viewers   int    `json:"viewers"`
	Thumbnail string `json:"thumbnail"`
}

var _ client = (*YoutubeClient)(nil)
var _ response = (*YoutubeAPIResponse)(nil)

func (yt *YoutubeClient) GetChannelByName(id string) (response, error) {

	// swapping back to query.Set() after I get this working
	url := serviceURLs["youtube"] + "?key=" + yt.APIKey +
		"&part=liveStreamingDetails,snippet,statistics" + "&id=" + id

	fmt.Println(url)
	res, err := yt.Get(url, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(res)

	var r YoutubeAPIResponse

	dec := json.NewDecoder(res.Body)
	defer res.Body.Close()
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (r *YoutubeAPIResponse) GetLive() bool {
	return r.Live
}

func (r *YoutubeAPIResponse) GetTitle() string {
	return r.Title
}

func (r *YoutubeAPIResponse) GetViewers() int {
	return r.Viewers
}

func (r *YoutubeAPIResponse) GetThumbnail() string {
	return r.Thumbnail
}
