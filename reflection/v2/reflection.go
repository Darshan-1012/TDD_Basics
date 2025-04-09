package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	//Second
	for i := 0; i < val.NumField(); i++ { //NumFields is to find the number of params in a struct
		field := val.Field(i) // Field is the number of iterations in a struct
		/*
			//Third
			if field.Kind() == reflect.String { //.Kind() - Knowing the type of data
				fn(field.String())
			}
			//Fourth
			if field.Kind() == reflect.Struct{
				walk(field.Interface(), fn)
			}
		*/
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
