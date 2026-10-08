package utils

import "strings"

// IsNullOrEmpty checks if a string is null or empty
// Returns true if the string is null or empty
func IsNullOrEmpty(input *string) bool {
	return input == nil || (strings.Trim(*input, " ") == "")
}

// IsNotNullOrEmpty checks if a string is not null or empty
// Returns true if the string is not null or empty
func IsNotNullOrEmpty(input *string) bool {
	return !IsNullOrEmpty(input)
}
