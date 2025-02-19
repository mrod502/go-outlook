package client

import (
	"net/http"

	"github.com/mrod502/go-outlook/types/message"
)

type Client struct {
	c *http.Client
	u string
}

func (c *Client) Message(id string, userId ...string) (*message.Message, error) {
	return nil, nil
}
