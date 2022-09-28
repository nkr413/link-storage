package main

import (
  "fmt"
	"context"
  "github.com/go-redis/redis/v8"
)

var (
	ctx = context.TODO()
)

func get(cln *redis.Client) {
	name := cln.Get(ctx, "name")
	fmt.Println(name)
}

func set(cln *redis.Client) {
	cln.Set(ctx, "user", "nkr", 0)
}

func checkConn(cln *redis.Client) {
	pong, err := cln.Ping(cln.Context()).Result()
	fmt.Println(pong, err)
}

func main() {
  fmt.Println("Testing Golang Redis")

	client := redis.NewClient(&redis.Options{
    Addr: "localhost:6379",
    Password: "",
    DB: 0,
  })

	checkConn(client)
	set(client)
	get(client)


	//fmt.Printf("%T\n", client)
}
