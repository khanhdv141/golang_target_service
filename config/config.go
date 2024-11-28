package config

import (
	//	Load .env file to environment
	"github.com/joho/godotenv"
	//	Map environment variables to struct
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Mongo struct {
		Host         string `envconfig:"MONGO_HOST"`
		Port         int    `envconfig:"MONGO_PORT"`
		User         string `envconfig:"MONGO_USER"`
		Password     string `envconfig:"MONGO_PASSWORD"`
		Database     string `envconfig:"MONGO_DB"`
		AuthDatabase string `envconfig:"MONGO_AUTH_DB"`
	}

	Mysql struct {
		Host            string `envconfig:"MYSQL_HOST"`
		Port            int    `envconfig:"MYSQL_PORT"`
		User            string `envconfig:"MYSQL_USER"`
		Password        string `envconfig:"MYSQL_PASSWORD"`
		Database        string `envconfig:"MYSQL_DATABASE"`
		MigrationFolder string `envconfig:"MYSQL_MIGRATION_FOLDER"`
	}

	Port int `envconfig:"PORT" default:"8080"`

	StorageDirectory string `envconfig:"STORAGE_DIRECTORY"`

	JWT struct {
		PublicKeyFilePath  string `envconfig:"JWT_PUBLIC_KEY_FILE_PATH"`
		PrivateKeyFilePath string `envconfig:"JWT_PRIVATE_KEY_FILE_PATH"`
	}

	Sentry struct {
		Dns string `envconfig:"SENTRY_DNS"`
	}
}

var ApplicationConfig *Config

func InitConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}
	ApplicationConfig = &Config{}
	return envconfig.Process("", ApplicationConfig)
}
