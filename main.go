package main

import (
	"fmt"
	client "github.com/zdunecki/k8s-publishing-bot-example-source-client"
)

func main() {
	c := client.NewClient()

	fmt.Print(c.HelloWorld())
}
