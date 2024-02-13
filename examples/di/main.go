package main

import (
	"log"

	"github.com/kerwanp/goland/pkgs/di"
)

//------ Database implementation

type Database interface {
	Fetch() []string
}

type (
	MySql struct{}
)

func NewMySql() MySql {
	return MySql{}
}

func (c MySql) Fetch() []string {
	return []string{"Tom", "Jane"}
}

//------ UserRepository

type UserRepository struct {
	database Database
}

// Arguments of this function will be autowired
func NewUserRepository(database Database) UserRepository {
	return UserRepository{
		database: database,
	}
}

// Example of a method that access autowired dependency
func (r *UserRepository) All() []string {
	return r.database.Fetch()
}

//------ DI and autowiring example

func main() {
	container := di.NewContainer()

	di.Register[Database](container, NewMySql)                // Register Database implementation
	di.Register[UserRepository](container, NewUserRepository) // Register UserRepository

	userRepository := di.Resolve[UserRepository](container) // Resolve the UserRepository

	users := userRepository.All()

	log.Print(users) // Tom, Jane
}
