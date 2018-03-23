package main

import (
	"log"
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)

var dbname string = "mydb"

func main() {
	database, err := sql.Open("sqlite3", dbname)
	if err != nil {
		log.Fatal(err)
	}

	err2 := database.Ping()
	if err2 != nil {
		log.Fatal(err)
	} else {
		fmt.Println("base ouverte")
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS findmydns (id INTEGER PRIMARY KEY, IPADDRESS TEXT , COOKIENUMBER TEXT , OTHER TEXT, Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	statement, err = database.Prepare("INSERT INTO findmydns (IPADDRESS, COOKIENUMBER, OTHER) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec("10.1", "de8", "rien")




}
	

	
