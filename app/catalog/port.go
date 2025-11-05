package catalog

import (
    ports "github.com/mytheresa/go-hiring-challenge/domain/ports"
)

// ProductReader aliases the domain output port; kept for backward compatibility inside the app layer.
type ProductReader = ports.CatalogRepository
