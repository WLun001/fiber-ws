package main

import (
	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/url"
)

func main() {
	app := fiber.New()

	con, err := connectWS()
	if err != nil {
		log.Fatal(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		msg := map[string]string {
			"hello": "world",
		}
		con.WriteJSON(msg)
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":8081")
}

func connectWS() (*websocket.Conn, error) {
	u := url.URL{Scheme: "ws", Host: "localhost:3000", Path: "/ws"}
	log.Printf("connecting ws at %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to WS")
	return c, nil
}

