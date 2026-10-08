package utils

import (
	"crypto/rand"
	"encoding/base64"

	mathRand "math/rand"

	"github.com/oklog/ulid/v2"
)

// Ternary is a function that returns a value based on a condition
func Ternary[T any](condition bool, trueValue, falseValue T) T {
	if condition {
		return trueValue
	}
	return falseValue
}

func IsError(input interface{}) (error, bool) {
	newErr, ok := input.(error)
	return newErr, ok
}

func RandomString(length int) string {

	const charset = "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[mathRand.Intn(len(charset))]
	}

	result := string(b)

	return result
}

func Ulid() string {
	return ulid.MustNew(ulid.Now(), rand.Reader).String()
}

func PaginationOffsetLimit(page int, limit int) (int, int) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	if page == 1 {
		page = 0
		return page, limit
	}

	limitAux := limit
	limit = limit * page
	page = limitAux * (page - 1)

	return page, limit
}

func RandomInt(min int, max int) int {
	return mathRand.Intn(max-min) + min
}

// Pointer returns a pointer to the value passed as argument
// Please use this function with caution
// if you already have a variable YOU DON'T NEED THIS FUNCTION
func Pointer[T any](value T) *T {
	return &value
}

func DefaultInt64(value *int64, defaultValue int64) int64 {
	if value == nil {
		return defaultValue
	}
	return *value
}

func DefaultUInt64(value *uint64, defaultValue uint64) uint64 {
	if value == nil {
		return defaultValue
	}
	return *value
}

// DecodeBase64 it decodes a string Base64
func DecodeBase64(s string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(s)
}
