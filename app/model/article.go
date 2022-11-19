package model

// Article ...
type Article struct {
	ID    int    `db:"id"`
	Title string `db:"title"`
}