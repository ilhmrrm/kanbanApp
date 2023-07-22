package utils

import (
	"kanbanApp/entity"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() error {
	// connect using gorm pgx
	// conn, err := gorm.Open(postgres.New(postgres.Config{
	// 	DriverName: "pgx",
	// 	DSN:        os.Getenv("DATABASE_URL"),
	// }), &gorm.Config{})
	conn, err := gorm.Open(mysql.New(mysql.Config{DSN: os.Getenv("DATABASE_URL")}), &gorm.Config{})
	if err != nil {
		return err
	}

	conn.AutoMigrate(entity.User{}, entity.Category{}, entity.Task{})
	SetupDBConnection(conn)

	return nil
}

func SetupDBConnection(DB *gorm.DB) {
	db = DB
}

func GetDBConnection() *gorm.DB {
	return db
}
