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
	t.Run("New word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		data := "this is just a test"

		err := dictionary.Add(key, data)

		assertError(t, err, nil)
		assertData(t, dictionary, key, data)
	})

	t.Run("Existing word", func(t *testing.T) {
		key := "test"
		data := "this is just a test"
		dictionary := Dictionary{key: data}

		err := dictionary.Add(key, "new test")

		assertError(t, err, ErrWordExists)
		assertData(t, dictionary, key, data)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Existing word", func(t *testing.T) {
		key := "test"
		data := "this is just a test"
		dictionary := Dictionary{key: data}
		newDefinition := "new definition"

		err := dictionary.Update(key, newDefinition)
		assertError(t, err, nil)
		assertData(t, dictionary, key, newDefinition)
	})
	t.Run("New word", func(t *testing.T) {
		key := "test"
		data := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(key, data)
		assertError(t, err, ErrWordDoesNotExists)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Key exists", func(t *testing.T) {
		key := "test"
		dictionary := Dictionary{key: "test definition"}

		err := dictionary.Delete(key)
		assertError(t, err, nil)

		_, err = dictionary.Search(key)
		assertError(t, err, ErrNotFound)
	})
	t.Run("Key doesn't exists", func(t *testing.T) {
		key := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(key)
		assertError(t, err, ErrWordDoesNotExists)
	})
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
		t.Errorf("%q is got %q is want", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("%q is got %q is want", got, want)
	}
}
