package main

import (
	"fmt"
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
	params := request(req)

	fmt.Println(params["certificate"])
	fmt.Println(params["private_key"])
}
