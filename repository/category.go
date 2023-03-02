package repository

import (
	"context"
	"errors"
	"kanbanApp/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	var result []entity.Category

	err := r.db.Model(&entity.Category{}).Where("user_id = ?", id).Scan(&result).Error
	if err != nil {
		return nil, ctx.Err()
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return []entity.Category{}, nil
	}

	return result, nil // TODO: replace this
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	err = r.db.Create(&category).Error
	if err != nil {
		return 0, ctx.Err() // TODO: replace this
	}

	return category.ID, nil
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	err := r.db.Create(&categories).Error
	if err != nil {
		return ctx.Err()
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	var res entity.Category

	err := r.db.Where("id = ?", id).Model(&entity.Category{}).Scan(&res).Error
	if err != nil {
		return entity.Category{}, ctx.Err()
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.Category{}, nil
	}

	return res, nil // TODO: replace this
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	err := r.db.Model(&entity.Category{}).Updates(category).Error
	if err != nil {
		return ctx.Err()
	}
	return nil // TODO: replace this
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	err := r.db.Where("id = ?", id).Delete(&entity.Category{}).Error
	if err != nil {
		return ctx.Err()
	}
	return nil // TODO: replace this
}
