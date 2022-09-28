package main

import (
  "fmt"
	"context"
  "github.com/go-redis/redis/v8"
)

var (
	ctx = context.TODO()
)

func get(cln *redis.Client) error {
	x, err := cln.Get(ctx, "user1").Result()

	if err == redis.Nil {
		fmt.Println("no value found")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println(x)
	}

	return nil
}

func set(cln *redis.Client) error {
	err := cln.Set(ctx, "user1", "lone", 0).Err()

	if err != nil {
		return err
	}

	return nil
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
	//set(client)
	get(client)


	//fmt.Printf("%T\n", client)
}
