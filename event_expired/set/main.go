package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client = redis.NewClient(&redis.Options{
	Addr: "127.0.0.1:6379",
})

func main() {
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("msg %d", i)
		client.Set(ctx, msg, msg, time.Second*5)
	}
}
