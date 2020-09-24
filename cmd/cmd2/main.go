package main

import (
	"flag"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

func main() {
	flag.Parse()
}
