package main

import (
	"fmt"
	"github.com/zdunecki/k8s-publishing-bot-example-source/pkg/client"
)

func main() {
	c := client.NewClient()

	fmt.Print(c.HelloWorld())
}
