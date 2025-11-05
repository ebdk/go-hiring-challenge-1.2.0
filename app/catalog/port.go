package catalog

import "github.com/mytheresa/go-hiring-challenge/models"

// ProductReader is the port that the catalog handler depends on for reading products.
// It allows swapping the concrete repository without changing the handler.
type ProductReader interface {
    GetAllProducts() ([]models.Product, error)
}

