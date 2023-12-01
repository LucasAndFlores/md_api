package service

import (
	"context"

	"github.com/LucasAndFlores/md_api/internal/entity"
	"github.com/LucasAndFlores/md_api/internal/repository"
)

type UserService struct {
    userRepository repository.IUserRepository
}

type IUserService interface {
    CreateUser(context.Context, entity.UserDTOInput) error
}

func NewUserService(r repository.IUserRepository) IUserService {
    return &UserService{
        userRepository: r,
    }
    
}

func (s *UserService) CreateUser(ctx context.Context, input entity.UserDTOInput) error {
    err := s.userRepository.Insert(ctx, input)

    if err != nil {
        return err
    }

    return nil
}
