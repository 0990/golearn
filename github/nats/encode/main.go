package main

import (
	"fmt"
	nats "github.com/nats-io/nats.go"
	"time"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()

	// Simple Publisher
	c.Publish("foo", "Hello World")

	// Simple Async Subscriber
	c.Subscribe("foo", func(s string) {
		fmt.Printf("Received a message: %s\n", s)
	})

	// EncodedConn can Publish any raw Go type using the registered Encoder
	type person struct {
		Name    string
		Address string
		Age     int
	}

	// Go type Subscriber
	c.Subscribe("hello", func(p *person) {
		fmt.Printf("Received a person: %+v\n", p)
	})

	me := &person{Name: "derek", Age: 22, Address: "140 New Montgomery Street, San Francisco, CA"}

	// Go type Publisher
	c.Publish("hello", me)

	// Unsubscribe
	sub, err := c.Subscribe("foo", nil)
	sub.Unsubscribe()

	// Requests
	var response string
	err = c.Request("help", "help me", &response, 10*time.Millisecond)
	if err != nil {
		fmt.Printf("Request failed: %v\n", err)
	}
}
