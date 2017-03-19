package main

import (
	"bufio"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"golang.org/x/net/websocket"
	"net/http"
	"net/url"
	"os"
)

func streamVisualEvents() {
	endpoint_str := ""
	host := os.Getenv("WAVE_HOST")
	if len(host) >= 5 && host[:5] == "https" {
		endpoint_str = "wss" + host[5:]
	} else {
		endpoint_str = "ws" + host[4:]
	}
	ws_endpoint, err := url.Parse(endpoint_str + "/streams/visualizer")
	if err != nil {
		log.WithFields(log.Fields{
			"endpoint_str": endpoint_str,
			"error":        err.Error(),
		}).Fatal("unable to parse wave uri")
	}
	ws_origin, err := url.Parse(host + "/streams/visualizer")
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
