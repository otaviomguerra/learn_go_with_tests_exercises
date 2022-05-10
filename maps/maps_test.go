package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, want, got)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is just a test")

		word := "test"
		definition := "this is just a test"
		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		err := dictionary.Add("test", "this is just a test")

		word := "test"
		definition := "this is just a test"
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("existing word/key", func(t *testing.T) {
		word := "test"
		updatedValue := "this is a updated key"

		dictionary := Dictionary{word: "this is just a test"}
		err := dictionary.Update(word, updatedValue)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, updatedValue)
	})

	t.Run("non existing word/key", func(t *testing.T) {
		word := "test"
		updatedValue := "this is a updated key"

		dictionary := Dictionary{}
		err := dictionary.Update(word, updatedValue)

		assertError(t, err, ErrNotFoundUpdate)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing word/key", func(t *testing.T) {
		word := "test"
		value := "this is a test"

		dictionary := Dictionary{word: value}
		dictionary.Delete(word)

		got, _ := dictionary.Search(word)
		if got != "" {
			t.Errorf("Expected key %s to be deleted", word)
		}
	})

}

func assertStrings(t testing.TB, want, got string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, want, got error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}

}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("Should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}
