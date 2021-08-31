/*
 * Copyright (c) 2021.
 */

package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test", "teste": "Isso e um teste"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")
		assertErrorDictionary(t, got, ErrNotFound)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		newDefinition := "new definition"
		dictionary = Dictionary{word: definition}
		err := dictionary.Update(word, newDefinition)
		assertErrorDictionary(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary = Dictionary{}
		err := dictionary.Update(word, definition)
		assertErrorDictionary(t, err, ErrWordDoesNotExist)
	})

	t.Run("count word", func(t *testing.T) {
		dictionary2 := Dictionary{"test": "this is just a test", "teste": "Isso e um teste"}
		qty := dictionary2.Count()
		assertInt(t, 2, qty)
	})

	t.Run("merge dictionary", func(t *testing.T) {
		dictionary2 := Dictionary{"test": "this is just a test", "teste": "Isso e um teste"}
		dictionary3 := Dictionary{"test3": "this is just another test", "teste4": "Isso e um outro teste"}
		newDictionary, err := dictionary2.Merge(dictionary3)
		qty := newDictionary.Count()
		assertErrorDictionary(t, err, nil)
		assertInt(t, 4, qty)
	})

}

func assertErrorDictionary(t testing.TB, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	newDefinition := "new definition"

	err := dictionary.Update(word, newDefinition)
	if err != nil {
		return
	}

	assertDefinition(t, dictionary, word, newDefinition)
}

func TestAddDictionary(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"

	err := dictionary.Add(word, definition)
	if err != nil {
		return
	}

	assertDefinition(t, dictionary, word, definition)
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertInt(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestDictionaryErr_Error(t *testing.T) {
	tests := []struct {
		name string
		e    DictionaryErr
		want string
	}{
		{
			name: "ErrNotFound",
			e:    ErrNotFound,
			want: "could not find the word you were looking for",
		},
		{
			name: "ErrWordExists",
			e:    ErrWordExists,
			want: "cannot add word because it already exists",
		},
		{
			name: "ErrWordDoesNotExist",
			e:    ErrWordDoesNotExist,
			want: "cannot update word because it does not exist",
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Error()
			if got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
