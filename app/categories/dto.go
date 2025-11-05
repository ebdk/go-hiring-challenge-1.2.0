package categories

type CategoryDTO struct {
    Code string `json:"code"`
    Name string `json:"name"`
}

type ResponseDTO struct {
    Categories []CategoryDTO `json:"categories"`
}

