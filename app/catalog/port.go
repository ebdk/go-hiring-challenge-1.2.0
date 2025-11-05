package catalog

import (
    "context"
    "github.com/mytheresa/go-hiring-challenge/domain"
)

// ProductReader is the port that the catalog handler depends on for reading products.
// It allows swapping the concrete repository without changing the handler.
type ProductReader interface {
    // ListProducts returns a page of products and the total count available.
    // category: optional category code filter; empty for no filter.
    // priceLT: optional price upper bound filter; nil for no filter.
    ListProducts(ctx context.Context, offset, limit int, category string, priceLT *float64) ([]domain.Product, int64, error)
    // GetByCode fetches a product by its code, including variants and category.
    // Returns found=false when the product does not exist.
    GetByCode(ctx context.Context, code string) (domain.Product, bool, error)
}
