package categories

import (
    "context"
    "github.com/mytheresa/go-hiring-challenge/models"
)

// CategoryRepo is the port for listing and creating categories.
type CategoryRepo interface {
    List(ctx context.Context) ([]models.Category, error)
    Create(ctx context.Context, c models.Category) (models.Category, error)
}
