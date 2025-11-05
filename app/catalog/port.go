package catalog

import (
    "context"
    "github.com/mytheresa/go-hiring-challenge/models"
)

// ProductReader is the port that the catalog handler depends on for reading products.
// It allows swapping the concrete repository without changing the handler.
type ProductReader interface {
    // ListProducts returns a page of products and the total count available.
    ListProducts(ctx context.Context, offset, limit int) ([]models.Product, int64, error)
}
