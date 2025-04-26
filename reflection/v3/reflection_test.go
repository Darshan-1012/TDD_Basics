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
		{ //Sixth
			"slices",
			[]Profile{
				{104, "Bengaluru"},
				{105, "Mysore"},
			},
			[]string{"Bengaluru", "Mysore"},
		},
		{ //seventh
			"arrays",
			[]Profile{
				{104, "Bengaluru"},
				{105, "Mysore"},
			},
			[]string{"Bengaluru", "Mysore"},
		},
		{ //Eight
			"maps",
			map[string]string{
				"Dog": "Barks",
				"Cat": "Meow",
			},
			[]string{"Barks", "Meow"},
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
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Dog": "Barks",
			"Cat": "Meow",
		}
		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})
		assertContains(t, got, "Barks")
		assertContains(t, got, "Meow")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{103, "Hyderabad"}
			aChannel <- Profile{104, "Chennai"}
			close(aChannel)
		}()
		var got []string
		want := []string{"Hyderabad", "Chennai"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{103, "Hyderabad"}, Profile{104, "Chennai"}
		}

		var got []string
		want := []string{"Hyderabad", "Chennai"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v ", got, want)
		}
	})
}

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but didn't", haystack, needle)
	}
}
