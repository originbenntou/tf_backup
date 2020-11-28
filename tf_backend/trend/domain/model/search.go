package model

type Search struct {
	Id         uint64 `db:"id"`
	SearchWord string `db:"search_word"`
	Date       string `db:"date"`
	Status     uint64 `db:"status"`
	CreatedAt  string `db:"created_at"`
	UpdatedAt  string `db:"updated_at"`
}
