package model

type Session struct {
	Id        uint64 `db:"id"`
	Token     string `db:"token"`
	UserUuid  string `db:"user_uuid"`
	CompanyId uint64 `db:"company_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
