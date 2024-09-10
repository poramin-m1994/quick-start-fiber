package db

import (

	// "fmt"
	"log"
	"quickstart/models"

	// configs "quickstart/conf"

	"github.com/jackc/pgx/v5/pgxpool"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Pool *pgxpool.Pool
	DB   *gorm.DB
)

func Connect() {
	var err error

	// กำหนดค่าการเชื่อมต่อฐานข้อมูล
	databaseUrl := "postgres://poramin:1234@localhost:5432/personal"
	DB, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// ทำการ AutoMigrate เพื่อสร้างตาราง users
	err = DB.AutoMigrate(&models.User{}, &models.Token{}, &models.Stock{}, &models.Portfolio{}, &models.PurchaseHistory{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("Connected to PostgreSQL and migrated schema succ!")

	// // ทำการ Migration สำหรับโมเดลต่างๆ
	// if err := models.MigrateUser(); err != nil {
	// 	log.Fatalf("Failed to migrate User model: %v", err)
	// }

	// if err := models.MigrateToken(); err != nil {
	// 	log.Fatalf("Failed to migrate Token model: %v", err)
	// }

}
