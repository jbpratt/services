package services

import (
	"encoding/json"
	"net/url"
	"strconv"
)

type YoutubeClient struct {
	service
	APIKey string
}

type YoutubeAPIResponse struct {
	Items []struct {
		Snippet struct {
			Title      string `json:"title"`
			Thumbnails struct {
				Medium struct {
					URL string `json:"url"`
				} `json:"medium"`
			} `json:"thumbnails"`
		} `json:"snippet"`
		LiveStreamingDetails struct {
			ConcurrentViewers string `json:"concurrentViewers"`
		} `json:"liveStreamingDetails"`
		Statistics struct {
			ViewCount string `json:"viewCount"`
		} `json:"statistics"`
	} `json:"items"`
}

var _ client = (*YoutubeClient)(nil)
var _ response = (*YoutubeAPIResponse)(nil)

func (yt *YoutubeClient) GetChannelByName(id string) (response, error) {

	u, err := url.Parse(serviceURLs["youtube"])
	if err != nil {
		return nil, err
	}

	// params func probably
	q := u.Query()
	q.Set("key", yt.APIKey)
	q.Set("part", "liveStreamingDetails,snippet,statistics")
	q.Set("id", id)
	u.RawQuery = q.Encode()

	res, err := yt.Get(u.String(), nil)
	if err != nil {
		return nil, err
	}

	var r YoutubeAPIResponse

	dec := json.NewDecoder(res.Body)
	defer res.Body.Close()
	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

// want a better way to determine this
func (r *YoutubeAPIResponse) GetLive() bool {
	if r.Items[0].LiveStreamingDetails.ConcurrentViewers != "0" {
		return true
	}
	return false
}

func (r *YoutubeAPIResponse) GetTitle() string {
	return r.Items[0].Snippet.Title
}

// need way to take into account whether LiveStreamingDetails
// has been populated or not
func (r *YoutubeAPIResponse) GetViewers() int {
	res, _ := strconv.Atoi(r.Items[0].Statistics.ViewCount)
	return res
}

func (r *YoutubeAPIResponse) GetThumbnail() string {
	return r.Items[0].Snippet.Thumbnails.Medium.URL
}
