package configs

import (
	"auth_authorization/pkg/utils"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

var (
	ApplicationPort           int
	PostgresHost              string
	PostgresPort              string
	PostgresUsername          string
	PostgresDBName            string
	PostgresSSLMode           string
	JWTSecret                 string
	LocalTimezone             string
	DefaultFormatDate         string
	RedisHost                 string
	RedisPort                 string
	RedisPassword             string
	RedisDBNumber             string
	DefaultFormatDateResponse string
	RedisExpire               string
	PasswordSalt              string
	UserContext               string
	FileDirPath               string
	AllowExtension            string
	LocalPath                 string
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ApplicationPort = utils.GetenvInt("APPLICATION_PORT", 7070)
	PostgresHost = utils.GetEnv("POSTGRES_HOST")
	PostgresPort = utils.GetEnv("POSTGRES_PORT")
	PostgresUsername = utils.GetEnv("POSTGRES_USERNAME")
	PostgresDBName = utils.GetEnv("POSTGRES_DB_NAME")
	PostgresSSLMode = utils.GetEnv("POSTGRES_SSL_MODE")

	JWTSecret = utils.GetEnv("JWT_SECRET")
	LocalTimezone = utils.GetEnv("TIME_ZONE")
	DefaultFormatDate = utils.GetEnv("DEFAULT_FORMAT_DATE")
	RedisHost = utils.GetEnv("REDIS_HOST")
	RedisPort = utils.GetEnv("REDIS_PORT")
	RedisPassword = utils.GetEnv("REDIS_PASSWORD")
	RedisDBNumber = utils.GetEnv("REDIS_DB_NUMBER")
	DefaultFormatDateResponse = utils.GetEnv("DEFAULT_FORMAT_DATE_RESPONSE")
	RedisExpire = utils.GetEnv("REDIS_EXPIRE")
	PasswordSalt = utils.GetEnv("PASSWORD_SALT")
	UserContext = utils.GetEnv("USER_CONTEXT")
	FileDirPath = utils.GetEnv("FILE_DIR_PATH")
	AllowExtension = utils.GetEnv("ALLOW_EXTENSION")
	LocalPath = utils.GetEnv("LOCAL_PATH")

	if ApplicationPort == 0 && JWTSecret == "" {
		_ = fmt.Errorf("environment variable is not set")
	}
}
