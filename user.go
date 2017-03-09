package main

import (
	"encoding/json"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"net/http"
	"strings"
)

func newUser(username string) {
	req, _ := http.NewRequest(
		"POST",
		endpoint+"/users/create",
		strings.NewReader(fmt.Sprintf(
			"{\"username\": \"%s\"}",
			username,
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

func updateUserName(name string) {
	req, _ := http.NewRequest(
		"POST",
		endpoint+"/users/name",
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

func updateUserPassword(old, new string) {
	req, _ := http.NewRequest(
		"POST",
		endpoint+"/users/password",
		strings.NewReader(fmt.Sprintf(
			"{\"old_password\": \"%s\", \"new_password\": \"%s\"}",
			old,
			new,
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

func deleteUser(username string) {
	req, _ := http.NewRequest(
		"POST",
		endpoint+"/users/delete",
		strings.NewReader(fmt.Sprintf(
			"{\"username\": \"%s\"}",
			username,
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

func assignUserPassword(username, password string) {
	req, _ := http.NewRequest(
		"POST",
		endpoint+"/users/assign-password",
		strings.NewReader(fmt.Sprintf(
			"{\"username\": \"%s\", \"password\": \"%s\"}",
			username,
			password,
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

func getUsers() {
	req, _ := http.NewRequest(
		"GET",
		endpoint+"/users",
		strings.NewReader(""),
	)
	body := request(req)

	var params = make([]map[string]string, 0)
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&params)
	if err != nil {
		log.Error(err)
		return
	}

	for _, user := range params {
		fmt.Println(fmt.Sprintf(
			"Username: %s\tName: %s\tAdmin: %s",
			user["username"],
			user["name"],
			user["admin"],
		))
	}
}
