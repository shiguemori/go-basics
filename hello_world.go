/*
 * Copyright (c) 2021.
 */

package main

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const portugueseHelloPrefix = "Ola, "
const defaultName = "World"

// This functions defines the language used to do the greetings.
// It receives a string that can be (Portuguese, Spanish and English).
// If something is sent different of those options the greetings will be in english
func greetingPrefix(language string) (prefix string) {
	switch language {
	case "Portuguese":
		prefix = portugueseHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	case "English":
	default:
		prefix = englishHelloPrefix
	}
	return
}

// Hello This functions greetings according to the
// parameters that receives (name string, language string)
// if none of those parameters was defined the return will be "Hello, World"
func Hello(name, language string) string {
	if name == "" {
		name = defaultName
	}
	return greetingPrefix(language) + name
}
