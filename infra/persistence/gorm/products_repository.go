package gormrepo

import (
    "context"
    "errors"
    "github.com/mytheresa/go-hiring-challenge/domain"
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

func (r *ProductsRepository) GetAllProducts() ([]domain.Product, error) {
    var products []Product
    if err := r.db.Preload("Variants").Preload("Category").Find(&products).Error; err != nil {
        return nil, err
    }
    out := make([]domain.Product, len(products))
    for i := range products {
        out[i] = toDomainProduct(products[i])
    }
    return out, nil
}

// ListProducts returns a page of products with total count.
func (r *ProductsRepository) ListProducts(ctx context.Context, offset, limit int, category string, priceLT *float64) ([]domain.Product, int64, error) {
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
        Order("code ASC").
        Offset(offset).
        Limit(limit).
        Find(&products).Error; err != nil {
        return nil, 0, err
    }

    out := make([]domain.Product, len(products))
    for i := range products {
        out[i] = toDomainProduct(products[i])
    }
    return out, total, nil
}

// GetByCode fetches a single product by its code with preloaded relations.
func (r *ProductsRepository) GetByCode(ctx context.Context, code string) (domain.Product, bool, error) {
    var p Product
    err := r.db.WithContext(ctx).
        Preload("Variants").
        Preload("Category").
        Where("code = ?", code).
        First(&p).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return domain.Product{}, false, nil
    }
    if err != nil {
        return domain.Product{}, false, err
    }
    return toDomainProduct(p), true, nil
}
