package intermediate

import (
	"fmt"
	"time"
)

// Topic 76: Time Format and Parse - Bridging Machines and Humans
// ==============================================================
// This lesson covers Go's unique approach to formatting and parsing time.
// We learn the "Magic Reference Date" and how it differs from other languages.
// Topics: Custom layouts, ISO 8601, timezone handling, error handling, and real-world parsing.

func main() {
	fmt.Println("=== Topic 76: Time Formatting and Parsing - The Translation Layer ===\n")

	lesson1MagicReferenceDate()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson2ParsingISO8601()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson3CustomFormats()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson4FormattingForDisplay()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson5ImportantGotchas()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson6TeachersInsight()
}

// LESSON 1: The "Magic Reference Date" Rule
// ==========================================
func lesson1MagicReferenceDate() {
	fmt.Println("LESSON 1: THE 'MAGIC REFERENCE DATE' RULE")
	fmt.Println("------------------------------------------\n")

	fmt.Println("THE PROBLEM WITH OTHER LANGUAGES:")
	fmt.Println("  Most languages use codes like:")
	fmt.Println("    • YYYY-MM-DD (Python, JavaScript, etc.)")
	fmt.Println("    • mm/dd/yyyy (Excel, some legacy systems)")
	fmt.Println("  These codes are intuitive but arbitrary.\n")

	fmt.Println("GO'S UNIQUE APPROACH:")
	fmt.Println("  Go doesn't use codes. Instead, you format a SPECIFIC REFERENCE DATE.")
	fmt.Println("  This date is:\n")

	refDate := "Mon Jan 2 15:04:05 MST 2006"
	fmt.Printf("  ┌─────────────────────────────────┐\n")
	fmt.Printf("  │  %s  │\n", refDate)
	fmt.Printf("  └─────────────────────────────────┘\n\n")

	fmt.Println("WHY THIS SPECIFIC DATE?")
	fmt.Println("  Because the numbers count up in American date/time order:\n")

	fmt.Println("  Mon = Weekday name")
	fmt.Println("  Jan = Month (1 = January)")
	fmt.Println("  2   = Day (2 = 2nd)")
	fmt.Println("  15  = Hour (3 = 3 PM in 24-hour format)")
	fmt.Println("  04  = Minute (4 = 4 minutes)")
	fmt.Println("  05  = Second (5 = 5 seconds)")
	fmt.Println("  MST = Timezone (-07:00)")
	fmt.Println("  2006= Year (6 = 2006)\n")

	fmt.Println("THE KEY INSIGHT:")
	fmt.Println("  When you write a layout string in Go, you're saying:")
	fmt.Println("  \"Rearrange THIS reference date into the pattern I want,")
	fmt.Println("   then apply that logic to my actual data.\"\n")

	// Practical demonstration
	now := time.Now().UTC()
	fmt.Println("PRACTICAL DEMONSTRATION:")
	fmt.Printf("  Reference:     %s\n", "Mon Jan 2 15:04:05 MST 2006")
	fmt.Printf("  Current time:  %s\n", now.Format("Mon Jan 2 15:04:05 MST 2006"))
	fmt.Printf("  (They use the same format!)\n")
}

// LESSON 2: Practical Coding - Parsing ISO 8601
// =============================================
func lesson2ParsingISO8601() {
	fmt.Println("LESSON 2: PRACTICAL CODING - PARSING ISO 8601")
	fmt.Println("----------------------------------------------\n")

	fmt.Println("SCENARIO:")
	fmt.Println("  You receive a timestamp string from an API (internet standard format)")
	fmt.Println("  Input: \"2024-07-04T14:30:18Z\"")
	fmt.Println("  Goal: Convert it to a Go time.Time object for calculations\n")

	fmt.Println("STEP 1: Define the Layout")
	fmt.Println("  The layout describes the input format using the reference date:")

	layoutISO := "2006-01-02T15:04:05Z07:00"
	fmt.Printf("    Layout: %s\n", layoutISO)
	fmt.Println("    • 2006 = Year (4-digit)")
	fmt.Println("    • 01   = Month (2-digit, zero-padded)")
	fmt.Println("    • 02   = Day (2-digit, zero-padded)")
	fmt.Println("    • T    = Literal 'T' character")
	fmt.Println("    • 15   = Hour in 24-hour format")
	fmt.Println("    • 04   = Minutes")
	fmt.Println("    • 05   = Seconds")
	fmt.Println("    • Z07:00 = Timezone offset\n")

	fmt.Println("STEP 2: The Input String")
	dateString := "2024-07-04T14:30:18Z"
	fmt.Printf("    Input: %s\n", dateString)
	fmt.Println("    (This is real data, maybe from an API)\n")

	fmt.Println("STEP 3: Parse the Time")
	fmt.Println("    Code: t, err := time.Parse(layoutISO, dateString)")
	fmt.Println("    Returns TWO values:")
	fmt.Println("      • t   = the time.Time object (if successful)")
	fmt.Println("      • err = error (if parsing failed)\n")

	t, err := time.Parse(layoutISO, dateString)

	fmt.Println("STEP 4: Professional Error Handling")
	fmt.Println("    ALWAYS check if err != nil when parsing!")

	if err != nil {
		fmt.Printf("    Error: %v\n", err)
		return
	}

	fmt.Println("    ✓ Parse successful!\n")

	fmt.Println("RESULT:")
	fmt.Printf("  Parsed Time:      %v\n", t)
	fmt.Printf("  Year:             %d\n", t.Year())
	fmt.Printf("  Month:            %v\n", t.Month())
	fmt.Printf("  Day:              %d\n", t.Day())
	fmt.Printf("  Hour:             %d\n", t.Hour())
	fmt.Printf("  Minute:           %d\n", t.Minute())
	fmt.Printf("  Second:           %d\n", t.Second())
	fmt.Printf("  Timezone:         %v\n\n", t.Location())

	fmt.Println("WHY THIS MATTERS:")
	fmt.Println("  • ISO 8601 is the internet standard for timestamps")
	fmt.Println("  • JSON APIs use this format")
	fmt.Println("  • Databases often store/return timestamps in this format")
	fmt.Println("  • Knowing how to parse it is essential for modern Go development")
}

// LESSON 3: Custom Formats - Real-World Messiness
// ================================================
func lesson3CustomFormats() {
	fmt.Println("LESSON 3: CUSTOM FORMATS - REAL-WORLD MESSINESS")
	fmt.Println("-----------------------------------------------\n")

	fmt.Println("THE REALITY:")
	fmt.Println("  Not all data arrives in ISO 8601 format.")
	fmt.Println("  Real-world data is often 'messy' and inconsistent.")
	fmt.Println("  You need to write custom layouts to match it.\n")

	// Example 1
	fmt.Println("EXAMPLE 1: Month Name with AM/PM")
	fmt.Println("  Input: \"July 03, 2024 03:18 PM\"")
	str1 := "July 03, 2024 03:18 PM"
	layout1 := "January 02, 2006 03:04 PM"

	fmt.Printf("  Layout: %s\n", layout1)
	fmt.Println("    • January = Full month name (Go uses reference month)")
	fmt.Println("    • 02      = Day with zero-padding")
	fmt.Println("    • 2006    = 4-digit year")
	fmt.Println("    • 03      = Hour in 12-hour format (not 15!)")
	fmt.Println("    • 04      = Minutes")
	fmt.Println("    • PM      = AM/PM designator\n")

	t1, err := time.Parse(layout1, str1)
	if err != nil {
		fmt.Printf("  Error: %v\n\n", err)
	} else {
		fmt.Printf("  ✓ Parsed: %v\n\n", t1)
	}

	// Example 2
	fmt.Println("EXAMPLE 2: Abbreviated Month with Different Separators")
	fmt.Println("  Input: \"04-Jul-2024 14:30\"")
	str2 := "04-Jul-2024 14:30"
	layout2 := "02-Jan-2006 15:04"

	fmt.Printf("  Layout: %s\n", layout2)
	fmt.Println("    • 02      = Day")
	fmt.Println("    • Jan     = 3-letter month abbreviation")
	fmt.Println("    • 2006    = 4-digit year")
	fmt.Println("    • 15      = 24-hour format")
	fmt.Println("    • 04      = Minutes\n")

	t2, err := time.Parse(layout2, str2)
	if err != nil {
		fmt.Printf("  Error: %v\n\n", err)
	} else {
		fmt.Printf("  ✓ Parsed: %v\n\n", t2)
	}

	// Example 3
	fmt.Println("EXAMPLE 3: Database Format (No Timezone)")
	fmt.Println("  Input: \"2024-01-15 14:30:45\"")
	str3 := "2024-01-15 14:30:45"
	layout3 := "2006-01-02 15:04:05"

	fmt.Printf("  Layout: %s\n", layout3)
	fmt.Println("    • Simple, clean format")
	fmt.Println("    • Often used in databases")
	fmt.Println("    • No timezone info (Go assumes UTC)\n")

	t3, err := time.Parse(layout3, str3)
	if err != nil {
		fmt.Printf("  Error: %v\n\n", err)
	} else {
		fmt.Printf("  ✓ Parsed: %v\n\n", t3)
	}

	fmt.Println("KEY PRINCIPLE:")
	fmt.Println("  Your layout string must EXACTLY match the input string.")
	fmt.Println("  Every space, dash, letter must correspond.")
	fmt.Println("  If they don't match, time.Parse will return an error.")
}

// LESSON 4: Formatting for Display
// =================================
func lesson4FormattingForDisplay() {
	fmt.Println("LESSON 4: FORMATTING FOR DISPLAY")
	fmt.Println("---------------------------------\n")

	fmt.Println("THE INVERSE OPERATION:")
	fmt.Println("  While parsing converts strings TO time objects,")
	fmt.Println("  formatting converts time objects TO strings.")
	fmt.Println("  You use the SAME layout strings!\n")

	now := time.Now().UTC()

	fmt.Println("EXAMPLE FORMATS:\n")

	// Format 1: ISO 8601 (standard)
	iso := now.Format("2006-01-02T15:04:05Z07:00")
	fmt.Printf("ISO 8601:                %s\n", iso)
	fmt.Println("  Use for: APIs, JSON, web services\n")

	// Format 2: Simple date
	simple := now.Format("2006-01-02")
	fmt.Printf("Simple Date (YYYY-MM-DD): %s\n", simple)
	fmt.Println("  Use for: Databases, logs\n")

	// Format 3: Human readable
	human := now.Format("Monday, January 2, 2006 at 3:04 PM")
	fmt.Printf("Human Readable:           %s\n", human)
	fmt.Println("  Use for: User displays, emails\n")

	// Format 4: RFC3339 (predefined)
	rfc := now.Format(time.RFC3339)
	fmt.Printf("RFC3339 (predefined):     %s\n", rfc)
	fmt.Println("  Use for: Standard internet format\n")

	// Format 5: Custom business format
	custom := now.Format("Jan 02, 2006 - 3:04 PM MST")
	fmt.Printf("Custom Business Format:   %s\n", custom)
	fmt.Println("  Use for: Reports, business documents\n")

	fmt.Println("PREDEFINED FORMATS IN GO:")
	fmt.Println("  time.RFC3339   = \"2006-01-02T15:04:05Z07:00\" (Standard)")
	fmt.Println("  time.RFC1123   = \"Mon, 02 Jan 2006 15:04:05 MST\"")
	fmt.Println("  time.RFC1123Z  = \"Mon, 02 Jan 2006 15:04:05 -0700\"")
	fmt.Println("  time.Kitchen   = \"3:04PM\"")
	fmt.Println("  time.UnixDate  = \"Mon Jan 2 15:04:05 MST 2006\" (The reference!)\n")

	fmt.Println("RULE OF THUMB:")
	fmt.Println("  ✓ Store time in database as Unix timestamp (number)")
	fmt.Println("  ✓ Parse/format strings only for I/O (APIs, display)")
	fmt.Println("  ✓ Use RFC3339 for APIs and JSON")
	fmt.Println("  ✓ Use custom formats only when necessary for legacy systems")
}

// LESSON 5: Important Gotchas (Common Mistakes)
// =============================================
func lesson5ImportantGotchas() {
	fmt.Println("LESSON 5: IMPORTANT GOTCHAS - COMMON MISTAKES")
	fmt.Println("---------------------------------------------\n")

	fmt.Println("GOTCHA #1: Timezone Ambiguity")
	fmt.Println("  If you parse a time WITHOUT timezone info:")
	fmt.Println("    Input:  \"2024-01-15 14:30:45\"")
	fmt.Println("    Layout: \"2006-01-02 15:04:05\"")
	fmt.Println("  Go assumes UTC. But your user might be in New York or India!")
	fmt.Println("  This causes silent bugs that only appear in certain timezones.\n")

	fmt.Println("  SOLUTION:")
	fmt.Println("  Use time.ParseInLocation() to parse in a specific timezone:")

	loc, _ := time.LoadLocation("America/New_York")
	str := "2024-01-15 14:30:45"
	layout := "2006-01-02 15:04:05"
	t, err := time.ParseInLocation(layout, str, loc)
	if err == nil {
		fmt.Printf("    Parsed in NY timezone: %v\n", t)
	}
	fmt.Println()

	fmt.Println("GOTCHA #2: Layout Mismatch Causes Silent Failures")
	fmt.Println("  If your layout doesn't match your input, time.Parse fails:")

	badLayout := "2006-01-02" // Expects YYYY-MM-DD
	badInput := "01/15/2024"  // But got MM/DD/YYYY
	_, err = time.Parse(badLayout, badInput)

	fmt.Printf("    Layout: %s\n", badLayout)
	fmt.Printf("    Input:  %s\n", badInput)
	fmt.Printf("    Error:  %v\n", err)
	fmt.Println("    → Always test your layout strings with real data!\n")

	fmt.Println("GOTCHA #3: Variable Scope and Re-declaration")
	fmt.Println("  In Go, you can't re-declare a variable with :=")
	fmt.Println("    Wrong:  t, err := time.Parse(layout, str1)")
	fmt.Println("            t, err := time.Parse(layout, str2)  // ← Error!")
	fmt.Println("            (Can't re-declare 'err')")
	fmt.Println("    Right:  t, err := time.Parse(layout, str1)")
	fmt.Println("            t, err = time.Parse(layout, str2)   // ← Use '=' not ':='\n")

	fmt.Println("GOTCHA #4: Padding and Format Strictness")
	fmt.Println("  Layouts must be EXACT:")
	fmt.Println("    • Use '02' for day (zero-padded)")
	fmt.Println("    • Use '2' for day (not padded)")
	fmt.Println("    • Use '01' for month (zero-padded)")
	fmt.Println("    • Use '1' for month (not padded)")
	fmt.Println("    If input is '5' but layout says '05', it will fail!\n")

	fmt.Println("GOTCHA #5: 12-hour vs 24-hour Format")
	fmt.Println("  Use '15' for 24-hour (0-23)")
	fmt.Println("  Use '03' for 12-hour (1-12, requires AM/PM)")
	fmt.Println("    Wrong:  \"14:30\" with layout \"03:04\" (mismatch!)")
	fmt.Println("    Right:  \"02:30 PM\" with layout \"03:04 PM\"")
}

// LESSON 6: Teacher's Insight - The Time Trilogy Completion
// ==========================================================
func lesson6TeachersInsight() {
	fmt.Println("LESSON 6: TEACHER'S INSIGHT - COMPLETING THE TIME TRILOGY")
	fmt.Println("---------------------------------------------------------\n")

	fmt.Println("YOU HAVE NOW LEARNED THE COMPLETE TIME STORY:\n")

	fmt.Println("1. EPOCH (Lesson 75)")
	fmt.Println("   • What is time? Raw numbers.")
	fmt.Println("   • Seconds since Jan 1, 1970 UTC")
	fmt.Println("   • This is how computers think\n")

	fmt.Println("2. FORMATTING/PARSING (Lesson 76)")
	fmt.Println("   • How to convert between strings and time objects")
	fmt.Println("   • The Magic Reference Date: Mon Jan 2 15:04:05 MST 2006")
	fmt.Println("   • This is the translation layer\n")

	fmt.Println("3. PRACTICAL WORKFLOW")
	fmt.Println("   • Receive data (string from API/user/database)")
	fmt.Println("   • Parse it with time.Parse()")
	fmt.Println("   • Do calculations with time.Time objects")
	fmt.Println("   • Format results with .Format() for display\n")

	fmt.Println("THE FUNDAMENTAL PRINCIPLE:")
	fmt.Println("┌─────────────────────────────────────────────────────────────┐")
	fmt.Println("│  Computers like NUMBERS (epoch/Unix timestamps)             │")
	fmt.Println("│  Humans like TEXT (formatted dates)                         │")
	fmt.Println("│  Go bridges them with Parse (text→time) and Format (time→text) │")
	fmt.Println("└─────────────────────────────────────────────────────────────┘\n")

	fmt.Println("PROFESSIONAL CHECKLIST:")
	fmt.Println("  ☐ Always handle parse errors (if err != nil)")
	fmt.Println("  ☐ Test layouts with actual data before going to production")
	fmt.Println("  ☐ Be aware of timezone implications")
	fmt.Println("  ☐ Use RFC3339 for APIs and interop with other systems")
	fmt.Println("  ☐ Use custom layouts only when necessary")
	fmt.Println("  ☐ Store time as numbers (Unix epoch) in databases")
	fmt.Println("  ☐ Format to human-readable only for display\n")

	fmt.Println("YOU'RE NOW READY FOR:")
	fmt.Println("  ✓ Parsing API responses with timestamps")
	fmt.Println("  ✓ Reading timestamps from databases")
	fmt.Println("  ✓ Formatting dates for user display")
	fmt.Println("  ✓ Working with timezones safely")
	fmt.Println("  ✓ Building production-grade time handling code")
}
