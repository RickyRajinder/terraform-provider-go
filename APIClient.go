package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func getIP(service string) (string, error) {
	rsp, err := http.Get(service)
	if err != nil {
		return "", err
	}

	defer rsp.Body.Close()

	if rsp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP request error. Response code: %d", rsp.StatusCode)
	}

	buf, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes.TrimSpace(buf)), err
}