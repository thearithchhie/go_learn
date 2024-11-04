package user

import "github.com/jmoiron/sqlx"

type UserService struct {
	r *UserRepo
}

func NewUserService(db *sqlx.DB) *UserService {
	return &UserService{
		r: NewUserRepo(db),
	}
}
