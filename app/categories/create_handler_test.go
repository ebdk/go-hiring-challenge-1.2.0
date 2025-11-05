package categories

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"
    "context"

    "github.com/mytheresa/go-hiring-challenge/models"
    "github.com/mytheresa/go-hiring-challenge/domain"
    "github.com/stretchr/testify/assert"
)

type fakeCategoryRepoCreate struct{
    items []models.Category
    createErr error
}

func (f *fakeCategoryRepoCreate) List(ctx context.Context) ([]models.Category, error) { return f.items, nil }
func (f *fakeCategoryRepoCreate) Create(ctx context.Context, c models.Category) (models.Category, error) {
    if f.createErr != nil { return models.Category{}, f.createErr }
    return c, nil
}

func TestHandleCreate_Success(t *testing.T) {
    repo := &fakeCategoryRepoCreate{}
    h := NewHandler(repo)

    body := bytes.NewBufferString(`{"code":"bags","name":"Bags"}`)
    req := httptest.NewRequest(http.MethodPost, "/categories", body)
    rec := httptest.NewRecorder()

    h.HandleCreate(rec, req)

    assert.Equal(t, http.StatusCreated, rec.Code)
    assert.Equal(t, "application/json", rec.Header().Get("Content-Type"))
    assert.JSONEq(t, `{"code":"bags","name":"Bags"}`, rec.Body.String())
}

func TestHandleCreate_BadRequest(t *testing.T) {
    repo := &fakeCategoryRepoCreate{}
    h := NewHandler(repo)

    body := bytes.NewBufferString(`{"code":"","name":"X"}`)
    req := httptest.NewRequest(http.MethodPost, "/categories", body)
    rec := httptest.NewRecorder()

    h.HandleCreate(rec, req)
    assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestHandleCreate_Conflict(t *testing.T) {
    repo := &fakeCategoryRepoCreate{createErr: domain.ErrAlreadyExists}
    h := NewHandler(repo)

    body := bytes.NewBufferString(`{"code":"bags","name":"Bags"}`)
    req := httptest.NewRequest(http.MethodPost, "/categories", body)
    rec := httptest.NewRecorder()

    h.HandleCreate(rec, req)
    assert.Equal(t, http.StatusConflict, rec.Code)
}
