package utils

import "math/rand"

func Random(n int) []byte {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnm")

	result := make([]byte, n)

	for i := range result {

		result[i] = letters[rand.Intn(len(letters))]
	}

	return result

}
