package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nats-io/nats.go"
)

// Core Publish-Subscribe in Messaging
/*
https://natsbyexample.com/examples/messaging/pub-sub/go
This example demonstrates the core NATS publish-subscribe behavior. This is the fundamental pattern that all other NATS patterns and higher-level APIs build upon. There are a few takeaways from this example:

Delivery is an at-most-once. For MQTT users, this is referred to as Quality of Service (QoS) 0.
There are two circumstances when a published message won’t be delivered to a subscriber:
The subscriber does not have an active connection to the server (i.e. the client is temporarily offline for some reason)
There is a network interruption where the message is ultimately dropped
Messages are published to subjects which can be one or more concrete tokens, e.g. greet.bob. Subscribers can utilize wildcards to show interest on a set of matching subjects.
*/
func main() {

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}

	nc, _ := nats.Connect(url)

	defer nc.Drain()

	nc.Publish("greet.joe", []byte("hello"))

	sub, _ := nc.SubscribeSync("greet.*")

	msg, _ := sub.NextMsg(10 * time.Millisecond)
	fmt.Println("subscribed after a publish...")
	fmt.Printf("msg is nil? %v\n", msg == nil)

	nc.Publish("greet.joe", []byte("hello"))
	nc.Publish("greet.pam", []byte("hello"))

	msg, _ = sub.NextMsg(10 * time.Millisecond)
	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)

	msg, _ = sub.NextMsg(10 * time.Millisecond)
	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)

	nc.Publish("greet.bob", []byte("hello"))

	msg, _ = sub.NextMsg(10 * time.Millisecond)
	fmt.Printf("msg data: %q on subject %q\n", string(msg.Data), msg.Subject)
}
