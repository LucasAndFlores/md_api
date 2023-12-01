package repository

import (
	"context"

	"github.com/LucasAndFlores/md_api/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type (
    UserRepository struct {
        pg *pgxpool.Pool
    }
)

type IUserRepository interface {
    Insert(context.Context, entity.UserDTOInput) error
}

func NewUserRepository(conn *pgxpool.Pool) IUserRepository {
    return &UserRepository{
        pg: conn,
    }
}

func (ur *UserRepository) Insert(ctx context.Context, u entity.UserDTOInput) error {
    _, err := ur.pg.Exec(ctx, "INSERT INTO users_md (name, surname, email, cellphone, password) VALUES ($1, $2, $3, $4, $5)", u.Name, u.Surname, u.Email, u.Cellphone, u.Password)

    if err != nil {
        return  err
    }

    return nil
}
