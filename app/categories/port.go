package categories

import (
    ports "github.com/mytheresa/go-hiring-challenge/domain/ports"
)

// CategoryRepo aliases the domain output port; kept local for app wiring convenience.
type CategoryRepo = ports.CategoryRepository
