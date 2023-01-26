package koneksi

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect() (db *sql.DB) {
	connStr := "user=postgres password=password dbname=db_mahasiswa sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Println(err)
	}
	return
}
