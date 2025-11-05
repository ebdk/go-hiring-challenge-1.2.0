package catalog

import (
    "context"
    "github.com/mytheresa/go-hiring-challenge/domain"
)

// ListCatalog defines the input port for listing catalog products.
type ListCatalog interface {
    Execute(ctx context.Context, q ListCatalogQuery) ([]domain.Product, int64, error)
}

// GetProduct defines the input port for fetching a product by code.
type GetProduct interface {
    Execute(ctx context.Context, code string) (domain.Product, bool, error)
}

