package repository

import (
	"context"
	"github.com/LXJ0000/go-rpc/app/user/internal/domain"
	"github.com/LXJ0000/go-rpc/orm"
)

type userRepository struct {
	db orm.Database
}

func NewUserRepository(db orm.Database) domain.UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) GetByUserName(ctx context.Context, userName string) (*domain.User, error) {
	filter := map[string]interface{}{
		"user_name": userName,
	}
	user, err := u.db.FindOne(ctx, &domain.User{}, filter)
	if err != nil {
		return nil, err
	}
	return user.(*domain.User), nil
}

func (u *userRepository) Create(ctx context.Context, user *domain.User) error {
	return u.db.InsertOne(ctx, &domain.User{}, user)
}
