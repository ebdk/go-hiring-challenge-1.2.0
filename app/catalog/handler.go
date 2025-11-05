package catalog

import (
	"net/http"
	"strconv"

	"github.com/mytheresa/go-hiring-challenge/app/api"
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
            api.ErrorResponse(w, http.StatusBadRequest, "invalid offset")
            return
        }
        offset = n
    }
    // limit
    limit := 10
    if s := q.Get("limit"); s != "" {
        n, err := strconv.Atoi(s)
        if err != nil {
            api.ErrorResponse(w, http.StatusBadRequest, "invalid limit")
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
        api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
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

    response := Response{Total: int(total), Products: products}
    api.OKResponse(w, response)
}

// ProductDetail represents a single product with variants for the detail endpoint.
type ProductDetail struct {
    Code     string     `json:"code"`
    Price    float64    `json:"price"`
    Category Category   `json:"category"`
    Variants []Variant  `json:"variants"`
}

type Variant struct {
    Name  string  `json:"name"`
    SKU   string  `json:"sku"`
    Price float64 `json:"price"`
}

// HandleGetByCode returns a single product identified by code, including its variants.
func (h *CatalogHandler) HandleGetByCode(w http.ResponseWriter, r *http.Request) {
    code := r.PathValue("code")
    if code == "" {
        api.ErrorResponse(w, http.StatusBadRequest, "missing code")
        return
    }

    p, found, err := h.repo.GetByCode(r.Context(), code)
    if err != nil {
        api.ErrorResponse(w, http.StatusInternalServerError, err.Error())
        return
    }
    if !found {
        api.ErrorResponse(w, http.StatusNotFound, "not found")
        return
    }

    // Build variants, inheriting price when variant price is not set (treated as zero)
    variants := make([]Variant, len(p.Variants))
    for i, v := range p.Variants {
        price := v.Price.InexactFloat64()
        if price == 0 {
            price = p.Price.InexactFloat64()
        }
        variants[i] = Variant{
            Name:  v.Name,
            SKU:   v.SKU,
            Price: price,
        }
    }

    resp := ProductDetail{
        Code:  p.Code,
        Price: p.Price.InexactFloat64(),
        Category: Category{
            Code: p.Category.Code,
            Name: p.Category.Name,
        },
        Variants: variants,
    }

    api.OKResponse(w, resp)
}
