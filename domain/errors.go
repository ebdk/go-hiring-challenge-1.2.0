package domain

import "errors"

var (
    // ErrAlreadyExists indicates a unique constraint or conflict creating a resource.
    ErrAlreadyExists = errors.New("already exists")
    // ErrNotFound indicates an entity was not found.
    ErrNotFound    = errors.New("not found")
    // ErrInvalid indicates validation failed for a request.
    ErrInvalid     = errors.New("invalid")
)

