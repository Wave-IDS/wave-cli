package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"os"
	"strings"
	"time"
)

var client http.Client
var endpoint string
var session_id string
var username string
var password string

func parse_envars() {
	endpoint = os.Getenv("WAVE_HOST")
	username = os.Getenv("WAVE_USER")
	password = os.Getenv("WAVE_PASSWORD")
}

func login() {
	parse_envars()
	client = http.Client{}

	req, err := http.NewRequest(
		"POST",
		endpoint+"/sessions/create",
		strings.NewReader(fmt.Sprintf(
			"{\"username\": \"%s\", \"password\": \"%s\"}",
			username,
			password,
		)),
	)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	if header, ok := resp.Header["Set-Cookie"]; ok {
		components := strings.Split(header[0], " ")
		cookie := components[0]
		session_id = cookie[13 : len(cookie)-1]
	} else {
		log.Fatal("unsuccessful authentication")
	}
}

func request(req *http.Request) (params map[string]string) {
	cookie := http.Cookie{
		Name:     "wave_session",
		Value:    session_id,
		Path:     "/",
		Domain:   "Wave-cli",
		MaxAge:   int(time.Now().AddDate(1, 0, 1).Unix()),
		Secure:   false,
		HttpOnly: true,
		Raw:      "wave_session=" + session_id,
		Unparsed: []string{"wave_session=" + session_id},
	}
	req.AddCookie(&cookie)

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return
	}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&params)
	if err != nil {
		log.Error(err)
		return
	}
	return
}
