package main

import (
	"auth_authorization/configs"
	"auth_authorization/configs/database"
	"auth_authorization/router"
	"auth_authorization/routes"
	"log"
	"strconv"

	"github.com/golang-module/carbon/v2"
	_ "github.com/lib/pq"
)

func main() {
	carbon.SetDefault(carbon.Default{
		Layout:       carbon.DateTimeLayout,
		Timezone:     carbon.Local,
		WeekStartsAt: carbon.Sunday,
	})
	configs.InitConfig()
	// Connection to postgres
	poolDB := database.GetDB()

	app := router.New()
	routes.SetupRoute(app, poolDB)
	log.Fatal(app.Listen(configs.PostgresHost + ":" + strconv.Itoa(configs.ApplicationPort)))
}
