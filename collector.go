package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

func newCollector(name string) {
	req, _ := http.NewRequest(
		"POST",
		endpoint+"/collectors/create",
		strings.NewReader(fmt.Sprintf(
			"{\"name\": \"%s\"}",
			name,
		)),
	)
	body := request(req)

	var params = make(map[string]string)
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Error(err)
		return
	}

	cert := []byte(params["certificate"])
	err = ioutil.WriteFile(name+".crt", cert, 0644)
	if err == nil {
		fmt.Println("wrote " + name + ".crt")
	} else {
		log.Error(err)
	}
	key := []byte(params["private_key"])
	err = ioutil.WriteFile(name+".key", key, 0644)
	if err == nil {
		fmt.Println("wrote " + name + ".key")
	} else {
		log.Error(err)
	}
}

func getCollectors() {
	req, _ := http.NewRequest(
		"GET",
		endpoint+"/collectors",
		strings.NewReader(""),
	)
	body := request(req)

	var params = make([]string, 0)
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Error(err)
		return
	}

	for _, name := range params {
		fmt.Println(name)
	}
}

func deleteCollector(name string) {
	req, _ := http.NewRequest(
		"POST",
		endpoint+"/collectors/delete",
		strings.NewReader(fmt.Sprintf(
			"{\"name\": \"%s\"}",
			name,
		)),
	)
	body := request(req)

	var params = make(map[string]string)
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Error(err)
		return
	}

	if params["success"] != "" {
		fmt.Println(fmt.Sprintf("success: %s", params["success"]))
	} else {
		fmt.Println(params["error"])
	}
}
