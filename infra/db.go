package infra

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	connStr := "user=dadi1nho password=strongpass@ dbname=urlshortener sslmode=disable"

	return sql.Open("postgres", connStr)
}

func TestConnection() {
	db, err := Connect()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error on Ping:", err)
		return
	}

	fmt.Println("Conexão estabelecida com sucesso!")

}
