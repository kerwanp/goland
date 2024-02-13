package di

import (
	"log"
	"testing"
)

type (
	Database interface {
		Fetch() []string
	}
	MySql struct{}
)

func NewMySql() MySql {
	return MySql{}
}

func (c MySql) Fetch() []string {
	return []string{"Tom", "Jane"}
}

type UserRepository struct {
	database Database
}

func NewUserRepository(database Database) UserRepository {
	return UserRepository{
		database: database,
	}
}

func (r *UserRepository) All() []string {
	return r.database.Fetch()
}

func TestContainer(t *testing.T) {
	container := NewContainer()
	Register[Database](container, NewMySql)
	Register[UserRepository](container, NewUserRepository)

	userRepository := Resolve[UserRepository](container)

	users := userRepository.All()

	log.Print(users) // Tom, Jane
}
