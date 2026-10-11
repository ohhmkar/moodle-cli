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
}

const baseURL = "https://moodle.spit.ac.in/webservice/rest/server.php"

var batchRe = regexp.MustCompile(`(?:^|-)([A-D][1-4])(?:-|W)`)

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
	params.Set("timesortfrom", strconv.FormatInt(time.Now().Unix(), 10))
	params.Set("limitnum", "50")

	var r eventsResponse
	err := call("core_calendar_get_action_events_by_timesort", params, &r)
	if err != nil {
		return nil, err
	}
	var mine []Event
	for _, e := range r.Events {
		if isMine(e.Name) {
			mine = append(mine, e)
		}
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
