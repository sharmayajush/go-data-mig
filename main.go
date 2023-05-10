package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "mig/migrations"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var (
	flags    = flag.NewFlagSet("goose", flag.ExitOnError)
	dbstring = flags.String("dbstring", "postgresql://yajush:test123@localhost:5432/postgres?sslmode=disable", "connection string")
	dir      = flags.String("dir", "./migrations", "directory with migration files")
)

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[1]
	fmt.Println(command)

	db, err := goose.OpenDBWithDriver("postgres", *dbstring)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 2 {
		arguments = append(arguments, args[2:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
