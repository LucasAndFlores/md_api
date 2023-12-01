package routes

import (

	"github.com/LucasAndFlores/md_api/internal/controller"
	"github.com/LucasAndFlores/md_api/internal/repository"
	"github.com/LucasAndFlores/md_api/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupUserRoutes (api fiber.Router, postgres *pgxpool.Pool) {

    repository := repository.NewUserRepository(postgres)

    service := service.NewUserService(repository)
    createUserController := controller.NewUserController(service)

    api.Post("/users", createUserController.Handle)
}
