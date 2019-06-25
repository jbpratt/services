package main

import (
	"flag"
	"fmt"

	"github.com/jbpratt78/services"
)

var service = flag.String("s", "", "service")
var name = flag.String("n", "", "name")

func main() {
	flag.Parse()
	if *service == "" {
		panic("must supply service")
	}

	if *name == "" {
		panic("must supply user name")
	}

	switch *service {
	case "angelthump":
		client := services.AngelThumpClient{}
		res, err := client.GetChannelByName(*name)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	case "twitch":
		client := services.TwitchClient{}
		res, err := client.GetChannelByName(*name)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	case "smashcast":
		client := services.SmashcastClient{}
		res, err := client.GetChannelByName(*name)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	case "mixer":
		client := services.MixerClient{}
		res, err := client.GetChannelByName(*name)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	case "youtube":
		client := services.YoutubeClient{}
		res, err := client.GetChannelByName(*name)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	default:
		panic("error: bad service")
	}

}
