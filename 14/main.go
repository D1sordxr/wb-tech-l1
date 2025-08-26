package main

import (
	"fmt"
	"reflect"
)

func main() {
	checkType(114)
	checkType("l1.14")
	checkType(true)
	checkType(make(chan string))
	checkType(make(chan int))
	checkType(nil)

	checkTypeReflect(114)
	checkTypeReflect(1.14)
	checkTypeReflect("l1.14")
	checkTypeReflect(true)
	checkTypeReflect(make(chan string))
	checkTypeReflect(make(chan int))
	checkTypeReflect(nil)
}

func checkType(v any) {
	if v == nil {
		fmt.Println("nil")
		return
	}

	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			fmt.Println("chan")
		} else {
			fmt.Println("not supported")
		}
	}
}
func checkTypeReflect(v interface{}) {
	if v == nil {
		fmt.Println("nil")
		return
	}

	switch reflect.TypeOf(v).Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	case reflect.Bool:
		fmt.Println("bool")
	case reflect.Float32, reflect.Float64:
		fmt.Println("float")
	case reflect.Chan:
		fmt.Println("chan")
	default:
		fmt.Println("not supported")
	}
}
