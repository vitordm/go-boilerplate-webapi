package utils

import (
	"crypto/sha256"
	"fmt"
)

// Hash256 returns a sha256 hash of the given data.
func Hash256(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	bs := hash.Sum(nil)
	//return string(hash.Sum(nil))
	return fmt.Sprintf("%x", bs)
}

// Hash256StringToBytes returns a sha256 hash of the given data.
func Hash256StringToBytes(data string) []byte {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hash.Sum(nil)
}
