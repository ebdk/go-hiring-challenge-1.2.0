package categories

import (
    "context"
    "github.com/mytheresa/go-hiring-challenge/models"
)

// CategoryReader is the port for listing categories.
type CategoryReader interface {
    List(ctx context.Context) ([]models.Category, error)
}

