package catalog

import (
    "context"
    "github.com/mytheresa/go-hiring-challenge/domain"
)

// listAdapter adapts UseCase.List to the ListCatalog interface.
type listAdapter struct{ uc *UseCase }

func (a *listAdapter) Execute(ctx context.Context, q ListCatalogQuery) ([]domain.Product, int64, error) {
    return a.uc.List(ctx, q)
}

// getAdapter adapts UseCase.GetByCode to the GetProduct interface.
type getAdapter struct{ uc *UseCase }

func (a *getAdapter) Execute(ctx context.Context, code string) (domain.Product, bool, error) {
    return a.uc.GetByCode(ctx, code)
}

