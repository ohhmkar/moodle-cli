package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/x/term"
)

type loginJson struct {
	Token string `json:"token"`
	Error string `json:"error"`
}

func login() error {
	fmt.Print("username:  ")
	reader := bufio.NewReader(os.Stdin)
	username, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	username = strings.TrimSpace(username)

	fmt.Println("password: ")
	pw, err := term.ReadPassword(os.Stdin.Fd())
	if err != nil {
		return err
	}
	fmt.Println()
	password := string(pw)

	params := url.Values{}
	params.Set("username", username)
	params.Set("password", password)
	params.Set("service", "moodle_mobile_app")

	resp, err := http.PostForm("https://moodle.spit.ac.in/login/token.php", params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var r loginJson

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return err
	}
	if r.Error != "" {
		return fmt.Errorf("moodle: %s", r.Error)
	}
	if r.Token == "" {
		return errors.New("no token in response")
	}

	path, err := tokenPath()
	if err != nil {
		return err
	}
	err = os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, []byte(r.Token), 0600)
	if err != nil {
		return err
	}
	fmt.Println("logged in ✓")
	return nil
}

func tokenPath() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "moodle-cli", "token"), nil
}

func loadToken() (string, error) {
	x := os.Getenv("MOODLE_TOKEN")
	if x != "" {
		return x, nil
	}
	path, err := tokenPath()
	if err != nil {
		return "", err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", errors.New("not logged in, run: moodle-cli login")
	}
	return strings.TrimSpace(string(data)), nil
}
