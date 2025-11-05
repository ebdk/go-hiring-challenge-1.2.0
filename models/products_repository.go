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
func (r *ProductsRepository) ListProducts(ctx context.Context, offset, limit int, category string, priceLT *float64) ([]Product, int64, error) {
    // Base query
    q := r.db.WithContext(ctx).Model(&Product{})
    if category != "" {
        q = q.Where("category_id = (SELECT id FROM categories WHERE code = ?)", category)
    }
    if priceLT != nil {
        q = q.Where("price < ?", *priceLT)
    }

    var total int64
    if err := q.Count(&total).Error; err != nil {
        return nil, 0, err
    }

    var products []Product
    if err := q.
        Preload("Variants").
        Preload("Category").
        Offset(offset).
        Limit(limit).
        Find(&products).Error; err != nil {
        return nil, 0, err
    }

    return products, total, nil
}
