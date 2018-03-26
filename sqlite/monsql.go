package main

import (
	"log"
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
	"strconv"
	"os"
)

var (
	dbname string	 = "mydb"
	id int 
	ipaddress string = os.Args[1]
	cookie string	 = os.Args[2]
	other string	 = os.Args[3]
	timestamp string = os.Args[4]	
)

func main() {

	if len(os.Args) != 3 {
		panic("argc")
		//log.Fatal("pas d'arg")
	} 

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

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS findmydns (id INTEGER PRIMARY KEY, ipaddress TEXT , cookie TEXT , other TEXT, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	statement, err = database.Prepare("INSERT INTO findmydns (ipaddress, cookie, other) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec(ipaddress, cookie , other )

	//rows, err := database.Query("SELECT id, IPADDRESS, COOKIENUMBER, OTHER, Timestamp FROM findmydns") 
	rows, err := database.Query("SELECT * FROM findmydns") 
	if err != nil {
		log.Fatal(err)
	} 

   for rows.Next() {
        rows.Scan(&id, &ipaddress, &cookie, &other, &timestamp)
        fmt.Println(strconv.Itoa(id) + ": " + ipaddress + " "+ cookie +" "+ other +" "+ timestamp)
    }
}



	

	
