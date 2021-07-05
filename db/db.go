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

/*************************************************
 *	specification;
 *	name 			= Init
 *	Function 	= Init database
 *	note			= init database and migration model from main fucntion
 *	date			= 07/05/2021
 *  author		= Yuma Matsuzaki
 *  History		= V1.00
 *  input 		= isReset : inittialize database flag
 * 						= models  : for migration models
 *  output    = none
 *  end of specification;
**************************************************/
func Init(isReset bool, models ...interface{}) {

	// 開発環境用
	// err := godotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// 	os.Exit(1)
	// }

	// db_name := os.Getenv("DB_NAME")
	// db_host := os.Getenv("DB_HOST")
	// db_port := os.Getenv("DB_PORT")
	// db_user := os.Getenv("DB_USER")
	// db_pass := os.Getenv("DB_PASS")
	// db_timezone := os.Getenv("DB_TIMEZONE")
	// db_sslmode := os.Getenv("DB_SSLMODE")

	// url := fmt.Sprintf("host=%s port=%s	dbname=%s user=%s password=%s",
	// 	db_host, db_port, db_name, db_user, db_pass)

	// 本番環境用（Heroku)
	databaseUrl := os.Getenv("DATABASE_URL")

	db, err = gorm.Open("postgres", databaseUrl)
	if err != nil {
		log.Fatal(err)
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
