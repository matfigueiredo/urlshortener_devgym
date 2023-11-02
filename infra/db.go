package infra

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/matfigueiredo/urlshortener_devgym/domain"
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

type URLRepositoryDB struct {
	db *sql.DB
}

func NewURLRepositoryDB() *URLRepositoryDB {
	db, _ := Connect() // Melhor manejar este erro em produção
	return &URLRepositoryDB{db: db}
}

func (r *URLRepositoryDB) Save(url domain.URL) error {
	_, err := r.db.Exec("INSERT INTO urls(original, code) VALUES($1, $2)", url.Original, url.Code)
	return err
}

func (r *URLRepositoryDB) FindByCode(code string) (domain.URL, error) {
	var url domain.URL
	err := r.db.QueryRow("SELECT original, code FROM urls WHERE code = $1", code).Scan(&url.Original, &url.Code)
	return url, err
}
