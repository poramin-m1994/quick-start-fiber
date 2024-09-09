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

	// // เชื่อมต่อกับฐานข้อมูล
	// Pool, err = pgxpool.New(context.Background(), databaseUrl)
	// if err != nil {
	// 	log.Fatalf("Unable to connect to database: %v\n", err)
	// }

	// log.Println("Connected to PostgreSQL!")

	// เชื่อมต่อกับฐานข้อมูล PostgreSQL โดยใช้ GORM
	DB, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// ทำการ AutoMigrate เพื่อสร้างตาราง users
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Connected to PostgreSQL and migrated schema for users!")

}
