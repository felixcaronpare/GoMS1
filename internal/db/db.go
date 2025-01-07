package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
  "gorm.io/driver/postgres"
)


func DbConn() *gorm.DB{
  dsn := GetDsnFromEnv()

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

  if err != nil {
    log.Fatalf("There was error connecting to the database: %v", err)
  }

  return db
}


//=========== UTILS ===========
func GetDsnFromEnv() string {
  err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

  host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

  dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=" + sslmode
  return dsn
}