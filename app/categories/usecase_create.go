package categories

import (
    "context"
    "github.com/mytheresa/go-hiring-challenge/domain"
)

// CreateCategoryUseCase validates and creates a category.
type CreateCategoryUseCase struct {
    Repo CategoryRepo
}

func (u *CreateCategoryUseCase) Execute(ctx context.Context, c domain.Category) (domain.Category, error) {
    if c.Code == "" || c.Name == "" {
        return domain.Category{}, domain.ErrInvalid
    }
    return u.Repo.Create(ctx, c)
}

