package categories

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "context"

    "github.com/mytheresa/go-hiring-challenge/domain"
    "github.com/stretchr/testify/assert"
)

type fakeCategoryRepo struct {
    items []domain.Category
    err   error
}

func (f *fakeCategoryRepo) List(ctx context.Context) ([]domain.Category, error) {
    return f.items, f.err
}

func (f *fakeCategoryRepo) Create(ctx context.Context, c domain.Category) (domain.Category, error) {
    return c, nil
}

func TestHandleList_ReturnsCategories(t *testing.T) {
    repo := &fakeCategoryRepo{items: []domain.Category{
        {Code: "clothing", Name: "Clothing"},
        {Code: "shoes", Name: "Shoes"},
        {Code: "accessories", Name: "Accessories"},
    }}
    h := NewHandler(repo)

    req := httptest.NewRequest(http.MethodGet, "/categories", nil)
    rec := httptest.NewRecorder()

    h.HandleList(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
    expected := `{"categories":[{"code":"clothing","name":"Clothing"},{"code":"shoes","name":"Shoes"},{"code":"accessories","name":"Accessories"}]}`
    assert.JSONEq(t, expected, rec.Body.String())
}
