package main

import (
	"fmt"
	"strconv"
	"os/exec"
	"context"
	"github.com/go-redis/redis/v8"

	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)


var (
	ctx = context.TODO()
)

// SHELL COMMANDS
func redis_server_stop() {
	cmd := exec.Command("/bin/sh", "-c", "sudo service redis-server stop")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	} else { fmt.Println(string(stdout)) }
}
func redis_server_start() {
	cmd := exec.Command("/bin/sh", "-c", "sudo service redis-server start")
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	} else { fmt.Println(string(stdout)) }
}
// ----------------------------


// SERVER FUNCTIONS
func start_server(base []string) {
	engine := html.New("./public", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map {
			"Base": base,
		})
	})

	app.Static("/assets", "./assets")

	log.Fatal(app.Listen(":3000"))
}
// ----------------------------


// REDIS FUNCTIONS
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
	//var numbers [5]int = [5]int{1,2,3,4,5}
	err := cln.Set(ctx, "8", "eight", 0).Err()

	if err != nil {
		return err
	}

	return nil
}

func get_all(cln *redis.Client) []string {
	var (
		cursor int = 1
		values = []string{}
	)

	for {
		id := strconv.Itoa(cursor)
		x, err := cln.Get(ctx, id).Result()

		if err == redis.Nil {
			cursor = 0
		} else if err != nil {
			panic(err)
		} else {
			values = append(values, x)
		}

		if cursor == 0 {
			break
		} else {
			cursor += 1
		}
	}

	return values
}

func checkConn(cln *redis.Client) {
	pong, err := cln.Ping(cln.Context()).Result()
	fmt.Println(pong, err)
}
// ----------------------------


func main() {
	fmt.Println("Testing Golang Redis\n")
	redis_server_start()

	client := redis.NewClient(&redis.Options {
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
	checkConn(client)

	base := get_all(client)
	start_server(base)
}
