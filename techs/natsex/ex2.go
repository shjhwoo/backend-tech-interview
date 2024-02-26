package natsex

/*
The request-reply pattern is built on top of the core publish-subscribe model.
the client could indicate to NATS that multiple replies should be allowed.

This example shows the basics of the request-reply pattern including the standard “no responders” error if there are no subscribers available to handle and reply to the requesting message.
*/

package main


import (
	"fmt"
	"os"
	"time"


	"github.com/nats-io/nats.go"
)


func main() {

	url := os.Getenv("NATS_URL")
	if url == "" {
		url = nats.DefaultURL
	}


	nc, _ := nats.Connect(url)
	defer nc.Drain()


	sub, _ := nc.Subscribe("greet.*", func(msg *nats.Msg) {
		/*요청에 대해 응답해주는 기능을 추가할수도 있다.
		이러한 작업을 서비스라고 부른다.*/
		name := msg.Subject[6:]
		msg.Respond([]byte("hello, " + name))
	})


	rep, _ := nc.Request("greet.joe", nil, time.Second)
	fmt.Println(string(rep.Data))


	rep, _ = nc.Request("greet.sue", nil, time.Second)
	fmt.Println(string(rep.Data))


	rep, _ = nc.Request("greet.bob", nil, time.Second)
	fmt.Println(string(rep.Data))


	sub.Unsubscribe()


	_, err := nc.Request("greet.joe", nil, time.Second)
	fmt.Println(err)
}