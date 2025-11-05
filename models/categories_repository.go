package models

import (
    "context"
    "errors"
    "github.com/mytheresa/go-hiring-challenge/domain"
    "gorm.io/gorm"
)

type CategoriesRepository struct {
    db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) *CategoriesRepository {
    return &CategoriesRepository{db: db}
}

func (r *CategoriesRepository) List(ctx context.Context) ([]Category, error) {
    var items []Category
    if err := r.db.WithContext(ctx).Find(&items).Error; err != nil {
        return nil, err
    }
    return items, nil
}

func (r *CategoriesRepository) Create(ctx context.Context, c Category) (Category, error) {
    if err := r.db.WithContext(ctx).Create(&c).Error; err != nil {
        if errors.Is(err, gorm.ErrDuplicatedKey) {
            return Category{}, domain.ErrAlreadyExists
        }
        return Category{}, err
    }
    return c, nil
}
