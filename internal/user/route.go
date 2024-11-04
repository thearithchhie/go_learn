package user

import (
	"auth_authorization/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func UserRoute(root fiber.Router, db *sqlx.DB) {
	h := NewUserHandler(db)
	user := root.Group("/users")
	jwt := middleware.NewAuthMiddleware("secret")
	user.Get("/", jwt, h.Shows)

}
