package main

import (
	"github.com/bigmassa/gohttp-starter/internal/conf"
	"github.com/bigmassa/gohttp-starter/internal/models"
)

func init() {
	// wait for a db connection before we continue
	// useful when waiting for a docker container to become ready
	conf.WaitForDB()
}

func main() {
	// setup database
	db, _ := conf.ConnectDB()
	defer db.Close()
	models.SetDatabase(db)
	models.AutoMigrate()
}
