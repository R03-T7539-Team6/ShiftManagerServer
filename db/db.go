/*
ここではデータベースの基本設定のみ行う

autoMigrationについて
entityを追加するたびに、autoMigrationに追記すると
テーブルの作成を自動で行うことができる。
*/

package db

import (
	"github.com/R03-T7539-Team6/ShiftManagerSerer/entity"
	"github.com/jinzhu/gorm"

	// Use PostgresSQL in gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *gorm.DB
	err error
)

// Init is initialize db from main function
func Init() {

	// 現在は手元のWSL環境のPostgreにしている
	// curlを使うため。。。
	// [TODO]:後で直す。
	db, err = gorm.Open(
		"postgres",
		"host=172.29.48.1 port=5432 user=gorm dbname=gorm password=gorm sslmode=disable",
	)
	if err != nil {
		panic(err)
	}
	autoMigration()
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

// autoMigration is migration db
func autoMigration() {
	db.AutoMigrate(&entity.User{})
}
