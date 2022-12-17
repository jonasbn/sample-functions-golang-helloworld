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

	//m := args["__ow_headers"]

	for k, v := range args {
		switch c := v.(type) {
		case string:
			fmt.Printf("Item %q is a string, containing %q\n", k, c)
		case float64:
			fmt.Printf("Looks like item %q is a number, specifically %f\n", k, c)
		default:
			fmt.Printf("Not sure what type item %q is, but I think it might be %T\n", k, c)
		}
	}

	return &Response{
		Body: fmt.Sprintf("User-agent: %s\nIP: %s", userAgent, ip),
	}, nil
}
