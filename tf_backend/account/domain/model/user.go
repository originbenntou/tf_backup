package model

type User struct {
	Id        uint64 `db:"id"`
	Uuid      string `db:"uuid"`
	Email     string `db:"email"`
	PassHash  []byte `db:"password"`
	Name      string `db:"name"`
	CompanyId uint64 `db:"company_id"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}
