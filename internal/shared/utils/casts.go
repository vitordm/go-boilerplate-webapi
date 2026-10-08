package utils

import "strconv"

func ParseInt(s string) int {
	c, _ := strconv.Atoi(s)
	return c
}

// ParseUint is a function that parses a string to an unsigned integer
func ParseUint(s string) uint64 {
	c, _ := strconv.ParseUint(s, 10, 64)
	return c
}

// Uint64ToString converts an uint64 to a string
func Uint64ToString(number uint64) string {
	return strconv.FormatUint(number, 10)
}

func ParseInt64(s string) int64 {
	c, _ := strconv.ParseInt(s, 10, 64)
	return c
}

func ParseInt32(s string) int32 {
	c, _ := strconv.ParseInt(s, 10, 32)
	return int32(c)
}

func ParseFloat64(s string) float64 {
	c, _ := strconv.ParseFloat(s, 64)
	return c
}
