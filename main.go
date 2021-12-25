package main

import (
	"fmt"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to a server
	ncSub, err := nats.Connect(nats.DefaultURL,
		nats.DisconnectErrHandler(func(_ *nats.Conn, err error) {
			println("client disconnected: %v", err)
		}),
		nats.ReconnectHandler(func(_ *nats.Conn) {
			println("client reconnected")
		}),
		nats.ClosedHandler(func(_ *nats.Conn) {
			println("client closed")
		}))

	if err != nil {
		print("error in ncSub ")
		println(err.Error())
	}

	// Simple Async Subscriber
	ncSub.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	//----------------------------------------------------------------------

	nc, err := nats.Connect(nats.DefaultURL,
		nats.DisconnectErrHandler(func(_ *nats.Conn, err error) {
			println("client disconnected: %v", err)
		}),
		nats.ReconnectHandler(func(_ *nats.Conn) {
			println("client reconnected")
		}),
		nats.ClosedHandler(func(_ *nats.Conn) {
			println("client closed")
		}))

	if err != nil {
		print("error in nc ")
		println(err.Error())
	}

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World"))

	nc.Close()

	runtime.Goexit()
}
