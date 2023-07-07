package module

import (
	"context"
	"schoolcms/internal/constant/dto"
)

type User interface {
	CreateUser(ctx context.Context, ur dto.User) (dto.User, error)
}
