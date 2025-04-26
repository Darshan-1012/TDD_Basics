package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	/* The first code*/
	// expected := "Darshan"
	// var got []string

	// // Anonymous struct
	// x := struct {
	// 	Name string
	// }{expected}

	// walk(x, func(input string) {
	// 	got = append(got, input)
	// })

	// if len(got) != 1 {
	// 	t.Errorf("wrong number of function calls %d and %d ", len(got), 1)
	// }

	// if got[0] != expected {
	// 	t.Errorf("got is %q and want is %q ", got[0], expected)
	// }

	/* The second case*/
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Darshan"},
			[]string{"Darshan"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}
}
