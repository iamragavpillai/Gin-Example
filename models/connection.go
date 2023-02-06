package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func ConnectDB() *gorm.DB {
	dbhost := os.Getenv(string("DB_HOST"))
	dbuser := os.Getenv(string("DB_USER"))
	dbpassword := os.Getenv(string("DB_PASS"))
	dbname := os.Getenv(string("DB_NAME"))
	dbport := os.Getenv(string("DB_PORT"))

	dsn := "postgres://" + dbuser + ":" + dbpassword + "@" + dbhost + ":" + dbport + "/" + dbname

	var database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	database.AutoMigrate(&Book{})

	return database

}
