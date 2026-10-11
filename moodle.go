package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)

type SiteInfo struct {
	Fullname string `json:"fullname"`
	UserID   int    `json:"userid"`
}

const baseURL = "https://moodle.spit.ac.in/webservice/rest/server.php"

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

	return json.NewDecoder(resp.Body).Decode(out)
}
