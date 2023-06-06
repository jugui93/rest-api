package database

import (
	"os"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/jugui93/rest-api/models"
)

type Database interface {
	Connect(dsn string)
}

type GormDatabase struct {
	Db *gorm.DB
}

var DB GormDatabase

func (g *GormDatabase) Connect(dsn string)  {
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&models.Fact{})

	DB = GormDatabase{
		Db: db,
	}
}

