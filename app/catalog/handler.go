package catalog

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type Response struct {
    Total    int       `json:"total"`
    Products []Product `json:"products"`
}

type Product struct {
    Code  string  `json:"code"`
    Price float64 `json:"price"`
    Category Category `json:"category"`
}

type Category struct {
    Code string `json:"code"`
    Name string `json:"name"`
}

type CatalogHandler struct {
	repo ProductReader
}

func NewCatalogHandler(r ProductReader) *CatalogHandler {
	return &CatalogHandler{
		repo: r,
	}
}

func (h *CatalogHandler) HandleGet(w http.ResponseWriter, r *http.Request) {
    // Parse pagination params
    q := r.URL.Query()
    // offset
    offset := 0
    if s := q.Get("offset"); s != "" {
        n, err := strconv.Atoi(s)
        if err != nil || n < 0 {
            http.Error(w, "invalid offset", http.StatusBadRequest)
            return
        }
        offset = n
    }
    // limit
    limit := 10
    if s := q.Get("limit"); s != "" {
        n, err := strconv.Atoi(s)
        if err != nil {
            http.Error(w, "invalid limit", http.StatusBadRequest)
            return
        }
        if n < 1 {
            n = 1
        }
        if n > 100 {
            n = 100
        }
        limit = n
    }

    // Filters
    category := q.Get("category")
    var priceLT *float64
    if s := q.Get("price_lt"); s != "" {
        if v, err := strconv.ParseFloat(s, 64); err == nil {
            priceLT = &v
        } else {
            http.Error(w, "invalid price_lt", http.StatusBadRequest)
            return
        }
    }

    res, total, err := h.repo.ListProducts(r.Context(), offset, limit, category, priceLT)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	// Map response
    products := make([]Product, len(res))
    for i, p := range res {
        products[i] = Product{
            Code:  p.Code,
            Price: p.Price.InexactFloat64(),
            Category: Category{
                Code: p.Category.Code,
                Name: p.Category.Name,
            },
        }
    }

	// Return the products as a JSON response
	w.Header().Set("Content-Type", "application/json")

    response := Response{Total: int(total), Products: products}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
