package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("Unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := "keyword not found"

		if err == nil {
			t.Fatal("Expected to get error")
		}
		assertStrings(t, err.Error(), want)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}

	key := "test"
	data := "this is just a test"
	dictionary.Add(key, data)

	assertData(t, dictionary, key, data)
}

func assertData(t testing.TB, dictionary Dictionary, key, data string) {
	t.Helper()

	got, err := dictionary.Search(key)
	if err != nil {
		t.Fatal("Should find added key", err)
	}
	assertStrings(t, got, data)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("%s is got %s is want", got, want)
	}
}
