package di

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type container struct {
	components map[reflect.Type]interface{}
	definitions map[reflect.Type]definition
}

func NewContainer() *container {
	return &container{
		components:  make(map[reflect.Type]interface{}),
		definitions: make(map[reflect.Type]definition),
	}
}

func (c *container) Provide(constructors ...interface{})(err error) {
	err = c.register(constructors)
	if err != nil {
		return err
	}
	err = c.wire()
	if err != nil {
		return err
	}
	return nil
}

func (c *container) Component(target interface{}) {
	if target == nil {
		panic(errors.New("target cannot be nil"))
	}
	targetValue := reflect.ValueOf(target)
	targetType := targetValue.Type()
	targetTypeType := targetValue.Elem().Type()
	value, ok := c.components[targetTypeType]
	if !ok {
		panic(errors.New("no such component"),)
	}

	if targetType.Kind() != reflect.Ptr || targetValue.IsNil() {
		panic(errors.New("target must be a non-nil pointer"))
	}
	targetElemType := targetType.Elem()
	if !reflect.TypeOf(value).AssignableTo(targetElemType) {
		panic(errors.New("cant' assign component to pointer"))
	}
	targetValue.Elem().Set(reflect.ValueOf(value))
	return
}

func (c *container) Start() {
	for _, component := range c.components {
		if starter, ok := component.(StartListener); ok {
			starter.Start()
		}
	}
}

func (c *container) Stop() {
	for _, component := range c.components {
		if stopper, ok := component.(StopListener); ok {
			stopper.Stop()
		}
	}
}

func (c *container) register(constructors []interface{}) error {
	for _, constructor := range constructors {
		constructorType := reflect.TypeOf(constructor)
		if constructorType.Kind() != reflect.Func {
			return  fmt.Errorf("%s must be constructor", constructorType.Name())
		}
		if constructorType.NumOut() != 1 {
			return  fmt.Errorf("%s constructor must return only one result", constructorType.Name())
		}
		outType := constructorType.Out(0)
		if _, exists := c.definitions[outType]; exists {
			return  fmt.Errorf("ambiguous definition %s already exists", constructorType.Name())
		}
		paramsCount := constructorType.NumIn()
		c.definitions[outType] = definition{
			dependencies: paramsCount,
			constructor:  reflect.ValueOf(constructor),
		}
	}
	return nil
}

func (c *container) wire() error {
	rest := make(map[reflect.Type]definition, len(c.definitions))
	for key, value := range c.definitions {
		rest[key] = value
	}
	for {
		wired := 0
		for key, value := range rest {
			depsValues := make([]reflect.Value, 0)
			for i := 0; i < value.dependencies; i++ {
				depType := value.constructor.Type().In(i)
				if dep, exists := c.components[depType]; exists {
					depsValues = append(depsValues, reflect.ValueOf(dep))
				}
			}
			if len(depsValues) == value.dependencies {
				wired++
				component := value.constructor.Call(depsValues)[0].Interface()
				c.components[key] = component
				delete(rest, key)
				continue
			}
		}
		if len(rest) == 0 {
			return nil
		}
		if wired == 0 {
			for index, d := range rest {
				log.Printf("%d -dependency %v", index, d)
			}
			return fmt.Errorf("some components has unmet dependencies: %v", rest)
		}
	}
}

type definition struct {
	dependencies int
	constructor  reflect.Value
}
