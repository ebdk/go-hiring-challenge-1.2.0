package models

import (
    "context"
    "gorm.io/gorm"
)

type ProductsRepository struct {
	db *gorm.DB
}

func NewProductsRepository(db *gorm.DB) *ProductsRepository {
	return &ProductsRepository{
		db: db,
	}
}

func (r *ProductsRepository) GetAllProducts() ([]Product, error) {
    var products []Product
    if err := r.db.Preload("Variants").Preload("Category").Find(&products).Error; err != nil {
        return nil, err
    }
    return products, nil
}

// ListProducts returns a page of products with total count.
func (r *ProductsRepository) ListProducts(ctx context.Context, offset, limit int) ([]Product, int64, error) {
    var total int64
    if err := r.db.WithContext(ctx).Model(&Product{}).Count(&total).Error; err != nil {
        return nil, 0, err
    }

    var products []Product
    if err := r.db.WithContext(ctx).
        Preload("Variants").
        Preload("Category").
        Offset(offset).
        Limit(limit).
        Find(&products).Error; err != nil {
        return nil, 0, err
    }

    return products, total, nil
}
