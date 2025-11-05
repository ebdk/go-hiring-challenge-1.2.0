package ports

import (
    "context"
    "github.com/mytheresa/go-hiring-challenge/domain"
)

// CategoryRepository is the output port for category persistence.
type CategoryRepository interface {
    List(ctx context.Context) ([]domain.Category, error)
    Create(ctx context.Context, c domain.Category) (domain.Category, error)
}

