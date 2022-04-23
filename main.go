package main

import (
	"log"

	"github.com/ali2210/fictional-engine/stubs/client"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

var tasks *client.Task_features
var counter int64 = 0

func Counter() int64 { return counter + 1 }

func main() {

	app_web := fiber.New(fiber.Config{
		Views: html.New("./views", ".hbs"),
	})

	app_web.Get("/", func(c *fiber.Ctx) error {

		// page rendered
		return c.Render("index", fiber.Map{
			"Title": "Fictional-Engine",
		})

	})

	app_web.Post("/", func(c *fiber.Ctx) error {

		tasks = &client.Task_features{}

		tasks.TaskName = c.FormValue("taskin")
		tasks.Deadline = c.FormValue("date&time")

		if c.FormValue("deadbox") == "on" {
			tasks.Level = 0
		}

		if c.FormValue("busybox") == "on" {
			tasks.Level = 1
		}

		if c.FormValue("started") == "on" {
			tasks.Level = 2
		}

		tasks.Id = Counter()

		log.Println("Tasks details :", tasks)

		go func() error {

			client.StoreTask(tasks)
			client.Set_Route(client.Create)

			err := client.Client_Init()
			if err != nil {
				log.Println("Client is not started", err)
				return err
			}

			return nil
		}()

		// page rendered
		return c.Render("index", fiber.Map{
			"Title": "Fictional-Engine",
		})
	})

	app_web.Post("/delete", func(ctx *fiber.Ctx) error {

		tasks := &client.Task_features{}
		tasks.TaskName = ""
		tasks.Deadline = ctx.FormValue("date")
		tasks.Level = 0
		tasks.Id = Counter()

		log.Println("Tasks details :", tasks)
		go func() error {

			client.StoreTask(tasks)
			client.Set_Route(client.Delete)

			err := client.Client_Init()
			if err != nil {
				log.Println("Client is not started", err)
				return err
			}
			return nil
		}()

		return ctx.Render("index", fiber.Map{
			"Title": "Fictional-Engine",
		})
	})

	err := app_web.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
