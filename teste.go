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
	var personalData []datas
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

	// Iterate over the rows and append them to personalData.
	for rows.Next() {
		var data datas // Create a new datas instance for each row
		err = rows.Scan(&data.ID, &data.Name, &data.Nickname, &data.Email, &data.Password)
		if err != nil {
			panic(err)
		}
		personalData = append(personalData, data) // Append the current row to the slice
	}

	defer rows.Close() // Close when done reading from table.

	fmt.Println(personalData)

	defer db.Close() // Close when done with it!

}
