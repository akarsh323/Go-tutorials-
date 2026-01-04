package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ============================================================================
// 92 Environment Variables
// ============================================================================
//
// Environment variables are dynamic key-value pairs in the OS environment.
// They allow you to pass configuration to your app WITHOUT hard-coding it.
//
// Why use them?
// - Configuration: Dev vs. Production behavior
// - Security: Keep credentials out of source code
// - Flexibility: Change app behavior without recompiling
//
// ============================================================================

// ============================================================================
// PART 1: READING ENVIRONMENT VARIABLES (os.Getenv)
// ============================================================================

func Demo92_Part1_ReadingVariables() {
	fmt.Println("\n=== PART 1: READING VARIABLES (os.Getenv) ===")
	fmt.Println()

	fmt.Println("📌 Basic Usage - Reading System Variables:")
	// Common system variables that exist on most systems
	shell := os.Getenv("SHELL")  // /bin/bash or /bin/zsh
	home := os.Getenv("HOME")    // /Users/username
	user := os.Getenv("USER")    // username
	pathVar := os.Getenv("PATH") // /usr/local/bin:/usr/bin:...
	lang := os.Getenv("LANG")    // en_US.UTF-8

	fmt.Printf("  SHELL: %s\n", shell)
	fmt.Printf("  HOME:  %s\n", home)
	fmt.Printf("  USER:  %s\n", user)
	fmt.Printf("  LANG:  %s\n", lang)
	fmt.Printf("  PATH:  %s (truncated)\n", truncateString(pathVar, 50))
	fmt.Println()

	fmt.Println("📌 Non-Existent Variables Return Empty String:")
	nonExistent := os.Getenv("THIS_DOES_NOT_EXIST")
	fmt.Printf("  os.Getenv(\"THIS_DOES_NOT_EXIST\") = \"%s\" (empty string)\n", nonExistent)
	fmt.Printf("  Length: %d\n", len(nonExistent))
	fmt.Println()

	fmt.Println("📌 Safe Access Pattern (Check if Empty):")
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		fmt.Println("  API_KEY not set - using default or error handling")
	} else {
		fmt.Println("  API_KEY is set - proceeding with authentication")
	}
	fmt.Println()
}

// ============================================================================
// PART 2: SETTING & UNSETTING VARIABLES (os.Setenv, os.Unsetenv)
// ============================================================================

func Demo92_Part2_ModifyingVariables() {
	fmt.Println("\n=== PART 2: MODIFYING VARIABLES (os.Setenv, os.Unsetenv) ===")
	fmt.Println()

	fmt.Println("📌 Setting a Variable (os.Setenv):")
	err := os.Setenv("MY_APP_COLOR", "blue")
	if err != nil {
		fmt.Printf("  Error setting variable: %v\n", err)
	} else {
		color := os.Getenv("MY_APP_COLOR")
		fmt.Printf("  Set MY_APP_COLOR=\"blue\"\n")
		fmt.Printf("  Retrieved: %s\n", color)
	}
	fmt.Println()

	fmt.Println("📌 Updating a Variable:")
	os.Setenv("MY_APP_COLOR", "red")
	color := os.Getenv("MY_APP_COLOR")
	fmt.Printf("  Updated MY_APP_COLOR to \"red\"\n")
	fmt.Printf("  Retrieved: %s\n", color)
	fmt.Println()

	fmt.Println("📌 Unsetting a Variable (os.Unsetenv):")
	os.Setenv("TEMP_VAR", "temp_value")
	fmt.Printf("  Before unset: %s\n", os.Getenv("TEMP_VAR"))
	os.Unsetenv("TEMP_VAR")
	fmt.Printf("  After unset:  \"%s\" (empty)\n", os.Getenv("TEMP_VAR"))
	fmt.Println()

	fmt.Println("📌 Important: Changes Only Affect Current Process")
	fmt.Println("  Variables set in Go do NOT affect the parent shell.")
	fmt.Println("  Example:")
	fmt.Println("    $ export SHELL_VAR=test")
	fmt.Println("    $ go run main.go  # Inside: os.Setenv(\"SHELL_VAR\", \"new\")")
	fmt.Println("    $ echo $SHELL_VAR # Still shows 'test' (not affected)")
	fmt.Println()
}

// ============================================================================
// PART 3: LISTING ALL VARIABLES (os.Environ)
// ============================================================================

func Demo92_Part3_ListAllVariables() {
	fmt.Println("\n=== PART 3: LISTING ALL VARIABLES (os.Environ) ===")
	fmt.Println()

	fmt.Println("📌 os.Environ() Returns All Variables:")
	fmt.Println("  Returns: []string where each element is \"KEY=VALUE\"")
	fmt.Println()

	// Get all environment variables
	allVars := os.Environ()
	fmt.Printf("  Total environment variables: %d\n", len(allVars))
	fmt.Println()

	// Show first few
	fmt.Println("📌 First 5 Variables (Raw Format):")
	for i := 0; i < 5 && i < len(allVars); i++ {
		fmt.Printf("  [%d] %s\n", i, allVars[i])
	}
	fmt.Println()

	// Parse and display nicely
	fmt.Println("📌 Parsing Variables (Separated Key & Value):")
	fmt.Println("  Using strings.SplitN(entry, \"=\", 2):")
	for i := 0; i < 3 && i < len(allVars); i++ {
		parts := strings.SplitN(allVars[i], "=", 2)
		if len(parts) == 2 {
			fmt.Printf("  Key: %-20s | Value: %s\n", parts[0], truncateString(parts[1], 40))
		}
	}
	fmt.Println()
}

// ============================================================================
// PART 4: THE PARSING LOGIC - WHY strings.SplitN?
// ============================================================================

func Demo92_Part4_SplitNLogic() {
	fmt.Println("\n=== PART 4: PARSING LOGIC - Why strings.SplitN? ===")
	fmt.Println()

	fmt.Println("📌 The Problem: Variables Can Contain '=' Characters")
	fmt.Println("  Example: DB_CONN=postgres://user:pass@localhost:5432/db?key=value")
	fmt.Println()
	fmt.Println("  If we use Split(\"=\", -1), we get:")
	fmt.Println("    [\"DB_CONN\", \"postgres://user:pass@localhost:5432/db?key\", \"value\"]")
	fmt.Println("  ✗ The value is broken into multiple parts!")
	fmt.Println()

	// Demonstrate the problem
	dbConn := "DB_CONN=postgres://user:pass@localhost:5432/db?key=value"
	wrongSplit := strings.Split(dbConn, "=")
	fmt.Println("  strings.Split(entry, \"=\"):")
	for i, part := range wrongSplit {
		fmt.Printf("    [%d] %s\n", i, part)
	}
	fmt.Println()

	fmt.Println("📌 The Solution: strings.SplitN(string, separator, n)")
	fmt.Println("  The 'n' parameter controls how many pieces to return:")
	fmt.Println()

	exampleVar := "KEY=value1=value2=value3"

	fmt.Println("  Example variable: \"KEY=value1=value2=value3\"")
	fmt.Println()

	fmt.Println("  • SplitN(entry, \"=\", -1) - Split at ALL separators:")
	parts := strings.SplitN(exampleVar, "=", -1)
	fmt.Printf("    Result: %v\n", parts)
	fmt.Println("    ✗ Not useful for parsing env vars")
	fmt.Println()

	fmt.Println("  • SplitN(entry, \"=\", 1) - Returns whole string:")
	parts = strings.SplitN(exampleVar, "=", 1)
	fmt.Printf("    Result: %v\n", parts)
	fmt.Println("    ✗ No separation of key and value")
	fmt.Println()

	fmt.Println("  • SplitN(entry, \"=\", 2) - Split at FIRST '=' only:")
	parts = strings.SplitN(exampleVar, "=", 2)
	fmt.Printf("    Result: %v\n", parts)
	fmt.Printf("    ✓ Key: %s, Value: %s\n", parts[0], parts[1])
	fmt.Println("    ✓ This is the correct approach!")
	fmt.Println()

	// Real-world example
	fmt.Println("📌 Real-World Example: Database Connection String")
	realDBConn := "DATABASE_URL=postgresql://admin:SecurePass123@db.example.com:5432/mydb?sslmode=require&timeout=30s"
	pair := strings.SplitN(realDBConn, "=", 2)
	fmt.Printf("  Input: %s\n", realDBConn)
	fmt.Printf("  Key:   %s\n", pair[0])
	fmt.Printf("  Value: %s\n", pair[1])
	fmt.Println()
}

// ============================================================================
// PART 5: HELPER FUNCTION - GetEnv with Default
// ============================================================================

func Demo92_Part5_HelperFunction() {
	fmt.Println("\n=== PART 5: HELPER FUNCTION - GetEnv with Default ===")
	fmt.Println()

	fmt.Println("📌 Problem: What if a Variable Doesn't Exist?")
	fmt.Println("  os.Getenv returns empty string if not set.")
	fmt.Println("  We need a pattern for defaults.")
	fmt.Println()

	fmt.Println("📌 Solution: Create a Helper Function")
	fmt.Println()

	// Test the helper function
	port := getEnv("PORT", "8080")
	fmt.Printf("  getEnv(\"PORT\", \"8080\") = %s\n", port)

	timeout := getEnv("TIMEOUT", "30")
	fmt.Printf("  getEnv(\"TIMEOUT\", \"30\") = %s\n", timeout)

	// Set a variable and test again
	os.Setenv("DEBUG", "true")
	debug := getEnv("DEBUG", "false")
	fmt.Printf("  getEnv(\"DEBUG\", \"false\") = %s (variable was set)\n", debug)
	fmt.Println()

	fmt.Println("📌 Type-Specific Helpers:")
	portInt := getEnvInt("PORT", 8080)
	fmt.Printf("  getEnvInt(\"PORT\", 8080) = %d (type: int)\n", portInt)

	os.Setenv("DEBUG_MODE", "true")
	debugBool := getEnvBool("DEBUG_MODE", false)
	fmt.Printf("  getEnvBool(\"DEBUG_MODE\", false) = %v (type: bool)\n", debugBool)
	fmt.Println()
}

// Helper function: Get env var with default
func getEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// Helper function: Get env var as integer
func getEnvInt(key string, defaultValue int) int {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// Helper function: Get env var as boolean
func getEnvBool(key string, defaultValue bool) bool {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

// ============================================================================
// PART 6: BEST PRACTICES
// ============================================================================

func Demo92_Part6_BestPractices() {
	fmt.Println("\n=== PART 6: BEST PRACTICES ===")
	fmt.Println()

	fmt.Println("📌 Best Practice 1: UPPERCASE Naming Convention")
	fmt.Println("  ✓ Correct:  API_KEY, DATABASE_URL, DEBUG_MODE")
	fmt.Println("  ✗ Avoid:    api_key, database_url, debug_mode")
	fmt.Println()

	fmt.Println("📌 Best Practice 2: Never Log Sensitive Variables")
	fmt.Println("  ✗ WRONG:   fmt.Println(\"Password:\", os.Getenv(\"DB_PASSWORD\"))")
	fmt.Println("  ✓ RIGHT:   fmt.Println(\"Using database with password (not logged)\")")
	fmt.Println()

	fmt.Println("📌 Best Practice 3: Always Have Defaults")
	fmt.Println("  ✗ WRONG:")
	fmt.Println("    port := os.Getenv(\"PORT\") // Could be empty!")
	fmt.Println("    listener, _ := net.Listen(\"tcp\", \":\"+port)")
	fmt.Println()
	fmt.Println("  ✓ RIGHT:")
	fmt.Println("    port := getEnv(\"PORT\", \"8080\")")
	fmt.Println("    listener, _ := net.Listen(\"tcp\", \":\"+port)")
	fmt.Println()

	fmt.Println("📌 Best Practice 4: Validate at Startup")
	fmt.Println("  Check required variables when app starts, not later.")
	validateRequiredEnvVars()
	fmt.Println()

	fmt.Println("📌 Best Practice 5: Document Expected Variables")
	fmt.Println("  Create a .env.example file or documentation listing:")
	fmt.Println("  - Variable names (e.g., DATABASE_URL)")
	fmt.Println("  - What they control (e.g., \"PostgreSQL connection string\")")
	fmt.Println("  - Expected format (e.g., \"postgres://user:pass@host:5432/db\")")
	fmt.Println("  - Default value if applicable")
	fmt.Println()

	fmt.Println("📌 Best Practice 6: Cross-Platform Considerations")
	fmt.Println("  Windows, macOS, and Linux handle env vars slightly differently.")
	fmt.Println("  - Always use os.Getenv (handles platform differences)")
	fmt.Println("  - Use forward slashes for paths or os.PathSeparator")
	fmt.Println("  - Test on target platform")
	fmt.Println()
}

func validateRequiredEnvVars() {
	requiredVars := []string{"HOME", "SHELL"}
	for _, varName := range requiredVars {
		if os.Getenv(varName) == "" {
			fmt.Printf("  ⚠️  Required variable '%s' is not set\n", varName)
		} else {
			fmt.Printf("  ✓ Required variable '%s' is set\n", varName)
		}
	}
}

// ============================================================================
// PART 7: COMPLETE EXAMPLE - App Configuration
// ============================================================================

func Demo92_Part7_CompleteExample() {
	fmt.Println("\n=== PART 7: COMPLETE EXAMPLE - App Configuration ===")
	fmt.Println()

	fmt.Println("📌 Real-World App Configuration Pattern:")
	fmt.Println()

	// Simulate different environment configurations
	os.Setenv("APP_ENV", "production")
	os.Setenv("DATABASE_URL", "postgresql://user:pass@prod-db:5432/app")
	os.Setenv("LOG_LEVEL", "info")

	config := loadAppConfig()
	fmt.Printf("  Environment: %s\n", config.Env)
	fmt.Printf("  Database:    %s\n", truncateString(config.DatabaseURL, 50))
	fmt.Printf("  Log Level:   %s\n", config.LogLevel)
	fmt.Printf("  Port:        %d\n", config.Port)
	fmt.Printf("  Debug:       %v\n", config.Debug)
	fmt.Println()
}

type AppConfig struct {
	Env         string
	DatabaseURL string
	LogLevel    string
	Port        int
	Debug       bool
}

func loadAppConfig() AppConfig {
	return AppConfig{
		Env:         getEnv("APP_ENV", "development"),
		DatabaseURL: getEnv("DATABASE_URL", "sqlite://app.db"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),
		Port:        getEnvInt("PORT", 8080),
		Debug:       getEnvBool("DEBUG", false),
	}
}

// ============================================================================
// HELPER UTILITIES
// ============================================================================

func truncateString(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen] + "..."
	}
	return s
}

// ============================================================================
// MAIN DEMO FUNCTION
// ============================================================================

func Demo92EnvVars() {
	fmt.Println("-- 92 Environment Variables --")
	fmt.Println()
	fmt.Println("Environment variables are dynamic key-value pairs from the OS.")
	fmt.Println("Use them for configuration, secrets, and platform-specific settings.")
	fmt.Println()

	Demo92_Part1_ReadingVariables()
	Demo92_Part2_ModifyingVariables()
	Demo92_Part3_ListAllVariables()
	Demo92_Part4_SplitNLogic()
	Demo92_Part5_HelperFunction()
	Demo92_Part6_BestPractices()
	Demo92_Part7_CompleteExample()

	fmt.Println("\n=== SUMMARY ===")
	fmt.Println("✓ os.Getenv() - Read variables")
	fmt.Println("✓ os.Setenv() - Create/update variables (current process only)")
	fmt.Println("✓ os.Unsetenv() - Remove variables")
	fmt.Println("✓ os.Environ() - List all variables")
	fmt.Println("✓ Use strings.SplitN(str, \"=\", 2) for safe parsing")
	fmt.Println("✓ Always provide defaults for optional variables")
	fmt.Println("✓ Never hardcode sensitive values - use env vars instead")
}
