package main

import (
	"context"

	"github.com/LucasAndFlores/md_api/internal/config"
	"github.com/LucasAndFlores/md_api/routes"
	"github.com/LucasAndFlores/md_api/storage"
	"github.com/gofiber/fiber/v2"
)


func main() {
    e := config.LoadEnvVariables()

    if e != nil {
        panic(e)
    }

    ctx := context.Background()
    postgres, err := storage.NewPostgresConnection(ctx)
    
    if err != nil {
        panic(err)
    }

    defer postgres.Close()

    app := fiber.New()

    router := app.Group("/api")

    routes.SetupUserRoutes(router, postgres)
    routes.SetupAudioRoutes(router)

    app.Listen(":3000")


}
