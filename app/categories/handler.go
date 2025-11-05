package categories

import (
    "encoding/json"
    "net/http"
)

type Response struct {
    Categories []Category `json:"categories"`
}

type Category struct {
    Code string `json:"code"`
    Name string `json:"name"`
}

type Handler struct {
    repo CategoryReader
}

func NewHandler(r CategoryReader) *Handler {
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

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(Response{Categories: out}); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

