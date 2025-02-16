package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DB    *sql.DB
	Redis *redis.Client
)

func InitMySQL() {
	var err error
	dsn := "root:yourpassword@tcp(127.0.0.1:3306)/golang_assignment"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to MySQL: ", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("MySQL ping failed: ", err)
	}
	fmt.Println("Connected to MySQL")
}

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	fmt.Println("Connected to Redis")
}
