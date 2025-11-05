package categories

import "github.com/mytheresa/go-hiring-challenge/domain"

func toCategoryDTO(c domain.Category) CategoryDTO {
    return CategoryDTO{Code: c.Code, Name: c.Name}
}

func toResponseDTO(items []domain.Category) ResponseDTO {
    out := make([]CategoryDTO, len(items))
    for i := range items {
        out[i] = toCategoryDTO(items[i])
    }
    return ResponseDTO{Categories: out}
}

