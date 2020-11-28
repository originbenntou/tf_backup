package model

type History struct {
	Id        uint64 `db:"id"`
	UserUuid  string `db:"user_uuid"`
	SearchId  uint64 `db:"search_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
	// mapping用
	SearchWord string `db:"search_word"`
	Date       string `db:"date"`
	Status     uint64 `db:"status"`
}
