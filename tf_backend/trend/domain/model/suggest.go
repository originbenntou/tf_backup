package model

type Suggest struct {
	Id          uint64 `db:"id"`
	SearchId    string `db:"search_id"`
	SuggestWord string `db:"suggest_word"`
	CreatedAt   string `db:"created_at"`
	UpdatedAt   string `db:"updated_at"`
	// mapping用
	ChildSuggest
}
