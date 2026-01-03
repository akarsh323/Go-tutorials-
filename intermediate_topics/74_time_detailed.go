package intermediate

import (
	"fmt"
	"time"
)

// Topic 74: time
// Combines pragmatic examples with GoBootcamp patterns


func main() {

	fmt.Println("-- 74 Time --")

	// Get current time in local timezone
	now := time.Now()
	fmt.Println("Current time:", now)

	// Get current time in UTC
	nowUTC := time.Now().UTC()
	fmt.Println("Current UTC:", nowUTC)

	// Time components
	fmt.Printf("Year: %d, Month: %d, Day: %d\n", now.Year(), now.Month(), now.Day())
	fmt.Printf("Hour: %d, Minute: %d, Second: %d\n", now.Hour(), now.Minute(), now.Second())

	// Duration operations
	d := 2 * time.Hour
	fmt.Println("Duration (2 hours):", d)
	fmt.Println("Total seconds:", d.Seconds())
	fmt.Println("Total milliseconds:", d.Milliseconds())

	// Add time
	future := now.Add(24 * time.Hour)
	fmt.Println("Tomorrow:", future)

	// Subtract time using Add with negative duration
	yesterday := now.Add(-24 * time.Hour)
	fmt.Println("Yesterday:", yesterday)

	// Time subtraction: Sub returns a Duration
	diff := future.Sub(now)
	fmt.Println("Difference (future - now):", diff)
	fmt.Println("Hours difference:", diff.Hours())

	// Comparison operations
	t1 := time.Now()
	time.Sleep(1 * time.Millisecond)
	t2 := time.Now()

	fmt.Println("t1 Before t2:", t1.Before(t2))
	fmt.Println("t2 After t1:", t2.After(t1))
	fmt.Println("t1 Equal t2:", t1.Equal(t2))

	// Creating specific time
	birthday := time.Date(1990, time.December, 25, 15, 30, 0, 0, time.UTC)
	fmt.Println("Birthday:", birthday)

	// Time zones
	fmt.Println("Time in UTC:", now.In(time.UTC))

	// Weekday
	fmt.Println("Day of week:", now.Weekday())
	fmt.Println("Day of year:", now.YearDay())
}

func timeExample1() {
	
	fmt.Println("📚 Get the current time in various formats")

	now := time.Now()

	fmt.Printf("Current time: %v\n", now)
	fmt.Printf("Year: %d\n", now.Year())
	fmt.Printf("Month: %v\n", now.Month())
	fmt.Printf("Day: %d\n", now.Day())
	fmt.Printf("Hour: %d\n", now.Hour())
	fmt.Printf("Minute: %d\n", now.Minute())
	fmt.Printf("Second: %d\n", now.Second())
	fmt.Printf("Nanosecond: %d\n", now.Nanosecond())

	// Time zones
	fmt.Printf("\nTime zone: %s\n", now.Location())
	fmt.Printf("UTC: %v\n", now.UTC())
}

func timeExample10() {
	
	fmt.Println(`
Key Functions:

Current Time:
  time.Now()               → Current time
  time.Now().UTC()         → Current UTC time

Creating Times:
  time.Date(y, m, d, h, mn, s, ns, loc)
  time.Parse(layout, string)

Time Arithmetic:
  t.Add(duration)          → Add duration to time
  t.AddDate(y, m, d)       → Add years/months/days
  t.Sub(other)             → Difference between times
  time.Since(t)            → Time since t until now

Comparison:
  t.Before(other)          → Is t before other?
  t.After(other)           → Is t after other?
  t.Equal(other)           → Are they equal?

Formatting:
  t.Format(layout)         → Format as string
  
Parsing:
  time.Parse(layout, str)  → Parse string to time

Key Durations:
  time.Nanosecond
  time.Microsecond
  time.Millisecond
  time.Second
  time.Minute
  time.Hour

Predefined Layouts:
  time.RFC3339             → "2006-01-02T15:04:05Z07:00"
  time.RFC1123             → "Mon, 02 Jan 2006 15:04:05 MST"
  time.Kitchen             → "3:04PM"

Best Practices:
  ✓ Always use UTC internally
  ✓ Convert to local time only for display
  ✓ Use time.Now() for current time
  ✓ Test time-dependent code with fixed times
  ✓ Remember: format reference is Mon Jan 2 15:04:05 MST 2006
	`)
}
