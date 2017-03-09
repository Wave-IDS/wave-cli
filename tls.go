package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

func getTLS() {
	req, _ := http.NewRequest(
		"GET",
		endpoint+"/tls",
		strings.NewReader(""),
	)
	body := request(req)

	var params = make(map[string]string, 0)
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Error(err)
		return
	}

	fmt.Println(params["certificate"])
	fmt.Println(params["private_key"])
}

func setTLS(cert_file, key_file string) {
	cert_data, err := ioutil.ReadFile(cert_file)
	if err != nil {
		log.Error(err)
		return
	}
	key_data, err := ioutil.ReadFile(cert_file)
	if err != nil {
		log.Error(err)
		return
	}

	req, _ := http.NewRequest(
		"POST",
		endpoint+"/tls",
		strings.NewReader(fmt.Sprintf(
			"{\"ca_cert\": \"%s\", \"private_key\": \"%s\"}",
			string(cert_data),
			string(key_data),
		)),
	)
	body := request(req)

	var params = make(map[string]string)
	decoder := json.NewDecoder(body)
	err = decoder.Decode(&params)
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
