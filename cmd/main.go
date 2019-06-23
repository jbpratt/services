package main

import (
	"flag"

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
		err := client.GetChannelByName(*name)
		if err != nil {
			panic(err)
		}

	default:
	}

}
