package models

import (
    "context"
    "errors"
    "strings"
    "github.com/jackc/pgconn"
    pq "github.com/lib/pq"
    "github.com/mytheresa/go-hiring-challenge/domain"
    "gorm.io/gorm"
)

type CategoriesRepository struct {
    db *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) *CategoriesRepository {
    return &CategoriesRepository{db: db}
}

func (r *CategoriesRepository) List(ctx context.Context) ([]Category, error) {
    var items []Category
    if err := r.db.WithContext(ctx).Find(&items).Error; err != nil {
        return nil, err
    }
    return items, nil
}

func (r *CategoriesRepository) Create(ctx context.Context, c Category) (Category, error) {
    if err := r.db.WithContext(ctx).Create(&c).Error; err != nil {
        if isUniqueViolation(err) {
            return Category{}, domain.ErrAlreadyExists
        }
        return Category{}, err
    }
    return c, nil
}

// isUniqueViolation detects Postgres unique constraint errors across drivers and wrappers.
func isUniqueViolation(err error) bool {
    if errors.Is(err, gorm.ErrDuplicatedKey) {
        return true
    }
    var pgErr *pgconn.PgError
    if errors.As(err, &pgErr) && pgErr.Code == "23505" { // SQLSTATE unique_violation
        return true
    }
    var pqErr *pq.Error
    if errors.As(err, &pqErr) && string(pqErr.Code) == "23505" {
        return true
    }
    // Fallback: match common message when error wrapping prevents type assertion
    msg := err.Error()
    if strings.Contains(msg, "SQLSTATE 23505") || strings.Contains(strings.ToLower(msg), "duplicate key") {
        return true
    }
    return false
}
