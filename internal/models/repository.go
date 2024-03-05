package models

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user User) (*User, error)
}
