package catalog

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "time"

    "github.com/mytheresa/go-hiring-challenge/domain"
    "github.com/shopspring/decimal"
    "github.com/stretchr/testify/assert"
    "context"
)

// fakeRepo implements ProductReader for tests
type fakeRepo struct {
    product domain.Product
    found   bool
}

func (f *fakeRepo) ListProducts(ctx context.Context, offset, limit int, category string, priceLT *float64) ([]domain.Product, int64, error) {
    return nil, 0, nil
}
func (f *fakeRepo) GetByCode(ctx context.Context, code string) (domain.Product, bool, error) {
    if f.found && code == f.product.Code {
        return f.product, true, nil
    }
    return domain.Product{}, false, nil
}

func TestHandleGetByCode_Success(t *testing.T) {
    repo := &fakeRepo{
        found: true,
        product: domain.Product{
            Code:  "PROD001",
            Price: decimal.NewFromFloat(10.99),
            Category: domain.Category{Code: "clothing", Name: "Clothing"},
            Variants: []domain.Variant{
                {Name: "Variant A", SKU: "SKU001A", Price: decimal.NewFromFloat(11.99)},
                {Name: "Variant B", SKU: "SKU001B", Price: decimal.Zero}, // inherit
            },
        },
    }
    h := NewCatalogHandler(repo)

    req := httptest.NewRequest(http.MethodGet, "/catalog/PROD001", nil)
    req = req.WithContext(context.WithValue(req.Context(), http.ServerContextKey, &http.Server{ReadTimeout: time.Second}))
    req.SetPathValue("code", "PROD001")
    rec := httptest.NewRecorder()

    h.HandleGetByCode(rec, req)

    assert.Equal(t, http.StatusOK, rec.Code)
    var body struct{
        Code string `json:"code"`
        Price float64 `json:"price"`
        Category struct{ Code, Name string } `json:"category"`
        Variants []struct{ Name, SKU string; Price float64 } `json:"variants"`
    }
    err := json.Unmarshal(rec.Body.Bytes(), &body)
    assert.NoError(t, err)
    assert.Equal(t, "PROD001", body.Code)
    assert.Equal(t, 10.99, body.Price)
    assert.Equal(t, "clothing", body.Category.Code)
    assert.Len(t, body.Variants, 2)
    assert.Equal(t, 11.99, body.Variants[0].Price)
    // inherited price for second variant
    assert.Equal(t, 10.99, body.Variants[1].Price)
}

func TestHandleGetByCode_NotFound(t *testing.T) {
    repo := &fakeRepo{found: false}
    h := NewCatalogHandler(repo)

    req := httptest.NewRequest(http.MethodGet, "/catalog/PROD999", nil)
    req.SetPathValue("code", "PROD999")
    rec := httptest.NewRecorder()

    h.HandleGetByCode(rec, req)

    assert.Equal(t, http.StatusNotFound, rec.Code)
}
