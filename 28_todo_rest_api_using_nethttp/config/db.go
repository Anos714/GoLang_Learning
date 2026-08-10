package config

import (
	"fmt"
	"log"
	"os"
	"todo-rest-api/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase(){
	err := godotenv.Load()
		if err != nil {
			log.Println("Warning: No .env file found, relying on system environment variables")
		}

		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			log.Fatal("DATABASE_URL environment variable is missing!")
		}

		db,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})
		if err!=nil{
			log.Fatalf("Failed to connect to database: %v\n",err)
		}

		fmt.Println("Connected to Neon Postgres securely!")

		// 4. Run automatic migrations
		err = db.AutoMigrate(&models.Todo{})
		if err != nil {
			log.Fatalf("Migration failed: %v\n",err)
		}

		fmt.Println("Database migration done")
		DB=db

}
