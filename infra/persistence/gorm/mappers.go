package gormrepo

import "github.com/mytheresa/go-hiring-challenge/domain"

func toDomainCategory(c Category) domain.Category {
    return domain.Category{Code: c.Code, Name: c.Name}
}

func toDomainVariant(v Variant) domain.Variant {
    return domain.Variant{Name: v.Name, SKU: v.SKU, Price: v.Price}
}

func toDomainProduct(p Product) domain.Product {
    out := domain.Product{
        Code:  p.Code,
        Price: p.Price,
        Category: toDomainCategory(p.Category),
    }
    out.Variants = make([]domain.Variant, len(p.Variants))
    for i := range p.Variants {
        out.Variants[i] = toDomainVariant(p.Variants[i])
    }
    return out
}
