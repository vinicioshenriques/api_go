package dbconfig

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const PostgresDriver = "postgres"

const User = "postgres"

const Host = "localhost"

const Port = "5433"

const Password = "postgres"

const DbName = "ApiGo"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

func InitDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		Host, User, Password, DbName, Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
