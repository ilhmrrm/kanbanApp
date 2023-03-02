package repository

import (
	"context"
	"errors"
	"kanbanApp/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	var res entity.User
	err := r.db.Model(&entity.User{}).Where("id = ?", id).Scan(&res).Error

	if err != nil {
		return entity.User{}, ctx.Err()
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, nil
	}

	return res, nil // TODO: replace this
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	var userMail entity.User
	err := r.db.Model(&entity.User{}).Where("email = ?", email).Scan(&userMail).Error
	if err != nil {
		return entity.User{}, ctx.Err()
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, nil
	}

	return userMail, nil // TODO: replace this
}

func (r *userRepository) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return entity.User{}, ctx.Err()
	}

	return user, nil // TODO: replace this
}

func (r *userRepository) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	err := r.db.Model(&entity.User{}).Updates(user).Error
	if err != nil {
		return entity.User{}, ctx.Err()
	}
	return user, nil // TODO: replace this
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&entity.User{}).Error
	if err != nil {
		return ctx.Err()
	}
	return nil // TODO: replace this
}
