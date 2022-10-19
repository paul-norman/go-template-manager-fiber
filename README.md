# `templateManager` Fiber Integration

This is the package for the [`templateManager`](https://github.com/paul-norman/go-template-manager) View engine integration for the [Fiber](https://gofiber.io/) framework.

For all options, please see the main repository.

## Basic Usage

```go
package main

import (
	"log"
	
	"github.com/gofiber/fiber/v2"
	TM "github.com/paul-norman/go-template-manager-fiber"
)

func main() {
	engine := TM.Init("./views", ".html")
	engine.ExcludeDirectories([]string{"layouts", "partials"}).
			Reload(true).
			Debug(true)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home.html", fiber.Map{
			"Title": "Hello, World!",
		})
	})

	log.Fatal(app.Listen(":8080"))
}
```