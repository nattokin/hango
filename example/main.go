package main

import (
	"fmt"

	"github.com/zucc-hicc/gcposter"
)

const webhook = "webhook url"
const thread = "thread path"

func main() {
	c := gcposter.NewClient(webhook)
	var bs []byte

	// Post to a new thread.
	bs, err := c.Post("Test of post to a new thead")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", bs)

	// Post to an existing thread.
	bs, err = c.PostToThread("Test of post to an existing thread", thread)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%s\n", bs)
}
