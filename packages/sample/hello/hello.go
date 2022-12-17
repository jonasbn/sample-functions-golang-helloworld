package main

import (
	"fmt"
	"log"
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

	for _, y := range args {
		for _, z := range y.(interface{}).([]interface{}) {
			userAgent = z.(map[string]interface{})["user-agent"].(string)
			ip = z.(map[string]interface{})["do-connecting-ip"].(string)
		}
	}

	log.Printf("ip %s", ip)
	log.Printf("user-agent %s", userAgent)

	return &Response{
		Body: fmt.Sprintf("User-agent:%s\nIP:%s\n", userAgent, ip),
	}, nil
}
