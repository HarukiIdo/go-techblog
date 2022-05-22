package model

import "time"

// Article ...
type Article struct {
	ID        int       `db:"id" form:"id"`
	Title     string    `db:"title" form:"title"`
	Body      string    `db:"body" form:"body"`
	CreatedAt time.Time `db:"createdat"`
	UpdatedAt time.Time `db:"updatedat"`
}
