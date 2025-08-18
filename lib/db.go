package lib

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func Connection()(*sqlx.DB,error){
	 err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found, reading from environment variables")
    }

   
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        user, password, host, port, dbname)

	db, err := sqlx.Connect("mysql",dsn)
	if err != nil {
		return nil,err
	}

	err = db.Ping()
    if err != nil {
        fmt.Println("Koneksi gagal:", err)
    }

	return db,nil
}