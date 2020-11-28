package model

type RecoverSession struct {
	Id           uint64 `db:"id"`
	UserUuid     string `db:"user_uuid"`
	AuthKeyHash  []byte `db:"auth_key"`
	RecoverToken string `db:"recover_token"`
	CreatedAt    string `db:"created_at"`
	UpdatedAt    string `db:"updated_at"`
}
