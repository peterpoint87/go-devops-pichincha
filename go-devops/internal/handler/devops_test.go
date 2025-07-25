package handler

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "os"

    "github.com/gofiber/fiber/v2"
    "github.com/peterpoint87/go-devops-pichincha/internal/handler"
)

func setupTestApp() *fiber.App {
    os.Setenv("JWT_SECRET", "testsecret") // para pruebas
    app := fiber.New()
    app.Post("/DevOps", DevOpsHandler)
    app.All("/DevOps", func(c *fiber.Ctx) error {
        if c.Method() != "POST" {
            return c.SendString("ERROR")
        }
        return c.Next()
    })
    return app
}

func TestDevOpsHandler_Success(t *testing.T) {
    app := setupTestApp()
    payload := `{"message":"This is a test","to":"Juan Perez","from":"Rita Asturia","timeToLifeSec":45}`
    req := httptest.NewRequest("POST", "/DevOps", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-Parse-REST-API-Key", "2f5ae96c-b558-4c7b-a590-a501ae1c3f6c")

    resp, err := app.Test(req)
    if err != nil {
        t.Fatalf("Request failed: %v", err)
    }

    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected 200 OK, got %d", resp.StatusCode)
    }

    if resp.Header.Get("X-JWT-KWY") == "" {
        t.Errorf("Expected JWT header, got empty")
    }
}

func TestDevOpsHandler_InvalidAPIKey(t *testing.T) {
    app := setupTestApp()
    req := httptest.NewRequest("POST", "/DevOps", nil)
    req.Header.Set("X-Parse-REST-API-Key", "invalid-key")
    resp, _ := app.Test(req)

    if resp.StatusCode != http.StatusUnauthorized {
        t.Errorf("Expected 401 Unauthorized, got %d", resp.StatusCode)
    }
}

func TestDevOpsHandler_InvalidMethod(t *testing.T) {
    app := setupTestApp()
    req := httptest.NewRequest("GET", "/DevOps", nil)
    resp, _ := app.Test(req)

    if resp.StatusCode != http.StatusOK {
        t.Errorf("Expected 200 OK for GET fallback, got %d", resp.StatusCode)
    }
}