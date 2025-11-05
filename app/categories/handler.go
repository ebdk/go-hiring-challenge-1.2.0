package categories

import (
    "encoding/json"
    "net/http"
    "github.com/mytheresa/go-hiring-challenge/app/api"
    "github.com/mytheresa/go-hiring-challenge/models"
)

type Response struct {
    Categories []Category `json:"categories"`
}

type Category struct {
    Code string `json:"code"`
    Name string `json:"name"`
}

type Handler struct {
    repo CategoryRepo
}

func NewHandler(r CategoryRepo) *Handler {
    return &Handler{repo: r}
}

// HandleList returns all categories.
func (h *Handler) HandleList(w http.ResponseWriter, r *http.Request) {
    res, err := h.repo.List(r.Context())
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    out := make([]Category, len(res))
    for i, c := range res {
        out[i] = Category{Code: c.Code, Name: c.Name}
    }

    api.OKResponse(w, Response{Categories: out})
}

// HandleCreate creates a new category from JSON body.
func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
    var in Category
    if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
        http.Error(w, "invalid json", http.StatusBadRequest)
        return
    }
    if in.Code == "" || in.Name == "" {
        http.Error(w, "code and name are required", http.StatusBadRequest)
        return
    }

    created, err := h.repo.Create(r.Context(), models.Category{Code: in.Code, Name: in.Name})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    out := Category{Code: created.Code, Name: created.Name}
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    if err := json.NewEncoder(w).Encode(out); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
