package main

import (
	"errors"
	"fmt"
	"github.com/shuhrat-shokirov/di/pkg/di"
	"log"
)

type Msg string

func NewMsg() Msg {
	return "hello msg"
}

type dependency struct{
	value string
}

func NewDependency(message Msg) *dependency {
	log.Print("dependency created")
	return &dependency{string(message)}
}

func (d *dependency) Start() {
	log.Print("dependency started")
}

func (d *dependency) Stop() {
	log.Print("dependency stopped")
}

type consumer struct {
	dep *dependency
}

func NewConsumer(dep *dependency) *consumer {
	if dep == nil {
		log.Print(errors.New("dependency can't be nil"))
	}
	log.Print("consumer created")
	return &consumer{dep: dep}
}

func main() {
	{
		container := di.NewContainer()
		container.Provide(
			NewMsg,
			NewDependency,
		)
	}
	{
		container := di.NewContainer()
		container.Provide(
			NewMsg,
			NewConsumer,
			NewDependency,
		)

		var component *consumer
		container.Component(&component)
		fmt.Print(component.dep)
	}
}