package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type datas struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	// Open a connection to the SQLite database.
	db, err := sql.Open("sqlite3", "./TestesGoLang/database/personalData")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Execute the query.
	rows, err := db.Query("SELECT * FROM personalData;")
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
	// Iterate over the rows and print the results.
	for rows.Next() {
		var id int
		var name string
		var nickname string
		var email string
		var password string
		err = rows.Scan(&id, &name, &nickname, &email, &password)
		if err != nil {
			panic(err)
		}

		// var personalData = []datas{
		// 	{ID: id, Name: name, Nickname: nickname, Email: email, Password: password},
		// }
	}

	// fmt.Println(personalData)

	// Close the connection to the database.
	// db.Close()

	defer db.Close() // Close when done with it!

}
