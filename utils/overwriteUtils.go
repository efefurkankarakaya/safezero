package utils

import "math/rand"

// For performance reasons, keeping this array as file-static.
// const characters = []byte{
// 	'a', 'b', 'c', 'd', 'e', 'f', 'g', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'r', 's', 't', 'u', 'v', 'y', 'z', 'q', 'x', 'w',
// 	'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
// 	'\'', '"', '!', '^', '+', '%', '&', '/', '(', ')', '=', '?', '_', ',', ';', '.', '*', '-', '<', '>', '#', '$', '{', '}', '[', ']', '|',
// }

// For performance reasons, keeping this string as file-static.
const CHARACTERS string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890/\"!^+%&/()=?_,;.*-<>#${}[]|"

func getCharacters() string {
	// Defined above of the lines, as static.
	return CHARACTERS
}

func GenerateRandomInteger(min int64, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func GetRandomString(fileSize int64) string {
	// Length of characters string is 89.
	// (in 64-bit system) 89 = 000000...1011001, that means the lowest 7 bits represent 89 here and that's why we pick 7.
	const numberOfCharacterIndexBits = 7
	// Here is our mask is 1111111.
	const mask = (1 << numberOfCharacterIndexBits) - 1
	// 63 / 7 = 9 and that means, 9 times right shift (for 63 bit number) will equal to zero.

	var characters string = getCharacters()
	var data []byte = make([]byte, fileSize)
	var numberOfCharacters int64 = fileSize

	// Non-negative random 63 bits is being cached, after 63 / 7 = 9 times right shift cache will equal to 0.
	for i, cache := numberOfCharacters-1, rand.Int63(); i >= 0; {

		if cache == 0 {
			// Cache should be renew at this time, otherwise it makes random value is 0.
			cache = rand.Int63()
		}

		if randomIndex := int(cache & mask); randomIndex < len(characters) {
			data[i] = characters[randomIndex]
			i--
		}
		cache >>= numberOfCharacterIndexBits
	}

	return string(data)
}
