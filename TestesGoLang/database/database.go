package database

import (
	"database/sql"
	"fmt"
	"goprojects/utils"

	_ "github.com/mattn/go-sqlite3"
)

type personal struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewValues struct {
	Name     string
	Nickname string
	Email    string
	Password string
}

type DeleteValues struct {
	Email    string
	Password string
}

func GetData() []personal {
	var personalData []personal

	db, err := sql.Open("sqlite3", "./database/personalData")

	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT * FROM personalData;")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var data personal

		err = rows.Scan(&data.ID, &data.Name, &data.Nickname, &data.Email, &data.Password)

		if err != nil {
			panic(err)
		}

		personalData = append(personalData, data)

	}

	defer rows.Close()

	defer db.Close()

	return personalData

}

func AddNewValues(nv NewValues) {
	db, err := sql.Open("sqlite3", "./database/personalData")

	if err != nil {
		fmt.Println(err)
	}

	// Insert new values
	insertStatement := `
        INSERT INTO personalData (Name, Nickname, Email, Password)
        VALUES (?, ?, ?, ?);
    `

	_, err = db.Exec(insertStatement, nv.Name, nv.Nickname, nv.Email, nv.Password)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("New values inserted successfully %s", utils.GreenColor("✓\n"))

	defer db.Close()

}

func DeleteValue(dv DeleteValues) {
	db, err := sql.Open("sqlite3", "./database/personalData")

	if err != nil {
		fmt.Println(err)
	}

	// Delete values using e-mail and password
	deleteStatement := `
        DELETE FROM personalData
        WHERE email=? AND password=?;
    `

	_, err = db.Exec(deleteStatement, dv.Email, dv.Password)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Values deleted successfully %s", utils.GreenColor("✓\n"))

	defer db.Close()

}
