package categories

import (
    "encoding/json"
    "net/http"
    "github.com/mytheresa/go-hiring-challenge/app/api"
    "github.com/mytheresa/go-hiring-challenge/domain"
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
    // Limit body size and reject unknown fields
    r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MiB
    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()

    var in Category
    if err := dec.Decode(&in); err != nil {
        api.ErrorResponse(w, http.StatusBadRequest, "invalid json")
        return
    }
    if in.Code == "" || in.Name == "" {
        api.ErrorResponse(w, http.StatusBadRequest, "code and name are required")
        return
    }

    created, err := h.repo.Create(r.Context(), models.Category{Code: in.Code, Name: in.Name})
    if err != nil {
        if err == domain.ErrAlreadyExists {
            api.ErrorResponse(w, http.StatusConflict, "category already exists")
            return
        }
        api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
        return
    }

    out := Category{Code: created.Code, Name: created.Name}
    api.CreatedResponse(w, out)
}
