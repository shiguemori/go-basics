/*
 * Copyright (c) 2021.
 */

package main

type Dictionary map[string]string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

// This function return the error
func (e DictionaryErr) Error() string {
	return string(e)
}

// This function add a new word and definition to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

// This function will search the word
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// This function will update the word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

// This function will delte the word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

// This function will count how many words has in the dictionary
func (d Dictionary) Count() int {
	return len(d)
}

// This function will merge 2 dictionarys
func (d Dictionary) Merge(dictionary Dictionary) (Dictionary, error) {
	newDict := Dictionary{}
	for k, v := range d {
		err := newDict.Add(k, v)
		if err != nil {
			return newDict, err
		}
	}
	for k, v := range dictionary {
		err := newDict.Add(k, v)
		if err != nil {
			return newDict, err
		}
	}
	return newDict, nil
}
