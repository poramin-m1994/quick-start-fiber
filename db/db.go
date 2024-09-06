package db

import (
	"context"
	// "fmt"
	"log"
	// configs "quickstart/conf"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Connect() {
	var err error

	// กำหนดค่าการเชื่อมต่อฐานข้อมูล
	databaseUrl := "postgres://poramin:1234@localhost:5432/personal"

	// เชื่อมต่อกับฐานข้อมูล
	Pool, err = pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	log.Println("Connected to PostgreSQL!")
}
