package main

import (
	"CMS/config"
	"CMS/log"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	applicationConfig := config.ApplicationConfig
	databaseURL := fmt.Sprintf("mysql://%v:%v@tcp(%v:%v)/%v",
		applicationConfig.Mysql.User, applicationConfig.Mysql.Password,
		applicationConfig.Mysql.Host, applicationConfig.Mysql.Port,
		applicationConfig.Mysql.Database,
	)
	m, err := migrate.New(fmt.Sprintf("file://%v", applicationConfig.Mysql.MigrationFolder), databaseURL)
	if err != nil {
		panic(err)
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Panic(err)
	}
	msg := "migrate success"
	if errors.Is(err, migrate.ErrNoChange) {
		msg += ", but no changes"
	}
	log.Info(msg)
}
