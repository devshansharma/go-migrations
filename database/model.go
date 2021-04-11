package database

import (
	"time"
)

// Model for common fields
type Model struct {
	ID         uint64 `migrator:"primaryKey"`
	CreatedAt  time.Time
	ModifiedAt time.Time
	DeletedAt  time.Time `migrator:"index"`
}

// Idb interface helps you in selecting database in production and testing.
type Idb interface {
	Query(string)
	Exec(string)
}
