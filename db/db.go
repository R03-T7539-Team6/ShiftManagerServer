/*
ここではデータベースの基本設定のみ行う

autoMigrationについて
entityを追加するたびに、autoMigrationに追記すると
テーブルの作成を自動で行うことができる。
*/

package db

import (
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

	db, err = gorm.Open(
		"postgres",
		"host=172.23.64.1 port=5432 user=gorm dbname=gorm password=gorm sslmode=disable",
	)
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
