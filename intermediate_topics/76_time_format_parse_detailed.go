package intermediate

import (
	"fmt"
	"time"
)

// Topic 76: time_format_parse
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 76 Time Formatting / Parsing --")

	now := time.Now().UTC()

	// Standard RFC3339 format (ISO 8601)
	rfc3339 := now.Format(time.RFC3339)
	fmt.Println("RFC3339:", rfc3339)

	// RFC1123 format (HTTP header style)
	rfc1123 := now.Format(time.RFC1123Z)
	fmt.Println("RFC1123Z:", rfc1123)

	// Unix date format
	unixDate := now.Format(time.UnixDate)
	fmt.Println("UnixDate:", unixDate)

	// ANSIC format
	ansic := now.Format(time.ANSIC)
	fmt.Println("ANSIC:", ansic)

	// Custom format layout (Go uses Mon Jan 2 15:04:05 MST 2006 as reference)
	custom := now.Format("2006-01-02 15:04:05")
	fmt.Println("Custom (YYYY-MM-DD HH:MM:SS):", custom)

	// Custom format with additional info
	custom2 := now.Format("Monday, January 2, 2006 at 3:04 PM")
	fmt.Println("Custom (readable):", custom2)

	// Parse RFC3339 formatted time back
	s := now.Format(time.RFC3339)
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}
	fmt.Println("Formatted string:", s)
	fmt.Println("Parsed time equals original:", t.Equal(now))

	// Parse custom format
	customStr := "2024-01-15 14:30:45"
	customLayout := "2006-01-02 15:04:05"
	parsed, err := time.Parse(customLayout, customStr)
	if err != nil {
		fmt.Println("Parse error:", err)
	} else {
		fmt.Println("Parsed custom format:", parsed)
	}

	// Parse with timezone information
	tsStr := "2024-01-15T14:30:45Z"
	ts, err := time.Parse(time.RFC3339, tsStr)
	if err == nil {
		fmt.Println("Parsed with timezone:", ts)
	}

	// ParseInLocation: Parse time in a specific timezone
	loc, _ := time.LoadLocation("America/New_York")
	locStr := "2024-01-15 14:30:45"
	locTime, err := time.ParseInLocation(customLayout, locStr, loc)
	if err == nil {
		fmt.Println("Parsed in NY timezone:", locTime)
	}

	// Format with timezone abbreviation
	fmt.Println("Time with timezone:", now.Format("2006-01-02 15:04:05 MST"))
}

func formatParseExample1() {
	
	fmt.Println("📚 Go uses Mon Jan 2 15:04:05 MST 2006 as pattern reference")

	t := time.Date(2025, time.January, 15, 14, 30, 45, 0, time.UTC)

	fmt.Println("The reference format:")
	fmt.Println("  Mon Jan 2 15:04:05 MST 2006")
	fmt.Println("  This represents a specific date/time pattern")
	fmt.Println("  - Mon = weekday")
	fmt.Println("  - Jan = month name")
	fmt.Println("  - 2 = day of month (single digit)")
	fmt.Println("  - 15 = hour (24-hour)")
	fmt.Println("  - 04 = minutes")
	fmt.Println("  - 05 = seconds")
	fmt.Println("  - MST = timezone")
	fmt.Println("  - 2006 = year")

	fmt.Printf("\nYour time formatted as reference: %s\n", t.Format("Mon Jan 2 15:04:05 MST 2006"))
}

func formatParseExample10() {
	
	fmt.Println(`
Format/Parse Reference:

The Magic Date: Mon Jan 2 15:04:05 MST 2006
- Mon = weekday name (Monday, Tuesday, etc.)
- Jan = month name (January, February, etc.)
- 2 = day (use 2 for both single/double digits)
- 15 = hour 24-hour format (3 for 12-hour)
- 04 = minutes (04 not 4 - padding matters)
- 05 = seconds
- MST = timezone name/offset
- 2006 = year (use 06 for 2-digit year)

Numbers to Remember:
- Hour: 15 = 3 PM (24-hour format)
- Month: 01 (January), 12 (December)
- Day: 02 (2nd), 31 (31st)
- Year: 2006 (4-digit), 06 (2-digit)

Formatting (time → string):
  t.Format(layout)

Parsing (string → time):
  time.Parse(layout, string)
  time.ParseInLocation(layout, string, location)

Predefined Formats:
  time.RFC3339    → "2006-01-02T15:04:05Z07:00"
  time.RFC1123    → "Mon, 02 Jan 2006 15:04:05 MST"
  time.Kitchen    → "3:04PM"

Tips:
  ✓ Use RFC3339 for JSON/APIs (ISO 8601)
  ✓ Use 2006-01-02 for database storage
  ✓ Use _ for space padding (2006-01-_2)
  ✓ Always test with actual data
  ✓ Be careful with timezone info
  ✓ ParseInLocation for specific timezone
	`)
}
