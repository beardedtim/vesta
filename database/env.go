package database

import (
	"mckp/vesta/env"
)

type DatabaseConfig struct {
	Port     int
	Host     string
	Username string
	Password string
	Database string
}

var Config = DatabaseConfig{
	Port:     env.GetEnvInt("DB_PORT"),
	Host:     env.GetEnvStr("DB_HOST"),
	Username: env.GetEnvStr("DB_USERNAME"),
	Password: env.GetEnvStr("DB_PASSWORD"),
	Database: env.GetEnvStr("DB_NAME"),
}
