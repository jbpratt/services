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
		fmt.Println(res)
		if err != nil {
			panic(err)
		}
	case "twitch":
		client := services.TwitchClient{}
		res, err := client.GetChannelByName(*name)
		fmt.Println(res)
		if err != nil {
			panic(err)
		}
	case "smashcast":
		client := services.SmashcastClient{}
		res, err := client.GetChannelByName(*name)
		fmt.Println(res)
		if err != nil {
			panic(err)
		}
	default:
		panic("error: bad service")
	}

}
