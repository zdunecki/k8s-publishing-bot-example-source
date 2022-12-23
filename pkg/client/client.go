package client

type client struct{}

func NewClient() *client {
	return &client{}
}

func (c *client) HelloWorld() string {
	return "hello world"
}
