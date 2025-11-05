package catalog

import "github.com/mytheresa/go-hiring-challenge/domain"

func toProductDTO(p domain.Product) ProductDTO {
    return ProductDTO{
        Code:  p.Code,
        Price: p.Price.InexactFloat64(),
        Category: CategoryDTO{Code: p.Category.Code, Name: p.Category.Name},
    }
}

func toResponseDTO(items []domain.Product, total int64) ResponseDTO {
    out := make([]ProductDTO, len(items))
    for i := range items {
        out[i] = toProductDTO(items[i])
    }
    return ResponseDTO{Total: int(total), Products: out}
}

func toProductDetailDTO(p domain.Product) ProductDetailDTO {
    variants := make([]VariantDTO, len(p.Variants))
    base := p.Price.InexactFloat64()
    for i, v := range p.Variants {
        price := v.Price.InexactFloat64()
        if price == 0 {
            price = base
        }
        variants[i] = VariantDTO{Name: v.Name, SKU: v.SKU, Price: price}
    }
    return ProductDetailDTO{
        Code:  p.Code,
        Price: base,
        Category: CategoryDTO{Code: p.Category.Code, Name: p.Category.Name},
        Variants: variants,
    }
}

