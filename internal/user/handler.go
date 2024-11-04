package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type UserHandler struct {
	s  *UserService
	db *sqlx.DB
}

func NewUserHandler(db *sqlx.DB) *UserHandler {
	return &UserHandler{
		s:  NewUserService(db),
		db: db,
	}
}

func (h *UserHandler) Shows(c *fiber.Ctx) error {
	return nil
}
