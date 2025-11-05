package models

import (
    "context"
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

