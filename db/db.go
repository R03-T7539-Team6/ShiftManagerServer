package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"

	// Use PostgresSQL in gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init(isReset bool, models ...interface{}) {
	// err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	// db_name := os.Getenv("DB_NAME")
	// db_host := os.Getenv("DB_HOST")
	// db_port := os.Getenv("DB_PORT")
	// db_user := os.Getenv("DB_USER")
	// db_pass := os.Getenv("DB_PASS")
	// db_timezone := os.Getenv("DB_TIMEZONE")
	// db_sslmode := os.Getenv("DB_SSLMODE")

	// url := fmt.Sprintf("host=%s port=%s	dbname=%s user=%s password=%s",
	// 	db_host, db_port, db_name, db_user, db_pass)

	db, err = gorm.Open("postgres", databaseUrl)
	if err != nil {
		panic(err)
	}

	// テーブル名の複数形化を無効にする
	db.SingularTable(true)

	if isReset {
		db.DropTableIfExists(models)
	}
	db.AutoMigrate(models...)
}

// GetDB is called in models
func GetDB() *gorm.DB {
	return db
}

// Close is closing db
func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}
