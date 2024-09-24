package main

import "fmt"

func main() {
	fmt.Println(characterOccurrences("Hello"))
	fmt.Scanln()
}

// Function counts the occurrence of each character in a string.
// Returns a map where the keys are the characters and the values are their counts.
func characterOccurrences(inputString string) map[rune]int {
	// Initialize a new map to store the character counts
	counts := make(map[rune]int)

	// Iterate over each char in the string
	for _, char := range inputString {
		// Increment the count for respective character
		counts[char]++
	}

	// Return the map of char counts
	return counts
}
