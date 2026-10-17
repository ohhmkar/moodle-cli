package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"time"
)

type SiteInfo struct {
	Fullname string `json:"fullname"`
	UserID   int    `json:"userid"`
}

type Course struct {
	ID        int
	Fullname  string
	Shortname string
}

type Event struct {
	Name     string
	TimeSort int64
	URL      string
	Overdue  bool
	Course   Course
}

type eventsResponse struct {
	Events []Event `json:"events"`
	LastId int     `json:"lastid"`
}

const baseURL = "https://moodle.spit.ac.in/webservice/rest/server.php"

var batchRe = regexp.MustCompile(`(?:^|[-\s\d])([A-D][1-9])(?:[-WT\s]|$)`)

func call(fn string, params url.Values, out any) error {
	if params == nil {
		params = url.Values{}
	}

	moodleToken := os.Getenv("MOODLE_TOKEN")
	params.Set("wstoken", moodleToken)
	params.Set("wsfunction", fn)
	params.Set("moodlewsrestformat", "json")

	resp, err := http.PostForm(baseURL, params)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var mErr struct {
		Exception string `json:"exception"`
		Message   string `json:"message"`
	}

	// error ignored : array responses wont match
	json.Unmarshal(body, &mErr)
	if mErr.Exception != "" {
		return fmt.Errorf("moodle: %s", mErr.Message)
	}
	return json.Unmarshal(body, out)
}

func upcomingDeadlines() ([]Event, error) {
	params := url.Values{}
	// look back 14 days so overdue (unsubmitted) assignments still show up
	params.Set("timesortfrom", strconv.FormatInt(time.Now().AddDate(0, 0, -14).Unix(), 10))
	params.Set("limitnum", "50")
	var mine []Event
	for {
		var r eventsResponse
		err := call("core_calendar_get_action_events_by_timesort", params, &r)
		if err != nil {
			return nil, err
		}
		for _, e := range r.Events {
			if isMine(e.Name) {
				mine = append(mine, e)
			}
		}
		if len(r.Events) < 50 {
			break
		}
		params.Set("aftereventid", strconv.Itoa(r.LastId)) // bookmark for next page
	}
	return mine, nil
}

func isMine(name string) bool {
	m := batchRe.FindStringSubmatch(name)
	if m == nil {
		return true
	}
	return m[1] == os.Getenv("MYBATCH")
}
