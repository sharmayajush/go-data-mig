package migrations

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

type Users struct {
	id       int
	username string
	name     string
	surname  string
}

// Loop through rows using only one struct

func init() {
	goose.AddMigration(upDataMig, downDataMig)
}

func upDataMig(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	users := []Users{}
	rows, _ := tx.Query("SELECT * FROM table1")
	for rows.Next() {
		var r Users
		err := rows.Scan(&r.id, &r.username, &r.name, &r.surname)
		if err != nil {
			fmt.Println(err)
		}
		users = append(users, r)
	}
	fmt.Println("Data found :")
	fmt.Println(users)

	if _, err := tx.Exec(`
        CREATE TABLE table2 (
            username TEXT,
			place TEXT
        )
    `); err != nil {
		fmt.Println("1----")
		return err
	}
	stmt, err := tx.Prepare("INSERT INTO table2 (username, place) VALUES ($1, $2)")
	if err != nil {
		fmt.Println("2----")
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	place := ""
	for _, row := range users {
		if row.id < 4 {
			place = "Banglore"
		} else {
			place = "Delhi"
		}
		_, err := stmt.Exec(row.username, place)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func downDataMig(tx *sql.Tx) error {
	if _, err := tx.Exec(`DROP TABLE table2`); err != nil {
		return err
	}
	return nil
}
