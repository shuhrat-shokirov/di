package main

import (
	"errors"
	"log"
	"reflect"
)

type myint int

func sample() {
	log.Print("called")
}

func sample2(msg string) {
	log.Print(msg)
}

func sample3(msg string) string {
	return msg
}

func callMe(c interface{}) (err error) {
	reflectType := reflect.TypeOf(c)
	if reflectType.Kind() != reflect.Func {
		err = errors.New("there should be error")
		return
	}

	if reflectType.NumIn() != 0 {
		err = errors.New("error with NumIn")
		return
	}

	reflectValue := reflect.ValueOf(c)
	reflectValue.Call(nil)
	return nil
}

func callMeAdvanced(c interface{}, args ...interface{}) (err error) {
	reflectType := reflect.TypeOf(c)
	if reflectType.Kind() != reflect.Func {
		err = errors.New("there should be error")
		return
	}

	reflectValue := reflect.ValueOf(c)
	valueArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		valueArgs[i] = reflect.ValueOf(arg)
	}

	reflectValue.Call(valueArgs)
	return nil
}

func callMeAdvancedWithResult(c interface{}, args... interface{}) (data interface{}, err error) {
	reflectType := reflect.TypeOf(c)
	if reflectType.Kind() != reflect.Func {
		err = errors.New("there should be error")
		return
	}

	reflectValue := reflect.ValueOf(c)
	valueArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		valueArgs[i] = reflect.ValueOf(arg)
	}

	resultValue := reflectValue.Call(valueArgs)[0]
	return resultValue.Interface(), nil
}

func main() {
	err := callMe(sample)
	if err != nil {
		log.Printf("can't call func: %v", err)
	}

	err = callMeAdvanced(sample2, "it works!")
	if err != nil {
		log.Printf("can't call func: %v", err)
	}

	result, err := callMeAdvancedWithResult(sample3,"Cosmos!!!")
	if err != nil {
		log.Printf("can't call func: %v", err)
	}

	log.Print(result)
}