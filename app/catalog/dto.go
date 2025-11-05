package catalog

// DTOs for HTTP responses

type ResponseDTO struct {
    Total    int          `json:"total"`
    Products []ProductDTO `json:"products"`
}

type ProductDTO struct {
    Code     string      `json:"code"`
    Price    float64     `json:"price"`
    Category CategoryDTO `json:"category"`
}

type CategoryDTO struct {
    Code string `json:"code"`
    Name string `json:"name"`
}

type ProductDetailDTO struct {
    Code     string        `json:"code"`
    Price    float64       `json:"price"`
    Category CategoryDTO   `json:"category"`
    Variants []VariantDTO  `json:"variants"`
}

type VariantDTO struct {
    Name  string  `json:"name"`
    SKU   string  `json:"sku"`
    Price float64 `json:"price"`
}

