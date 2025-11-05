package domain

import "github.com/shopspring/decimal"

// Variant is a domain entity for a product variant.
type Variant struct {
    Name  string
    SKU   string
    Price decimal.Decimal // zero means inherit from product in our current model
}

