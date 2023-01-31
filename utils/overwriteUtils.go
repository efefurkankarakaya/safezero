package utils

import "math/rand"

// For performance reasons, keeping this array as file-static.
var characters []rune = []rune{
	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'r', 's', 't', 'u', 'v', 'y', 'z', 'q', 'x', 'w',
	'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
	'\'', '"', '!', '^', '+', '%', '&', '/', '(', ')', '=', '?', '_', ',', ';', '.', '*', '-', '<', '>', '#', '$', '{', '}', '[', ']', '|',
}

func getCharacters() []rune {
	// Defined above of the lines, as static.
	return characters
}

func GenerateRandomInteger(min int, max int) int {
	return rand.Intn(max-min) + min
}

func GetRandomCharacter() string {
	var characters []rune = getCharacters()
	var randomIndex int
	var randomCharacter string

	randomIndex = GenerateRandomInteger(0, len(characters))
	randomCharacter = string(characters[randomIndex])

	return randomCharacter
}
