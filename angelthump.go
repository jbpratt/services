package services

import "fmt"

type AngelThumpClient struct {
	*service
}

type atResponse struct {
	Username  string `json:"username"`
	Live      bool   `json:"live"`
	Title     string `json:"title"`
	Viewers   int    `json:"viewers"`
	Thumbnail string `json:"thumbnail"`
}

var _ client = (*AngelThumpClient)(nil)

func (at *AngelThumpClient) GetChannelByName(name string) error {
	res, err := at.Get(serviceURLs["angelthump"] + name)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
