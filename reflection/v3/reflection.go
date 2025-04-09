package main

import "reflect"

func walk(x interface{}, fn func(input string)) {
	/*
		val := reflect.ValueOf(x)

		if val.Kind() == reflect.Pointer { // Fifth
			val = val.Elem()
		}
	*/
	/*
			if val.Kind() == reflect.Slice { //Sixth
				for i := 0; i < val.Len(); i++ {
					walk(val.Index(i).Interface(), fn)
				}
				return
			}

		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			switch field.Kind() {
			case reflect.String:
				fn(field.String())
			case reflect.Struct:
				walk(field.Interface(), fn)
			}
		}
	*/
	val := getValue(x)

	// numberOfValue := 0
	// var getField func(int) reflect.Value

	walkValue := func(value reflect.Value) {
		walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String, reflect.Array:
		fn(val.String())
	case reflect.Struct:
		// numberOfValue = val.NumField()
		// getField = val.Field
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice:
		// numberOfValue = val.Len()
		// getField = val.Index
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		// for {
		// 	if v, ok := val.Recv(); ok {
		// 		walkValue(v)
		// 	} else {
		// 		break
		// 	}
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walkValue(v)
		}

	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}

	// for i := 0; i < numberOfValue; i++ {
	// 	walk(getField(i).Interface(), fn)
	// }
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
