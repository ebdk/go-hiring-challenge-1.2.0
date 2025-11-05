package domain

import "github.com/shopspring/decimal"

// Product is a domain entity for catalog products.
type Product struct {
    Code     string
    Price    decimal.Decimal
    Category Category
    Variants []Variant
}

