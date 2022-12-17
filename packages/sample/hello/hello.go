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

	m := args["__ow_headers"]

	return &Response{
		Body: fmt.Sprintf("User-agent: %s\nIP: %s, args type: %T, __ow_headers type: %T", userAgent, ip, args, m)
	}, nil
}
