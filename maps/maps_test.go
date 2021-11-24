package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("existing key", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("non existing key", func(t *testing.T) {
		_, err := dictionary.Search("unknown")
		want := ErrKeyNotFound
		assertError(t, err, want)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is just a test"
		err := dictionary.Add(key, value)
		assertError(t, err, nil)
		assertDictionaryValue(t, dictionary, key, value)
	})

	t.Run("existing key", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		key := "test"
		value := "this is just a test"
		err := dictionary.Add(key, value)
		assertError(t, err, ErrKeyExists)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("new key", func(t *testing.T) {
		key := "test"
		dictionary := Dictionary{key: "this is just a test"}
		newKey := "new key"
		newValue := "new value"
		err := dictionary.Update(newKey, newValue)
		assertError(t, err, ErrKeyNotFound)
	})

	t.Run("existing key", func(t *testing.T) {
		key := "test"
		dictionary := Dictionary{key: "this is just a test"}
		newValue := "new value"
		dictionary.Update(key, newValue)
		assertDictionaryValue(t, dictionary, key, newValue)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	dictionary := Dictionary{key: "this is just a test"}

	dictionary.Delete(key)

	_, err := dictionary.Search(key)

	assertError(t, err, ErrKeyNotFound)
}

func assertDictionaryValue(t testing.TB, dictionary Dictionary, key, want string) {
	t.Helper()
	got, err := dictionary.Search(key)

	if err != nil {
		t.Fatal("should find added key:", err)
	}
	assertStrings(t, got, want)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if want == nil {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	} else {
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}
	}
}
