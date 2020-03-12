package main

import (
	"github.com/shuhrat-shokirov/di/pkg/di"
	"testing"
)

func TestNewMsg_NoDep(t *testing.T) {
	container := di.NewContainer()
	err := container.Provide(NewMsg)
	if err != nil {
		t.Fatalf("error must be nil: %v", err)
	}
}

func TestNewMsg_Err(t *testing.T) {
	container := di.NewContainer()
	err := container.Provide(NewMsg())
	if err == nil {
		t.Fatalf("error must not be nil: %v", err)
	}
}

func TestNewDependency_Dep_Msg(t *testing.T) {
	container := di.NewContainer()
	err := container.Provide(NewMsg, NewDependency)
	if err != nil {
		t.Fatalf("error must be nil: %v", err)
	}
}

func TestNewConsumer(t *testing.T) {
	container := di.NewContainer()
	err := container.Provide(NewMsg, NewDependency, NewConsumer)
	if err != nil {
		t.Fatalf("error must be nil: %v", err)
	}
}