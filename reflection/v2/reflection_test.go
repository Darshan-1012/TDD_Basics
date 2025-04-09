package main

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{ //Second
			"struct with two string fields",
			struct {
				Name string
				City string
			}{"Chris", "London"},
			[]string{"Chris", "London"},
		},
		{ //Third
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Darshan", 104},
			[]string{"Darshan"},
		},
		{ //Fourth
			"nested fields",
			Person{
				"Darshan",
				Profile{104, "Bengaluru"},
			},
			[]string{"Darshan", "Bengaluru"},
		},
		{ //Fifth
			"pointer to things",
			&Person{
				"Darshan",
				Profile{104, "Bengaluru"},
			},
			[]string{"Darshan", "Bengaluru"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}
