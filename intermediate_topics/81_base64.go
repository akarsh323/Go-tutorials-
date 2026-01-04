package main

import (
	"encoding/base64"
	"fmt"
)

/*
================================================================================
                       BASE64 ENCODING IN GO
================================================================================

WHAT IS ENCODING?

Encoding is simply the process of translating data from one format to another.
It is used to ensure that data can be:

1. STORED CORRECTLY: When different systems use different storage formats.
2. TRANSMITTED SAFELY: Sending data over networks where certain characters might
   be misinterpreted.
3. INTEROPERABLE: Ensuring different software systems (APIs, databases) can read
   the same data.

Common examples include:
- ASCII: 7-bit text
- UTF-8: Variable-width Unicode
- URL Encoding: Converting spaces to %20

================================================================================

WHAT IS BASE64 ENCODING?

Base64 is a "Binary-to-Text" encoding scheme. It takes binary data (like an
image or a complex string) and converts it into a safe, readable text format
using a specific set of 64 characters.

THE CHARACTER SET:
- Uppercase letters: A-Z
- Lowercase letters: a-z
- Digits: 0-9
- Plus sign: +
- Forward slash: /
- Padding: = (equal sign, used when data doesn't fit perfectly)

WHY USE BASE64?

1. EMAIL: Email protocols were originally designed for text. Sending an image
   attachment requires converting that binary image data into Base64 text.

2. DATA URLS: You can embed small icons directly into HTML or CSS files by
   encoding the image as a Base64 string.

3. APIs: Sending binary data (like a PDF) inside a JSON payload (which is
   text-only) requires Base64.

================================================================================

IMPLEMENTATION IN GO

Go provides the "encoding/base64" package to handle this easily.

================================================================================
*/

// EXAMPLE 1: STANDARD ENCODING & DECODING
func standardEncodingDecoding() {
	fmt.Println("\n--- STANDARD ENCODING & DECODING ---")

	// 1. Define input as a byte slice
	data := []byte("Hello Base64 encoding")

	// 2. Encode to String
	// StdEncoding.EncodeToString takes bytes and returns a string
	encoded := base64.StdEncoding.EncodeToString(data)

	fmt.Println("Original data:", string(data))
	fmt.Println("Encoded:", encoded)
	// Output: SGVsbG8gQmFzZTY0IGVuY29kaW5n

	// 3. Decode back to bytes
	// StdEncoding.DecodeString returns ([]byte, error)
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	// 4. Convert bytes to string to read it
	fmt.Println("Decoded:", string(decodedBytes))
	// Output: Hello Base64 encoding

	// 5. Verify they match
	if string(decodedBytes) == string(data) {
		fmt.Println("✓ Encoding/Decoding successful!")
	}
}

// EXAMPLE 2: URL-SAFE ENCODING
func urlSafeEncoding() {
	fmt.Println("\n--- URL-SAFE ENCODING ---")

	// Standard Base64 uses + and / characters.
	// These characters are problematic in URLs:
	// - / looks like a directory separator
	// - + is used for spaces in URL parameters
	//
	// Solution: Use base64.URLEncoding which:
	// - Replaces + (plus) with - (hyphen)
	// - Replaces / (slash) with _ (underscore)

	// A string that would result in + or / when encoded
	data := []byte("Subject~?")

	// Standard Encoding (Might contain + or /)
	stdEncoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Data:", string(data))
	fmt.Println("Standard Encoding:", stdEncoded)
	// Output might contain: / character

	// URL Safe Encoding (Safe for web links)
	urlEncoded := base64.URLEncoding.EncodeToString(data)
	fmt.Println("URL Safe Encoding:", urlEncoded)
	// Output replaces / with _ and + with -

	// Decoding URL-safe encoded data
	decodedBytes, err := base64.URLEncoding.DecodeString(urlEncoded)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	fmt.Println("Decoded from URL Safe:", string(decodedBytes))

	// The decoded result is the same regardless of encoding variant
	fmt.Println("✓ URL-safe encoding/decoding successful!")
}

// EXAMPLE 3: PRACTICAL USE CASE - EMBEDDING DATA IN JSON
func jsonWithBase64() {
	fmt.Println("\n--- EMBEDDING BINARY DATA IN JSON ---")

	// Imagine you have binary data (like an image or PDF bytes)
	binaryData := []byte("This could be image data or binary content")

	// Encode to Base64 for safe transmission in JSON
	encoded := base64.StdEncoding.EncodeToString(binaryData)

	// Now you can safely put this in JSON (JSON only supports text)
	jsonPayload := fmt.Sprintf(`{"filename": "document.pdf", "data": "%s"}`, encoded)
	fmt.Println("JSON with Base64 embedded:")
	fmt.Println(jsonPayload)

	// When receiving this JSON, extract and decode
	// (In real code, you'd use json.Unmarshal)
	decodedBytes, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println("Decoded back from JSON:", string(decodedBytes))
}

// EXAMPLE 4: ENCODING DIFFERENT TYPES OF DATA
func encodingDifferentData() {
	fmt.Println("\n--- ENCODING DIFFERENT DATA TYPES ---")

	// Strings
	str := "Hello, World!"
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println("String:", str)
	fmt.Println("Encoded:", encoded)

	// Numbers (converted to bytes)
	num := "12345"
	encoded = base64.StdEncoding.EncodeToString([]byte(num))
	fmt.Println("\nNumber:", num)
	fmt.Println("Encoded:", encoded)

	// Special characters
	special := "!@#$%^&*()"
	encoded = base64.StdEncoding.EncodeToString([]byte(special))
	fmt.Println("\nSpecial characters:", special)
	fmt.Println("Encoded:", encoded)
}

// EXAMPLE 5: ERROR HANDLING IN DECODING
func errorHandlingInDecoding() {
	fmt.Println("\n--- ERROR HANDLING IN DECODING ---")

	// Valid Base64 string
	validBase64 := "SGVsbG8gV29ybGQ="

	decoded, err := base64.StdEncoding.DecodeString(validBase64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Valid Base64:", string(decoded))
	}

	// Invalid Base64 string (wrong length or invalid characters)
	invalidBase64 := "This is not valid base64!!!"

	decoded, err = base64.StdEncoding.DecodeString(invalidBase64)
	if err != nil {
		fmt.Println("\nInvalid Base64 detected:")
		fmt.Println("Error:", err)
		// Always check for errors when decoding untrusted data
	}
}

/*
================================================================================

IMPORTANT SECURITY NOTE

BASE64 IS NOT ENCRYPTION.

Encryption implies secrecy - you need a key to read the data.
Encoding is just a format change. Anyone can decode a Base64 string if they
have the standard look-up table.

DO NOT use Base64 alone to hide sensitive data like passwords!

If you need to protect sensitive data:
1. Use encryption (e.g., AES, RSA)
2. Use hashing (e.g., bcrypt for passwords)
3. Use secure protocols (e.g., HTTPS/TLS)

Base64 is for:
- Format compatibility
- Safe transmission of binary data
- Data that doesn't need secrecy

================================================================================

KEY FUNCTIONS FROM encoding/base64:

1. StdEncoding.EncodeToString(src []byte) -> string
   - Encodes bytes to Base64 string using standard alphabet

2. StdEncoding.DecodeString(s string) -> ([]byte, error)
   - Decodes Base64 string back to bytes
   - Returns error if the string is not valid Base64

3. URLEncoding.EncodeToString(src []byte) -> string
   - URL-safe Base64 encoding (- and _ instead of + and /)

4. URLEncoding.DecodeString(s string) -> ([]byte, error)
   - URL-safe Base64 decoding

================================================================================
*/

func main() {
	fmt.Println("================================================================================")
	fmt.Println("                    BASE64 ENCODING EXAMPLES IN GO")
	fmt.Println("================================================================================")

	// Run all examples
	standardEncodingDecoding()
	urlSafeEncoding()
	jsonWithBase64()
	encodingDifferentData()
	errorHandlingInDecoding()

	fmt.Println("\n================================================================================")
	fmt.Println("Key Takeaways:")
	fmt.Println("1. Base64 converts binary data to text format for safe transmission")
	fmt.Println("2. Use StdEncoding for most cases")
	fmt.Println("3. Use URLEncoding when the data goes into URLs")
	fmt.Println("4. Base64 is NOT encryption - anyone can decode it")
	fmt.Println("5. Always handle errors when decoding untrusted input")
	fmt.Println("================================================================================")
}
