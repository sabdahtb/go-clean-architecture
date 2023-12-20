package bootstrap

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sabda.go/learn/domain"
)

func NewPostgreeDatabase(env *Env) *gorm.DB {
	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass
	dbName := env.DBName

	constring := "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai"
	dsn := fmt.Sprintf(constring, dbHost, dbUser, dbPass, dbName, dbPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&domain.UserModel{})
	db.AutoMigrate(&domain.Product{})

	return db
}
