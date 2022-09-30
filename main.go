package main

import (
	"fmt"
	"strconv"
	"os/exec"
	"context"
	"github.com/go-redis/redis/v8"

	"html/template"
	"io"
	"net/http"
	"github.com/labstack/echo/v4"
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
type TemplateRenderer struct {
	templates *template.Template
}
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func start_server() {
	e := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("static/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", map[string]interface{}{
			"name": "Dolly!",
		})
	}).Name = "foobar"

	e.Logger.Fatal(e.Start(":8000"))
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
	err := cln.Set(ctx, "4", "four", 0).Err()

	if err != nil {
		return err
	}

	return nil
}

func get_all(cln *redis.Client) {
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

	fmt.Println(values)
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
	//set(client)
	
	get_all(client)

	//fmt.Printf("%T\n", client)
}
