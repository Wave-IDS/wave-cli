package main

import (
	"bufio"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/websocket"
	"net/http"
	"net/url"
)

func streamVisualEvents() {
	ws_endpoint, err := url.Parse("ws://127.0.0.1:8080/streams/visualizer") // parse or set
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("unable to parse wave uri")
	}
	ws_origin, err := url.Parse("https://127.0.0.1:8080/streams/visualizer")
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Fatal("unable to parse wave uri")
	}

	cookie := new_cookie()
	config := websocket.Config{
		Location: ws_endpoint,
		Origin:   ws_origin,
		Version:  13,
		Header: http.Header(map[string][]string{
			"Cookie": []string{cookie.String()},
		}),
	}
	ws, err := websocket.DialConfig(&config)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("failed to dial websocket")
	} else {
		scanner := bufio.NewScanner(ws)
		for {
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
			if err := scanner.Err(); err != nil {
				log.Error(err)
				return
			}
		}
	}
}
