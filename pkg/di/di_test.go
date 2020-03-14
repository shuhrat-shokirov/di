package di

import (
	"testing"
)

type DSN string

func Test_container_Provide(t *testing.T) {
	container := NewContainer()
	err := container.Provide(
		func() DSN { return DSN("Hello") },
		func(dsn DSN) string { return "world" },
	)
	if err != nil {
		t.Errorf("error just be nil: %v", err)
	}

}

func Test_container_NoDependencies(t *testing.T) {
	container := NewContainer()
	err := container.Provide(
		func() DSN {return DSN("Hi")},)
	if err != nil {
		t.Errorf("error just be nil: %v", err)
	}
}

func TestNewMsg_Dep_Consumer(t *testing.T) {
	container := NewContainer()
	err := container.Provide(
		func() DSN{return DSN("why")},
		func(dsn DSN) string{return "what"})
	if err != nil {
		t.Fatalf("constructor must return only one result: %v", err)
	}
}

func TestNewMsg_Err_Consumer(t *testing.T) {
	container := NewContainer()
	err := container.Provide(
		func() DSN{return DSN("why")},
		func(func () DSN) string{return "what"})
	if err == nil {
		t.Fatalf("some components has unmet dependencies: %v", err)
	}
}

