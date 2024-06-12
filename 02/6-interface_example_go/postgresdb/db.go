package postgresdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
)

type Postgres struct {
	db *sql.DB
}

func New(user, password, host, port, dbName string) (*Postgres, error) {

	var err error

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=enabled", host, port, user, password, dbName)

	db, err := sql.Open(dbDriver, connStr)

	if err != nil {
		log.Printf("Postgres connection failure : %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Postgres ping failure : %v", err)
		return nil, err
	}

	return &Postgres{db: db}, nil

}

func (pgConn Postgres) Close() {
	err := pgConn.db.Close()
	if err != nil {
		log.Printf("Postgres close failure : %v", err)
	}
}

func (pgConn Postgres) InsertUser(userName string) error {
	pgConn.db.Exec("INSERT...")
	return nil
}

func (pgConn Postgres) SelectSingleUser(userName string) (string, error) {
	pgConn.db.Exec("SELECT...")
	return "user", nil
}
