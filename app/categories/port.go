package categories

import (
    "context"
    "github.com/mytheresa/go-hiring-challenge/domain"
)

// CategoryRepo is the port for listing and creating categories.
type CategoryRepo interface {
    List(ctx context.Context) ([]domain.Category, error)
    Create(ctx context.Context, c domain.Category) (domain.Category, error)
}
