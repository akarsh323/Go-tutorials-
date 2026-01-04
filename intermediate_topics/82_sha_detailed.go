package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

/*
================================================================================
                    HASHING & SHA IN GO - DETAILED GUIDE
================================================================================

WHAT IS HASHING?

Hashing is a process used in computing to transform data into a fixed-size
string of characters, which typically appears random. This transformation is
done using a special algorithm called a hash function.

================================================================================

BASIC CHARACTERISTICS OF HASHING

1. FIXED SIZE OUTPUT
   - No matter how large or small the input data is, the output hash will
     always be of fixed size.
   - Example: SHA256 always produces a 256-bit (32 byte) hash
   - Example: SHA512 always produces a 512-bit (64 byte) hash

2. DETERMINISTIC
   - The same input will always produce the same hash output.
   - If you hash the same data multiple times, you will always get the same result.

3. UNIQUE OUTPUT (AVALANCHE EFFECT)
   - Even a small change in the input will produce a completely different hash.
   - Example: "course" and "course" (one letter difference) produce completely
     different hashes - the output doesn't suggest they're similar.
   - This property is called the AVALANCHE EFFECT.

4. IRREVERSIBLE (ONE-WAY FUNCTION)
   - You cannot easily reverse a hash to get the original input data.
   - Hash functions are designed to be one-way functions.
   - Theoretical: impossible to get input from output
   - Practical: computationally infeasible with current technology

5. EFFICIENT
   - Hash functions are designed to be fast to compute.
   - This makes them practical for many applications.
   - Even low-powered computers can quickly compute hashes.

================================================================================

WHY DO WE USE HASHING?

1. PASSWORD STORAGE & VERIFICATION
   - Passwords are NOT stored directly in databases.
   - Password is converted to a hash, and THAT hash is stored.
   - During login: entered password → hashed → compared with stored hash
   - Even if database is breached, attacker gets hash (not original password)
   - Attacker cannot login with the hash (it gets hashed again, won't match)

2. DATA INTEGRITY VERIFICATION
   - Download a file and verify it hasn't been altered
   - Hash the downloaded file and compare with original hash from source
   - If hashes match: file is intact and unaltered
   - If hashes differ: file has been modified or corrupted
   - Example: Go installer downloads include hash verification

3. QUICK DATA RETRIEVAL
   - Used in hash tables and hash maps for O(1) lookup times

4. DIGITAL SIGNATURES & CERTIFICATES
   - Ensures authenticity and non-repudiation

================================================================================

WHY ARE HASHES IRREVERSIBLE?

Several factors make hashing irreversible:

1. INFORMATION LOSS
   - Hash functions produce fixed-size output regardless of input size
   - Information is lost during transformation
   - Impossible to recover original input from output

2. NO COLLISION DETECTION
   - Good hash functions are designed to minimize collisions
   - (Collision = different inputs produce same hash)
   - Even if collisions occur, finding two inputs that produce the same hash
     is computationally infeasible

3. AVALANCHE EFFECT
   - Small change in input produces drastically different hash
   - Output provides no useful clues about the input
   - Makes reversing practically impossible

4. MASSIVE NUMBER SPACE
   - SHA256 produces 2^256 possible hash values (2 multiplied 256 times)
   - To crack by brute force: try every possible input until finding one
     that produces desired hash
   - Time required increases exponentially with hash length
   - Impractical with current technology

PASSWORD VERIFICATION METHOD:
- We cannot reverse a hash
- We can verify if given input matches a hash by:
  1. Hash the input
  2. Compare result with stored hash
  3. If they match → input is correct
- This is the standard practice for password verification

================================================================================

SHA256 vs SHA512

SHA256 (Secure Hash Algorithm 256-bit):
- Produces 256-bit hash (32 bytes)
- Standard security level
- Good for most applications
- Faster than SHA512

SHA512 (Secure Hash Algorithm 512-bit):
- Produces 512-bit hash (64 bytes)
- Higher level of security
- Recommended for applications requiring stronger security
- Slower but more secure

================================================================================

HEX vs BYTE REPRESENTATION

When you hash data, the result is a byte slice. To make it human-readable,
you can convert it to hexadecimal (hex).

Hex values use characters 0-9 and a-f to represent each byte.
Example:
- Byte value 239 → Hex value "ef"
- Byte value 146 → Hex value "92"
- Byte value 183 → Hex value "b7"

================================================================================

PASSWORD SALTING

Salting adds an extra layer of security to password hashing.

WHAT IS SALT?
- A unique random value added to the password before hashing
- Can be randomly generated or a stored string
- Purpose: ensures different hashes even if users have same password

WHY USE SALT?

1. PROTECTS AGAINST DICTIONARY ATTACKS
   - Attacker cannot pre-compute hashes for common passwords

2. PROTECTS AGAINST RAINBOW TABLE ATTACKS
   - Rainbow tables are pre-computed hashes of common passwords
   - Salt makes these pre-computed tables useless

3. UNIQUE HASHES FOR SAME PASSWORD
   - User A: password "123" + salt "abc" → hash1
   - User B: password "123" + salt "xyz" → hash2
   - Same password, different salts, completely different hashes
   - Without salt, both would have identical hashes (security risk)

PASSWORD STORAGE WITH SALT:
Database columns:
- user_id: 1
- username: "john"
- salt: (base64 encoded random value)
- password_hash: (hash of password + salt, base64 encoded)

NOTE: Even the salt is encoded before storage (not stored as raw bytes)
This adds extra security - attacker doesn't know it's a salt or how to decode it

================================================================================

CRYPTOGRAPHICALLY SECURE RANDOM NUMBERS

When generating salt, use crypto/rand, NOT math/rand

DIFFERENCES:

Cryptographically Secure Random (crypto/rand):
- Generated using cryptographic algorithms designed to be secure
- Unpredictable and resistant to reverse engineering
- Even if part of data is known, predicting rest is computationally infeasible
- NOT seeded (uses OS entropy source)
- Use for: generating cryptographic keys, secure tokens, password salt,
           encryption keys, digital signatures

Pseudo-Random (math/rand):
- Generated using algorithms like Linear Congruential Generators (LCG) or
  Mersenne Twister
- Can be predicted if internal state is known
- Requires seeding
- Use for: simulations, games, non-security random data generation

KEY POINT: Always use crypto/rand for security-sensitive operations

================================================================================

io.ReadFull FUNCTION

Used to read cryptographically secure random numbers from crypto/rand.Reader

Signature:
  func ReadFull(r io.Reader, buf []byte) (n int, err error)

What it does:
- Reads from the io.Reader (crypto/rand.Reader in our case)
- Fills the byte slice with data from the reader
- Continues reading until the buffer is completely full
- Returns number of bytes read and any error

Example:
  salt := make([]byte, 16)
  _, err := io.ReadFull(rand.Reader, salt)
  // salt now contains 16 bytes of cryptographically secure random data

================================================================================
*/

// EXAMPLE 1: BASIC SHA256 HASHING
func basicSHA256Hashing() {
	fmt.Println("\n" + "="*80)
	fmt.Println("EXAMPLE 1: BASIC SHA256 HASHING")
	fmt.Println("=" * 80)

	password := "password123"

	// Hash using SHA256
	hash := sha256.Sum256([]byte(password))

	fmt.Println("Original password:", password)
	fmt.Println("Hash (byte slice):", hash[:])

	// Convert to hex for readable format
	hashHex := fmt.Sprintf("%x", hash)
	fmt.Println("Hash (hexadecimal):", hashHex)
}

// EXAMPLE 2: SHA256 vs SHA512
func sha256VsSha512() {
	fmt.Println("\n" + "="*80)
	fmt.Println("EXAMPLE 2: SHA256 vs SHA512")
	fmt.Println("=" * 80)

	password := "password123"

	// SHA256 hash
	hash256 := sha256.Sum256([]byte(password))
	hashHex256 := fmt.Sprintf("%x", hash256)
	fmt.Println("SHA256 Hash:", hashHex256)
	fmt.Println("SHA256 Byte slice:", hash256[:])

	// SHA512 hash
	hash512 := sha512.Sum512([]byte(password))
	hashHex512 := fmt.Sprintf("%x", hash512)
	fmt.Println("\nSHA512 Hash:", hashHex512)
	fmt.Println("SHA512 Byte slice:", hash512[:])

	fmt.Printf("\nSHA256 size: %d bytes\n", len(hash256))
	fmt.Printf("SHA512 size: %d bytes\n", len(hash512))
}

// EXAMPLE 3: DETERMINISTIC NATURE OF HASHING
func deterministicHashing() {
	fmt.Println("\n" + "="*80)
	fmt.Println("EXAMPLE 3: DETERMINISTIC NATURE")
	fmt.Println("=" * 80)

	password := "password123"

	// Hash same password multiple times
	hash1 := sha256.Sum256([]byte(password))
	hash2 := sha256.Sum256([]byte(password))
	hash3 := sha256.Sum256([]byte(password))

	hashHex1 := fmt.Sprintf("%x", hash1)
	hashHex2 := fmt.Sprintf("%x", hash2)
	hashHex3 := fmt.Sprintf("%x", hash3)

	fmt.Println("Hash 1:", hashHex1)
	fmt.Println("Hash 2:", hashHex2)
	fmt.Println("Hash 3:", hashHex3)
	fmt.Println("All hashes identical?", hash1 == hash2 && hash2 == hash3)
	fmt.Println("✓ Same input always produces same hash (deterministic)")
}

// EXAMPLE 4: AVALANCHE EFFECT
func avalancheEffect() {
	fmt.Println("\n" + "="*80)
	fmt.Println("EXAMPLE 4: AVALANCHE EFFECT (SMALL CHANGE = BIG DIFFERENCE)")
	fmt.Println("=" * 80)

	password1 := "course"
	password2 := "course" // Only one letter different!

	hash1 := sha256.Sum256([]byte(password1))
	hash2 := sha256.Sum256([]byte(password2))

	hashHex1 := fmt.Sprintf("%x", hash1)
	hashHex2 := fmt.Sprintf("%x", hash2)

	fmt.Printf("Input 1: '%s'\n", password1)
	fmt.Printf("Hash 1:  %s\n\n", hashHex1)

	fmt.Printf("Input 2: '%s'\n", password2)
	fmt.Printf("Hash 2:  %s\n\n", hashHex2)

	fmt.Println("Difference in input: just 1 letter ('o' vs 'u')")
	fmt.Println("Difference in hash:  completely different!")
	fmt.Println("✓ This is the AVALANCHE EFFECT")
}

// EXAMPLE 5: GENERATING SALT
func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// EXAMPLE 6: HASHING PASSWORD WITH SALT
func hashPasswordWithSalt(password string, salt []byte) string {
	// Combine salt and password
	saltedPassword := append(salt, []byte(password)...)

	// Hash the salted password
	hash := sha256.Sum256(saltedPassword)

	// Return as base64 encoded string
	return base64.StdEncoding.EncodeToString(hash[:])
}

// EXAMPLE 7: PASSWORD STORAGE AND VERIFICATION
func passwordStorageAndVerification() {
	fmt.Println("\n" + "="*80)
	fmt.Println("EXAMPLE 7: PASSWORD STORAGE & VERIFICATION WITH SALT")
	fmt.Println("=" * 80)

	// SIGNUP PHASE
	fmt.Println("\n--- SIGNUP PHASE ---")
	signupPassword := "password123"

	// Generate salt
	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	// Hash password with salt
	signupHash := hashPasswordWithSalt(signupPassword, salt)

	// Encode salt for storage
	saltStr := base64.StdEncoding.EncodeToString(salt)

	fmt.Println("Signup password:", signupPassword)
	fmt.Printf("Original salt (hex): %x\n", salt)
	fmt.Println("Encoded salt (base64):", saltStr)
	fmt.Println("Password hash (stored in DB):", signupHash)

	// WHAT GETS STORED IN DATABASE:
	fmt.Println("\n--- DATABASE STORAGE ---")
	fmt.Println("Stored salt:", saltStr)
	fmt.Println("Stored password_hash:", signupHash)

	// LOGIN PHASE
	fmt.Println("\n--- LOGIN PHASE: CORRECT PASSWORD ---")
	loginPassword := "password123"

	// Decode salt from database
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error decoding salt:", err)
		return
	}

	// Generate hash with login password
	loginHash := hashPasswordWithSalt(loginPassword, decodedSalt)

	fmt.Println("Login password:", loginPassword)
	fmt.Println("Generated login hash:", loginHash)
	fmt.Println("Stored password hash:", signupHash)

	if loginHash == signupHash {
		fmt.Println("✓ PASSWORD CORRECT - Login successful!")
	} else {
		fmt.Println("✗ Password incorrect - Login failed")
	}

	// LOGIN PHASE WITH WRONG PASSWORD
	fmt.Println("\n--- LOGIN PHASE: WRONG PASSWORD ---")
	wrongPassword := "password124" // One digit different

	wrongLoginHash := hashPasswordWithSalt(wrongPassword, decodedSalt)
	fmt.Println("Wrong password:", wrongPassword)
	fmt.Println("Generated hash from wrong password:", wrongLoginHash)
	fmt.Println("Stored password hash:", signupHash)

	if wrongLoginHash == signupHash {
		fmt.Println("✓ PASSWORD CORRECT - Login successful!")
	} else {
		fmt.Println("✗ PASSWORD INCORRECT - Login failed")
		fmt.Println("  (One digit change caused completely different hash)")
	}
}

// EXAMPLE 8: SALT PREVENTS IDENTICAL HASHES FOR SAME PASSWORD
func saltPreventsIdenticalHashes() {
	fmt.Println("\n" + "="*80)
	fmt.Println("EXAMPLE 8: SALT ENSURES DIFFERENT HASHES FOR SAME PASSWORD")
	fmt.Println("=" * 80)

	password := "password123"

	fmt.Println("Two users with same password: 'password123'\n")

	// User 1
	salt1, _ := generateSalt()
	hash1 := hashPasswordWithSalt(password, salt1)
	saltStr1 := base64.StdEncoding.EncodeToString(salt1)

	fmt.Println("User 1:")
	fmt.Println("  Salt:", saltStr1)
	fmt.Println("  Hash:", hash1)

	// User 2
	salt2, _ := generateSalt()
	hash2 := hashPasswordWithSalt(password, salt2)
	saltStr2 := base64.StdEncoding.EncodeToString(salt2)

	fmt.Println("\nUser 2:")
	fmt.Println("  Salt:", saltStr2)
	fmt.Println("  Hash:", hash2)

	fmt.Println("\n✓ SAME PASSWORD → DIFFERENT HASHES (due to different salts)")
	fmt.Println("✓ If there were no salt, both hashes would be identical!")
}

// EXAMPLE 9: COMPARING PASSWORD HASH WITH UNSALTED HASH
func saltingBenefit() {
	fmt.Println("\n" + "="*80)
	fmt.Println("EXAMPLE 9: BENEFIT OF SALTING")
	fmt.Println("=" * 80)

	password := "password123"

	// With salt
	salt, _ := generateSalt()
	saltedHash := hashPasswordWithSalt(password, salt)

	// Without salt (just plain hash)
	plainHash := sha256.Sum256([]byte(password))
	plainHashHex := fmt.Sprintf("%x", plainHash)
	plainHashB64 := base64.StdEncoding.EncodeToString(plainHash[:])

	fmt.Println("Password:", password)
	fmt.Println("\nHash WITHOUT salt:")
	fmt.Println("  ", plainHashHex)
	fmt.Println("  (base64):", plainHashB64)

	fmt.Println("\nHash WITH salt:")
	fmt.Println("  ", saltedHash)

	fmt.Println("\n✓ Different hashes, even though password is the same")
	fmt.Println("✓ Salt adds unpredictability to the hash")
}

// EXAMPLE 10: SECURITY CONSIDERATIONS
func securityNote() {
	fmt.Println("\n" + "="*80)
	fmt.Println("IMPORTANT: BASE64 IS NOT ENCRYPTION")
	fmt.Println("=" * 80)

	fmt.Println(`
Why we encode salt with Base64:
- Base64 encoding converts bytes to readable text format
- This is for STORAGE and TRANSMISSION, not security
- Anyone can decode Base64 (it's not encryption)

Example:
- Original salt (bytes):  [163 83 183 120 ...]
- Base64 encoded:         o1O3eJ... (readable text for database)
- Anyone with knowledge of Base64 can decode this back

Security comes from:
1. Using cryptographically secure random salt (unpredictable)
2. Unknown encoding/hashing algorithm (custom implementation)
3. Additional layers of security (encryption, secrets, etc.)

What encoding does:
- Makes binary data human-readable
- Makes it storable in text databases
- NOT a security measure by itself

Always use:
- crypto/rand for generating salt (not math/rand)
- Established cryptographic libraries
- Multiple layers of security for sensitive data
	`)
}

/*
================================================================================

KEY FUNCTIONS USED

1. sha256.Sum256(data []byte) -> [32]byte
   - Hashes data using SHA256
   - Returns fixed 32-byte hash

2. sha512.Sum512(data []byte) -> [64]byte
   - Hashes data using SHA512
   - Returns fixed 64-byte hash

3. io.ReadFull(r io.Reader, buf []byte) -> (n int, err error)
   - Reads exactly len(buf) bytes from reader
   - Used with rand.Reader for cryptographic randomness

4. base64.StdEncoding.EncodeToString(src []byte) -> string
   - Converts bytes to base64 string
   - Used for encoding salt before storage

5. base64.StdEncoding.DecodeString(s string) -> ([]byte, error)
   - Converts base64 string back to bytes
   - Used for retrieving salt from database

fmt.Sprintf("%x", hash) -> string
   - Converts byte array to hexadecimal string
   - Used for human-readable hash representation

================================================================================

SUMMARY

1. Hashing converts data to fixed-size, irreversible output
2. Deterministic: same input → same output
3. Avalanche effect: small input change → completely different output
4. Used for password storage, data integrity, security
5. SHA256 (256-bit) for standard security
6. SHA512 (512-bit) for higher security
7. Salt prevents identical hashes for same password
8. Use crypto/rand for secure random salt
9. Base64 encoding is for storage/transmission, not security
10. Always use established cryptographic libraries

================================================================================
*/

func main() {
	fmt.Println("\n" + "="*80)
	fmt.Println("                    HASHING AND SHA DETAILED")
	fmt.Println("=" * 80)

	// Run all examples
	basicSHA256Hashing()
	sha256VsSha512()
	deterministicHashing()
	avalancheEffect()
	passwordStorageAndVerification()
	saltPreventsIdenticalHashes()
	saltingBenefit()
	securityNote()

	fmt.Println("\n" + "="*80)
	fmt.Println("END OF EXAMPLES")
	fmt.Println("=" * 80)
}
