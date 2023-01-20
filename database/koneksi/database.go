package koneksi

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect() (db *sql.DB) {
	connStr := "user=postgres password=password dbname=db_test sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
	// defer db.Close()
	return

}

func Query(id int) (name string, age int) {

	db := Connect()

	err := db.QueryRow("SELECT name,age FROM users WHERE id = $1", id).Scan(&name, &age)

	if err != nil {
		log.Println("ERORRR DISINI")
		//panic(err)
	}

	return
}
