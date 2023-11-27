package config

import (
	"fmt"
	"github.com/useSreel-developers/football-quiz-grpc/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "localhost"
	USER     = "postgres"
	PASSWORD = "aswassaw"
	DBNAME   = "football-quiz"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Jakarta", HOST, USER, PASSWORD, DBNAME)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	return db
}
