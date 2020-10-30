package client

import "go-exercise/test/common"

type client struct {
}

func NewClient() *client {
	return &client{}
}

func (c *client) Sent(r *common.Request) {

}
