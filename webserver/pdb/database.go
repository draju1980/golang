package pdb

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// update the database details if its not hosted locally
const (
	host     = "localhost"
	port     = 6543
	user     = "postgres"
	password = "SuperPassword"
	dbname   = "postgres"
)

// function to connect to postgres database
func GetConnection() *sqlx.DB {
	psqlinfo := fmt.Sprintf("host=%s post=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sqlx.Open("postgres", psqlinfo)
	if err != nil {
		panic(err)
	}
	log.Println("DB Connection established.........")
	return db
}
