package controller

import (
	"context"
	"net/http"

	"github.com/LucasAndFlores/md_api/internal/entity"
	"github.com/LucasAndFlores/md_api/internal/service"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
    UserService service.IUserService
}

func NewUserController (s service.IUserService) *UserController {
    return &UserController{UserService: s}
}

func (c *UserController) Handle(fi *fiber.Ctx) error {
    var user entity.UserDTOInput

    if err := fi.BodyParser(&user); err != nil {
        return err
    }  

    ctx := context.Background()

    c.UserService.CreateUser(ctx, user)

    fi.Status(http.StatusCreated).Send(nil)

    return nil

}
