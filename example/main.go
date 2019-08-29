package main

import (
	"fmt"
	"log"

	"github.com/zucc-hicc/hango"
)

const webhook = "webhook url"
const thread = "thread path"

func main() {
	c := hango.NewClient(webhook)
	var bs []byte

	// Post to a new thread.
	bs, err := c.Post("Test of post to a new thead")
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%s\n", bs)

	// Post to an existing thread.
	bs, err = c.PostToThread("Test of post to an existing thread", thread)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("%s\n", bs)
}
