package routes

import (
	"auth_authorization/internal/user"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

func SetupRoute(app *fiber.App, db *sqlx.DB) {
	root := app.Group("/api/v1")
	//* User Route
	user.UserRoute(root, db)

}
