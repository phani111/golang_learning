package main

import (
	"6-interface_example_go/mysqldb"
	"fmt"
	"log"
	"os"
)

type dbcontract interface {
	Close()
	InsertUser(userName string) error
	SelectSingleUser(userName string) (string, error)
}

type Application struct {
	db dbcontract
}

func (this Application) Run() {
	userName := "user1"

	err := this.db.InsertUser(userName)
	if err != nil {
		log.Println("couldnt insert user: %s", userName)
	}

	user, err := this.db.SelectSingleUser(userName)
	if err != nil {
		log.Println("couldnt fetch user: %s", userName)
	}

	fmt.Println(user)
}

func NewApplication(db dbcontract) *Application {
	return &Application{db: db}
}

func main() {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	db, err := mysqldb.New(dbUser, dbPassword, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatalf("failed to intiate dbase connection: %v", err)
	}
	defer db.Close()

	app := NewApplication(db)

	app.Run()

}
