package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io"
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
	if endpoint == "" {
		log.Fatal("WAVE_HOST must be set")
	}
	username = os.Getenv("WAVE_USER")
	if username == "" {
		log.Fatal("WAVE_USER must be set")
	}
	password = os.Getenv("WAVE_PASSWORD")
	if password == "" {
		log.Fatal("WAVE_PASSWORD must be set")
	}
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

func logout() {
	req, err := http.NewRequest(
		"POST",
		endpoint+"/sessions/delete",
		strings.NewReader(fmt.Sprintf("")),
	)
	if err != nil {
		log.Fatal(err)
	}
	request(req)
}

func request(req *http.Request) io.ReadCloser {
	cookie := new_cookie()
	req.AddCookie(&cookie)

	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return nil
	}
	return resp.Body
}

func new_cookie() http.Cookie {
	return http.Cookie{
		Name:     "wave_session",
		Value:    session_id,
		Path:     "/",
		Domain:   "Wave-cli",
		MaxAge:   int(time.Now().AddDate(1, 0, 1).Unix()),
		Secure:   false,
		HttpOnly: true,
	}
}
