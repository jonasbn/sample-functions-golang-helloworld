package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(args map[string]interface{}) (*Response, error) {

	var ip string = ""
	var userAgent string = ""

	m := args["__ow_headers"]

	val, _ := m.(map[string]interface{})

	userAgent = val["user-agent"].(string)
	ip = val["do-connecting-ip"].(string)

	log.WithFields(log.Fields{
		"ip":         ip,
		"user-agent": userAgent,
	}).Info("Example message")

	return &Response{
		Body: fmt.Sprintf("User-agent:%s\nIP:%s\n", userAgent, ip),
	}, nil
}
