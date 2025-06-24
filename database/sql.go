package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // importa o driver MySQL
)

var DB *sql.DB

func ConnectionDB() {
	var err error

	DB, err = sql.Open("mysql", "root:savio2002@tcp(localhost:3306)/tendasclub")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
}