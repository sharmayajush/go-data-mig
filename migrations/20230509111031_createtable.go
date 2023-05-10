package migrations

import (
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreatetable, downCreatetable)
}

func upCreatetable(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	if _, err := tx.Exec(`
        CREATE TABLE table1 (
            id SERIAL,
            username TEXT,
            name TEXT,
            surname TEXT
        )
    `); err != nil {
		log.Println("1----")
		return err
	}

	// Insert data into the users table
	if _, err := tx.Exec(`
        INSERT INTO table1 (id, username, name, surname) VALUES
        (0, 'root', '', ''),
        (1, 'vojtechvitek', 'Vojtech', 'Vitek'),
		(2, 'abc', 'someone', 'good'),
		(3, 'pqr', 'some', 'good'),
		(4, 'xyz', 'someone', 'bad'),
		(6, 'abcd', 'nicey', 'good'),
		(7, 'abc', 'nicey', 'good')
    `); err != nil {
		log.Println("2----")
		return err
	}

	return nil
}

func downCreatetable(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	if _, err := tx.Exec(`DROP TABLE table1`); err != nil {
		return err
	}
	return nil
}
