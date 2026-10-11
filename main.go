package main

import (
	"fmt"
	"log"
	"time"
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
	for _, e := range events {
		due := time.Unix(e.TimeSort, 0)
		fmt.Printf("%-30s %-20s %s %s\n", e.Name, e.Course.Shortname, due.Format("Mon 02 Jan 15:04"), humanize(time.Until(due)))
	}
}
