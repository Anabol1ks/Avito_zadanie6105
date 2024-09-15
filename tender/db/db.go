package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func ConDB() {
	dsn := os.Getenv("POSTGRES_CONN")
	fmt.Println(dsn)
	var err error
	DB, err = sqlx.Open("postgres", dsn)
	if err != nil {
		panic(fmt.Sprintf("Ошибка подключения к бд: %v", err))
	} else {

		fmt.Println("Успешшное подключение")
	}
}
