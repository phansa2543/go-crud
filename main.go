package main

import (
	"github.com/phansa2543/go-crud/databases"
	"github.com/phansa2543/go-crud/servers"
)

func main() {
	db := databases.NewPostgresDB()
	servers.NewFiberServer(db).Start()
}