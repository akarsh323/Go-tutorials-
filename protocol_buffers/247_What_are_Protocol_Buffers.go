package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"log"
)

/*
TOPIC: PROTOCOL BUFFERS (PROTOBUF) CONCEPT

INSTRUCTION:
Real Protobuf requires installing the 'protoc' compiler and generating .pb.go files.
To keep this lesson contained in ONE file, we are using Go's 'encoding/gob' (Binary)
to SIMULATE the efficiency of Protobuf compared to 'encoding/json' (Text).

THE CONCEPT:
1. JSON (Text): Human readable, but heavy. Sends field names ("name", "id") with every message.
2. PROTOBUF (Binary): Machine efficient. Sends numeric tags (1, 2) and packed bytes.
   It acts like a "Vacuum Seal" for your data.

THE WORKFLOW (Real World):
1. Define .proto file (Schema).
2. Compile with 'protoc' -> Generates code (structs).
3. Serialize data using generated code.
*/

// The Data Structure we want to send
type Person struct {
	ID    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ---------------------------------------------------------
// 1. JSON Simulation (The "Loose Trunk" Approach)
// ---------------------------------------------------------
// Represents standard text-based APIs.
func simulateJSON(p Person) []byte {
	fmt.Println("--- 1. JSON Serialization (Text) ---")

	// Marshaling converts the struct to JSON bytes
	data, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Payload: %s\n", string(data))
	fmt.Printf("Size:    %d bytes\n", len(data))
	return data
}

// ---------------------------------------------------------
// 2. Binary Simulation (The "Protobuf" Approach)
// ---------------------------------------------------------
// Represents binary formats like Protobuf.
// We use 'gob' here which is Go's built-in binary serializer.
// Real Protobuf is even more efficient/cross-language than this!
func simulateBinary(p Person) []byte {
	fmt.Println("\n--- 2. Binary/Protobuf Simulation (Vacuum Sealed) ---")

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	// Encoding converts the struct to a raw stream of bytes
	err := encoder.Encode(p)
	if err != nil {
		log.Fatal(err)
	}

	data := buffer.Bytes()

	// Binary data is not human readable, so we print hex/raw values
	fmt.Printf("Payload: %v\n", data)
	fmt.Printf("Size:    %d bytes\n", len(data))
	return data
}

// ---------------------------------------------------------
// 3. Application: Why It Matters
// ---------------------------------------------------------
func comparePerformance(jsonSize, binSize int) {
	fmt.Println("\n--- 3. The Comparison ---")

	diff := jsonSize - binSize
	percent := (float64(diff) / float64(jsonSize)) * 100

	fmt.Printf("JSON Size:   %d bytes\n", jsonSize)
	fmt.Printf("Binary Size: %d bytes\n", binSize)
	fmt.Printf("Savings:     %d bytes (%.2f%% smaller)\n", diff, percent)

	fmt.Println("\nNOTE: In massive systems (Microservices), saving 40-50% bandwidth")
	fmt.Println("translates to millions of dollars in infrastructure costs and faster apps.")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: PROTOCOL BUFFERS (BINARY SERIALIZATION)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	// Our Object
	user := Person{
		ID:    101,
		Name:  "Alice Wonderland",
		Email: "alice.w@example.com",
	}

	// 1. Run JSON (Text)
	jsonData := simulateJSON(user)

	// 2. Run Binary (Simulated Protobuf)
	binData := simulateBinary(user)

	// 3. Compare
	comparePerformance(len(jsonData), len(binData))

	fmt.Println("\n═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. PROTOBUF IS BINARY: It transmits raw bytes, not text strings.
   This makes it unreadable to humans but incredibly fast for computers.

2. SCHEMA REQUIRED: Unlike JSON, you need a blueprint (.proto file) beforehand.
   This ensures "Type Safety" (Service A and Service B agree on the format).

3. USE CASES:
   - Microservices (gRPC)
   - IoT / Mobile (Low Bandwidth)
   - High Performance Systems
	`)
}
