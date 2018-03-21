package main

// https://siongui.github.io/2016/01/09/go-sqlite-example-basic-usage/

import (
	"database/sql"
	"fmt"
	_"github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	db, err := sql.Open("sqlite3", "truc.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	fmt.Println("test")

/*sqlStmt := `
	create table foo (id integer not null primary key, name text);
	delete from foo;
	`*/

/*_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}*/

	fmt.Println(os.Args[1])

	var a string = os.Args[1]

	fmt.Println(a)
	/*b, err := os.Args[2]
	c, err := os.Args[3]*/

_, err = db.Exec("insert into foo(id, name) values(1, "%s"), (2,'a'), (3, 'a')", a)
	if err != nil {
		log.Fatal(err)
	}
}
