package categories

import (
    "context"
    "github.com/mytheresa/go-hiring-challenge/domain"
)

// ListCategories defines the input port for listing categories.
type ListCategories interface {
    Execute(ctx context.Context) ([]domain.Category, error)
}

// CreateCategory defines the input port for creating a category.
type CreateCategory interface {
    Execute(ctx context.Context, c domain.Category) (domain.Category, error)
}

