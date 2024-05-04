package domain

import (
	"context"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID         int64  `gorm:"primaryKey"`
	UserName       string `gorm:"unique"`
	NickName       string
	PasswordDigest string
}

func (User) TableName() string {
	return `user`
}

type UserService interface {
}

type UserRepository interface {
	GetByUserName(ctx context.Context, userName string) (*User, error)
	Create(ctx context.Context, user *User) error
}
