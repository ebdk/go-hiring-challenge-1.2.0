package catalog

import (
    "encoding/json"
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"
    "context"

    "github.com/mytheresa/go-hiring-challenge/models"
    "github.com/shopspring/decimal"
    "github.com/stretchr/testify/assert"
)

// fakeProductRepo implements ProductReader for testing the list handler
type fakeProductRepo struct{
    // inputs captured
    gotOffset   int
    gotLimit    int
    gotCategory string
    gotPriceLT  *float64

    // outputs configured
    products []models.Product
    total    int64
    retErr   error
}

func (f *fakeProductRepo) ListProducts(ctx context.Context, offset, limit int, category string, priceLT *float64) ([]models.Product, int64, error) {
    f.gotOffset, f.gotLimit, f.gotCategory, f.gotPriceLT = offset, limit, category, priceLT
    return f.products, f.total, f.retErr
}

func (f *fakeProductRepo) GetByCode(ctx context.Context, code string) (models.Product, bool, error) {
    return models.Product{}, false, nil
}

func TestHandleGet_DefaultsAndMapping(t *testing.T) {
    repo := &fakeProductRepo{
        products: []models.Product{
            {Code: "PROD001", Price: decimal.NewFromFloat(10.99), Category: models.Category{Code: "clothing", Name: "Clothing"}},
            {Code: "PROD002", Price: decimal.NewFromFloat(12.49), Category: models.Category{Code: "shoes", Name: "Shoes"}},
        },
        total: 2,
    }
    h := NewCatalogHandler(repo)

    req := httptest.NewRequest(http.MethodGet, "/catalog", nil)
    rec := httptest.NewRecorder()

    h.HandleGet(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
    // defaults should be applied
    assert.Equal(t, 0, repo.gotOffset)
    assert.Equal(t, 10, repo.gotLimit)
    assert.Equal(t, "", repo.gotCategory)
    assert.Nil(t, repo.gotPriceLT)

    var body struct{
        Total int `json:"total"`
        Products []struct{
            Code string `json:"code"`
            Price float64 `json:"price"`
            Category struct{ Code, Name string } `json:"category"`
        } `json:"products"`
    }
    err := json.Unmarshal(rec.Body.Bytes(), &body)
    assert.NoError(t, err)
    assert.Equal(t, 2, body.Total)
    assert.Len(t, body.Products, 2)
    assert.Equal(t, "clothing", body.Products[0].Category.Code)
}

func TestHandleGet_PaginationAndFilters(t *testing.T) {
    repo := &fakeProductRepo{}
    h := NewCatalogHandler(repo)

    req := httptest.NewRequest(http.MethodGet, "/catalog?offset=5&limit=200&category=shoes&price_lt=10.5", nil)
    rec := httptest.NewRecorder()

    h.HandleGet(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code)
    assert.Equal(t, 5, repo.gotOffset)
    // limit should be clamped to 100
    assert.Equal(t, 100, repo.gotLimit)
    assert.Equal(t, "shoes", repo.gotCategory)
    if assert.NotNil(t, repo.gotPriceLT) {
        assert.InDelta(t, 10.5, *repo.gotPriceLT, 0.00001)
    }
}

func TestHandleGet_BadQueryParams(t *testing.T) {
    h := NewCatalogHandler(&fakeProductRepo{})

    // invalid offset
    rec := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/catalog?offset=oops", nil)
    h.HandleGet(rec, req)
    assert.Equal(t, http.StatusBadRequest, rec.Code)

    // invalid limit
    rec = httptest.NewRecorder()
    req = httptest.NewRequest(http.MethodGet, "/catalog?limit=bad", nil)
    h.HandleGet(rec, req)
    assert.Equal(t, http.StatusBadRequest, rec.Code)

    // invalid price_lt
    rec = httptest.NewRecorder()
    req = httptest.NewRequest(http.MethodGet, "/catalog?price_lt=bad", nil)
    h.HandleGet(rec, req)
    assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandleGet_RepoError(t *testing.T) {
    repo := &fakeProductRepo{retErr: errors.New("boom")}
    h := NewCatalogHandler(repo)
    rec := httptest.NewRecorder()
    req := httptest.NewRequest(http.MethodGet, "/catalog", nil)
    h.HandleGet(rec, req)
    assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

