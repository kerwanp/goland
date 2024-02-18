package main

import (
	"log"

	"github.com/kerwanp/goland/pkgs/di"
	"github.com/kerwanp/goland/pkgs/routing"
)

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (database *Database) FindUser(id int) string {
	return "User %s"
}

func main() {
	router := routing.NewRouter().
		// Provide services
		Service(di.T[Database], NewDatabase).
		// Arguments are autowired
		Route("/users/{id}", func(id int, db *Database) {
			user := db.FindUser(id),
				log.Print(user)
		})

	routing.Serve(router, ":8080")
}
