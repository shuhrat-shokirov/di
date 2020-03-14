package di

import (
	"testing"
)

type DSN string

func Test_container_Provide(t *testing.T) {

	container := NewContainer()
	container.Provide(
		func() DSN { return DSN("Hello") },
		func(dsn DSN) string { return "world" },

	)

}

func Test_container_NoDependencies(t *testing.T) {
	container := NewContainer()
	container.Provide(
		func() DSN {return DSN("Hi")},)
}
