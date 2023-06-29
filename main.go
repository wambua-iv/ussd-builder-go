package main

import (
	"fmt"
	"log"

	//"net/http"
	ussdbuilder "ussd-builder/ussd-builder"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var ussd ussdbuilder.UssdMenu

func main() {
	ussdBuilder := fiber.New()
	ussdBuilder.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	ussdBuilder.Get("init", initialize)
	ussdBuilder.Post("ussd", register)

	log.Fatal(ussdBuilder.Listen(":3080"))

}

func initialize(ctx *fiber.Ctx) error {
	return ctx.SendString(ussd.CON("Hello World"))
}

func register(ctx *fiber.Ctx) error {
	ctx.BodyParser(&ussd.Args)

	res := ussd.BuildState(map[int]string{
		1: "here",
		2: "Build Home",
	})

	fmt.Printf("URL.Body = %q\n", res)
	fmt.Printf("URL.Body = %q\n", ctx.Body())

	fmt.Printf("URL.Body = %q\n", ussd.Args)
	fmt.Printf("URL.Path = %q\n", ussd.GetRoutes("1*2*4*5*6"))
	fmt.Printf("URL.Path = %q\n", ussd.GetCurrentRoute("1*2*4*5*7"))
	return ctx.SendString(ussd.CON("HELLO WORLD \n 1. Build Home"))
}

// handler echoes the Path component of the requested URL.
// func handler(w http.ResponseWriter, r *http.Request) string {

//     server := http.NewServeMux()
//     server.HandleFunc("/", handler) // each request calls handler
//     log.Fatal(http.ListenAndServe(":8000", nil))

//     body, _ := ioutil.ReadAll(r.Body)
//     fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
//     fmt.Fprintf(w, "URL.Body = %q\n", body)

//     var ussd ussdbuilder.UssdMenu

//     fmt.Fprintf(w, "URL.Path = %q\n", ussd.CON("world"))
//     fmt.Fprintf(w, "URL.Path = %q\n", ussd.GetRoute("1*2*4*5*6"))

// }
