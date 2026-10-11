package main

import (
	"fmt"
	"log"
)

func main() {
	var info SiteInfo
	err := call("core_webservice_get_site_info", nil, &info)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Hello", info.Fullname)

	events, err := upcomingDeadlines()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deadlines: ", len(events))
}
