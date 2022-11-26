package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client = redis.NewClient(&redis.Options{
	Addr: "127.0.0.1:6379",
})

func main() {

	sub := client.PSubscribe(ctx, "__keyevent@0__:expired")

	for {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(msg.Payload)
	}
}
