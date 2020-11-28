package model

type ChildSuggest struct {
	Id               uint64 `db:"id"`
	SuggestId        uint64 `db:"suggest_id"`
	ChildSuggestWord string `db:"child_suggest_word"`
	Short            uint64 `db:"short"`
	Medium           uint64 `db:"medium"`
	Long             uint64 `db:"long"`
	ShortGraphs      string `db:"short_graphs"`
	MediumGraphs     string `db:"medium_graphs"`
	LongGraphs       string `db:"long_graphs"`
	CreatedAt        string `db:"created_at"`
	UpdatedAt        string `db:"updated_at"`
}
