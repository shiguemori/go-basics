/*
 * Copyright (c) $originalComment.match("Copyright \(c\) (\d+)", 1, "-")2022.
 */

package main

import "fmt"

func findAllHobbyists(hobby string, hobbies map[string][]string) []string {
	var hobbyists []string
	for name, h := range hobbies {
		for _, v := range h {
			if v == hobby {
				hobbyists = append(hobbyists, name)
			}
		}
	}
	return hobbyists
}

func main() {
	hobbies := map[string][]string {
		"Steve": []string{"Fashion", "Piano", "Reading"},
		"Patty": []string{"Drama", "Magic", "Pets"},
		"Chad": []string{"Puzzles", "Pets", "Yoga"},
	}

	fmt.Println(findAllHobbyists("Yoga", hobbies))
}
