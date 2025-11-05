package categories

import (
    "context"
    "github.com/mytheresa/go-hiring-challenge/domain"
)

// ListCategoriesUseCase lists all categories via the repository port.
type ListCategoriesUseCase struct {
    Repo CategoryRepo
}

func (u *ListCategoriesUseCase) Execute(ctx context.Context) ([]domain.Category, error) {
    return u.Repo.List(ctx)
}

