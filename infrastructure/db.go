package infrastructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"path"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	// import _ "github.com/jinzhu/gorm/dialects/mssql"
	"os"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

// Opening a database and save the reference to `Database` struct.
func InitDb() *gorm.DB {

	dialect := os.Getenv("DB_DIALECT")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")

	var db *gorm.DB
	var err error
	if dialect == "sqlite3" {
		db, err = gorm.Open("sqlite3", path.Join(".", "app.db"))
	} else {
		// db, err := gorm.Open("mysql", "root:root@/go_api_blog_gonic?charset=utf8")
		databaseUrl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable ", host, username, password, dbName)
		db, err = gorm.Open(dialect, databaseUrl)
	}

	if err != nil {
		fmt.Println("db err: ", err)
		os.Exit(-1)
	}

	db.DB().SetMaxIdleConns(10)
	db.LogMode(true)
	DB = db
	return DB
}

// Delete the database after running testing cases.
func RemoveDb(db *gorm.DB) error {
	db.Close()
	err := os.Remove("./app.db")
	return err
}

// Using this function to get a connection, you can create your connection pool here.
func GetDB() *gorm.DB {
	return DB
}
