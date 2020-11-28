package model

type Company struct {
	Id        uint64 `db:"id"`
	Name      string `db:"name"`
	PlanId    uint64 `db:"plan_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
