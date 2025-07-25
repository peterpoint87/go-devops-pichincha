package main

import (
    "log"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/peterpoint87/go-devops-pichincha/internal/handler"
)

func main() {
    app := fiber.New()

    app.Post("/DevOps", handler.DevOpsHandler)
    app.All("/DevOps", func(c *fiber.Ctx) error {
        if c.Method() != "POST" {
            return c.SendString("ERROR")
        }
        return c.Next()
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Fatal(app.Listen(":" + port))
}