package handler

import (
    "github.com/gofiber/fiber/v2"
    "github.com/peterpoint87/go-devops-pichincha/pkg/jwtutil"
)

const expectedAPIKey = "2f5ae96c-b558-4c7b-a590-a501ae1c3f6c"

type DevOpsRequest struct {
    Message       string `json:"message"`
    To            string `json:"to"`
    From          string `json:"from"`
    TimeToLifeSec int    `json:"timeToLifeSec"`
}

type DevOpsResponse struct {
    Message string `json:"message"`
}

func DevOpsHandler(c *fiber.Ctx) error {
    apiKey := c.Get("X-Parse-REST-API-Key")
    if apiKey != expectedAPIKey {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid API Key"})
    }

    var req DevOpsRequest
    if err := c.BodyParser(&req); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
    }

    token := jwtutil.GenerateJWT(req.To, req.TimeToLifeSec)
    c.Set("X-JWT-KWY", token)

    return c.JSON(DevOpsResponse{Message: "Hello " + req.To + " your message will be send"})
}