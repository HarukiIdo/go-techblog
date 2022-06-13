package model

import "time"

// Article expresses ...
// _form.htmlのname属性と構造体フィールドを紐付ける
type Article struct {
	ID        int       `db:"id" form:"id"`
	Title     string    `db:"title" form:"title"`
	Body      string    `db:"body" form:"body"`
	CreatedAt time.Time `db:"createdat"`
	UpdatedAt time.Time `db:"updatedat"`
}
