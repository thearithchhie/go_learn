package router

import (
	"auth_authorization/configs"

	go_json "github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func New() *fiber.App {
	app := fiber.New(
		fiber.Config{
			// Change default encoding and decoding
			JSONEncoder: go_json.Marshal,
			JSONDecoder: go_json.Unmarshal,
			BodyLimit:   5 * 1024 * 1024, // Set limit to 5 MB, When user upload image size
		},
	)

	//logger
	app.Use(logger.New(logger.Config{
		TimeZone: configs.LocalTimezone,
	}))
	// i18n
	// app.Use(
	// 	fiberi18n.New(&fiberi18n.Config{
	// 		RootPath:        configs.LocalPath,
	// 		AcceptLanguages: []language.Tag{language.Chinese, language.English, language.Khmer},
	// 		DefaultLanguage: language.English,
	// 	}),
	// )
	// For handler panic, by default fiber not handle panic errors
	app.Use(recover.New())

	// cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, PUT, PATCH, POST, DELETE",
	}))

	// app.Use(jwt)
	// app.Use(reteLimit)
	return app

}
