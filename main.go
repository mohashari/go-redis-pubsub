package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client = redis.NewClient(&redis.Options{
	Addr: "127.0.0.1:6379",
})

func main() {
	for i := 0; i < 100000; i++ {
		msg := fmt.Sprintf("hello %d", i)
		id := client.Publish(ctx, "message", msg)
		log.Println(id)
	}

}
