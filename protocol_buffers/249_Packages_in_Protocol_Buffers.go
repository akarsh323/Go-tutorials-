package main

import (
	"fmt"
)

/*
TOPIC: PROTOBUF PACKAGES & IMPORTS

CONCEPT:
As projects grow, you will have hundreds of .proto files.
If two files define a message named "User", the compiler will panic (Naming Collision).

THE SOLUTION:
We use the 'package' keyword to create Namespaces.
- website.User
- admin.User
This allows the compiler (and humans) to distinguish between them.
*/

// ---------------------------------------------------------
// Example 1: The Namespace Solution
// ---------------------------------------------------------
// Demonstrating how to define packages in .proto files to avoid conflicts.
func example1_PackageDefinitions() {
	fmt.Println("--- Example 1: Defining Packages (Namespaces) ---")

	// File 1: Public Website User
	publicUserProto := `
// FileName: public_user.proto
syntax = "proto3";

// 1. DECLARE PACKAGE
// This puts 'User' inside the 'website' namespace.
package website;

// Ideally, also specify the Go package path:
option go_package = "example.com/project/website";

message User {
  string username = 1;
  string email = 2;
}
`
	// File 2: Internal Admin User
	adminUserProto := `
// FileName: admin_user.proto
syntax = "proto3";

// 1. DECLARE PACKAGE
// This puts 'User' inside the 'admin' namespace.
package admin;

option go_package = "example.com/project/admin";

message User {
  int32 admin_id = 1;
  string clearance_level = 2;
}
`

	fmt.Println("File A (Public):")
	fmt.Println(publicUserProto)
	fmt.Println("\nFile B (Admin):")
	fmt.Println(adminUserProto)

	fmt.Println("\nRESULT:")
	fmt.Println("The compiler sees 'website.User' and 'admin.User'.")
	fmt.Println("They are now distinct types and can coexist in the same project.")
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 2: Importing Packages
// ---------------------------------------------------------
// Demonstrating how one .proto file can reuse messages from another.
func example2_Imports() {
	fmt.Println("--- Example 2: Importing (Modularity) ---")

	// The Scenario: A 'Company' message needs to store a list of 'Employees'.
	// But 'Employee' is defined in a different file (common/person.proto).

	companyProto := `
// FileName: company.proto
syntax = "proto3";

package enterprise;

// 1. IMPORT STATEMENT
// We look for the file path relative to the project root.
import "common/person.proto";

message Company {
  string company_name = 1;

  // 2. USE THE TYPE
  // We refer to the imported message by its PACKAGE name (common),
  // not the file name.
  repeated common.Person employees = 2;
}
`
	fmt.Println(companyProto)
	fmt.Println("\nKEY TAKEAWAY:")
	fmt.Println("We define 'Person' once in 'common/person.proto',")
	fmt.Println("and we can import it into Company, Payroll, and Shipping files.")
	fmt.Println("----------------------------------------\n")
}

// ---------------------------------------------------------
// Example 3: Generated Code Simulation (Go)
// ---------------------------------------------------------
// What does the generated Go code actually look like?

// SIMULATION: Generated from public_user.proto
// This would be in a folder named "website"
type WebsiteUser struct {
	Username string
	Email    string
}

// SIMULATION: Generated from admin_user.proto
// This would be in a folder named "admin"
type AdminUser struct {
	AdminID        int32
	ClearanceLevel string
}

func example3_GoSimulation() {
	fmt.Println("--- Example 3: Generated Code Simulation ---")

	// In real Go code, you would import these packages:
	// import "example.com/project/website"
	// import "example.com/project/admin"

	// Usage
	webUser := WebsiteUser{Username: "alice", Email: "alice@example.com"}
	adminUser := AdminUser{AdminID: 999, ClearanceLevel: "Top Secret"}

	fmt.Printf("Website User: %+v\n", webUser)
	fmt.Printf("Admin User:   %+v\n", adminUser)

	fmt.Println("\nNote: In Go, the Protobuf 'package' usually determines")
	fmt.Println("the Go generated package name (unless 'option go_package' is used).")
	fmt.Println("----------------------------------------\n")
}

func main() {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("TOPIC: PROTOBUF ORGANIZATION (PACKAGES & IMPORTS)")
	fmt.Println("═══════════════════════════════════════════════════════════\n")

	example1_PackageDefinitions()
	example2_Imports()
	example3_GoSimulation()

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("KEY TAKEAWAYS")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println(`
1. NAMESPACES: Always define 'package my.name;' at the top of .proto files.
   This prevents "Duplicate Name" errors.

2. CONVENTION: Package names should be lowercase (e.g., 'billing', 'shipping').

3. IMPORTS: Use 'import "path/to/file.proto";' to reuse messages.
   This allows you to create a "Shared Library" of standard types (Address, User).

4. FOLDER STRUCTURE: Align your file folders with your package names.
   (e.g., package 'finance.payroll' -> folder 'finance/payroll/').
	`)
}
