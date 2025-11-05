package categories

import (
    "encoding/json"
    "net/http"
    "github.com/mytheresa/go-hiring-challenge/app/api"
    "github.com/mytheresa/go-hiring-challenge/domain"
)

// DTOs are in dto.go

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

    api.OKResponse(w, toResponseDTO(res))
}

// HandleCreate creates a new category from JSON body.
func (h *Handler) HandleCreate(w http.ResponseWriter, r *http.Request) {
    // Limit body size and reject unknown fields
    r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MiB
    dec := json.NewDecoder(r.Body)
    dec.DisallowUnknownFields()

    var in CategoryDTO
    if err := dec.Decode(&in); err != nil {
        api.ErrorResponse(w, http.StatusBadRequest, "invalid json")
        return
    }
    if in.Code == "" || in.Name == "" {
        api.ErrorResponse(w, http.StatusUnprocessableEntity, "code and name are required")
        return
    }

    created, err := h.repo.Create(r.Context(), domain.Category{Code: in.Code, Name: in.Name})
    if err != nil {
        if err == domain.ErrAlreadyExists {
            api.ErrorResponse(w, http.StatusConflict, "category already exists")
            return
        }
        api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
        return
    }

    out := toCategoryDTO(created)
    api.CreatedResponse(w, out)
}
