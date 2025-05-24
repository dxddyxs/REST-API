package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		5432,
		"postgres",
		"1234",
		"postgres",
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexao: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco: %v", err)
	}

	fmt.Println("Conexao com o banco estabelecida com sucesso!")
	return db, nil
}
