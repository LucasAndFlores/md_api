package routes

import "github.com/gofiber/fiber/v2"

func SetupAudioRoutes(api fiber.Router) {
    api.Get("/audio", handler)
}

func handler (fi *fiber.Ctx) error {
      return fi.SendStatus(200)
}
