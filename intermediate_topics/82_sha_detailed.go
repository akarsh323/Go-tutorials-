package intermediate

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

// Topic 82: SHA Hashing
// Computing SHA-1 and SHA-256 hashes

func main() {

	// SHA-256 hashing strings
	data := "Hello, World!"
	hash := sha256.Sum256([]byte(data))
	hashStr := hex.EncodeToString(hash[:])
	fmt.Println("SHA-256 of 'Hello, World!':")
	fmt.Println(hashStr)

	// SHA-1 hashing (older, less secure)
	sha1Hash := sha1.Sum([]byte(data))
	sha1Str := hex.EncodeToString(sha1Hash[:])
	fmt.Println("\nSHA-1 of 'Hello, World!':")
	fmt.Println(sha1Str)

	// Using hash.Hash interface for streaming
	h := sha256.New()
	io.WriteString(h, "Hello")
	io.WriteString(h, " ")
	io.WriteString(h, "World")
	hashStream := hex.EncodeToString(h.Sum(nil))
	fmt.Println("\nSHA-256 of 'Hello World' (streamed):")
	fmt.Println(hashStream)

	// Comparing hashes
	data1 := "secret"
	data2 := "secret"
	hash1 := sha256.Sum256([]byte(data1))
	hash2 := sha256.Sum256([]byte(data2))
	fmt.Printf("\nAre '%s' and '%s' the same? %v\n", data1, data2, hash1 == hash2)

	data3 := "different"
	hash3 := sha256.Sum256([]byte(data3))
	fmt.Printf("Are '%s' and '%s' the same? %v\n", data1, data3, hash1 == hash3)

	// Hash of longer content
	longData := "The quick brown fox jumps over the lazy dog"
	longHash := sha256.Sum256([]byte(longData))
	fmt.Println("\nSHA-256 of longer text:")
	fmt.Println(hex.EncodeToString(longHash[:]))

}
