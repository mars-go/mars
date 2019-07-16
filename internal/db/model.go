package db

import (
	"time"
)

type Model struct {
	ID        uint `db:"id;pk"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
