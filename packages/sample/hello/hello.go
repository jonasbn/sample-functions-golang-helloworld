package main

import (
	"fmt"
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

	for k1, v1 := range args {
		fmt.Printf("%s is type %T\n", v1, v1)
		if k1 == "__ow_headers" {
			for k2, v2 := range v1 {
				fmt.Printf("%s is type %T\n", v2, v2)

				if k2 == "do-connecting-ip" {
					ip = v2
				}

				if k2 == "user-agent" {
					userAgent = v2
				}
			}
		}
	}

	return &Response{
		Body: fmt.Sprintf("User-agent: %s\nIP %s", userAgent, ip),
	}, nil
}
