package dmImpl

import (
	"cf-service/database"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

/*
@author galab pokharel
*/

type postgresConnection struct {
}

func NewPostgresConnection() database.Database {
	return &postgresConnection{}
}

const (
	host     = "127.0.0.1"
	user     = "postgres"
	password = "root"
	dbname   = "postgres"
	ssl      = "disable"
)

func (s postgresConnection) GetDatabase() (*sql.DB, error) {
	log.Println("Getting mysql connection")
	psqlInfo := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s", user, password, host, dbname, ssl)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
