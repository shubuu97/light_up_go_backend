package middleware

import "github.com/micro/go-micro/client"

type ClientWrapper struct {
	client.Client
}

func NewClientWrapper(c client.Client) ClientWrapper {
	if c == nil {
		c = client.NewClient()
	}
	return ClientWrapper{
		Client: c,
	}
}
