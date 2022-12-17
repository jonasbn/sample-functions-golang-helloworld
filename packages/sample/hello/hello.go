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

	for _, y := range args {
		for _, z := range y.(interface{}).([]interface{}) {
			userAgent = z.(map[string]interface{})["user-agent"].(string)
			ip = z.(map[string]interface{})["do-connecting-ip"].(string)
		}
	}

	return &Response{
		Body: fmt.Sprintf("User-agent:%s\nIP:%s\nargs type:%T\now_headers type: %T\n", userAgent, ip, args, m)}, nil
}
