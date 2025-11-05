package ports

import (
    "context"
    "github.com/mytheresa/go-hiring-challenge/domain"
)

// CatalogRepository is the output port for reading products from a persistence adapter.
type CatalogRepository interface {
    // ListProducts returns a page of products and total count, with optional filters.
    ListProducts(ctx context.Context, offset, limit int, category string, priceLT *float64) ([]domain.Product, int64, error)
    // GetByCode fetches a single product by its code.
    GetByCode(ctx context.Context, code string) (domain.Product, bool, error)
}

