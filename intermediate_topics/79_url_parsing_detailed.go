package intermediate

import (
	"fmt"
	"net/url"
)

// Topic 79: URL Parsing - Breaking Down and Building URLs
// ========================================================
// This lesson covers the anatomy of URLs, parsing them into components,
// reading query parameters, and safely building URLs programmatically.
// Essential for backend development and API interactions.

func main() {
	fmt.Println("=== Topic 79: URL Parsing - The Complete Guide ===\n")

	lesson1URLAnatomy()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson2ParsingURLs()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson3QueryParameters()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson4BuildingURLs()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson5URLEncoding()
	fmt.Println("\n" + string([]byte{61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61, 61}) + "\n")

	lesson6PracticalExercise()
}

// LESSON 1: The Anatomy of a URL
// =============================
func lesson1URLAnatomy() {
	fmt.Println("LESSON 1: THE ANATOMY OF A URL")
	fmt.Println("------------------------------\n")

	fmt.Println("A URL (Uniform Resource Locator) is a structured address.")
	fmt.Println("Understanding its parts is crucial for parsing and building.\n")

	fmt.Println("COMPLETE URL STRUCTURE:")
	fmt.Println("  scheme://userinfo@host:port/path?query#fragment\n")

	fmt.Println("COMPONENT BREAKDOWN:\n")

	fmt.Println("1. SCHEME (Protocol)")
	fmt.Println("   • How to access the resource")
	fmt.Println("   • Examples: http, https, ftp, file, ws (WebSocket)")
	fmt.Println("   • Most common: https (secure)\n")

	fmt.Println("2. USERINFO (Optional)")
	fmt.Println("   • Username and password")
	fmt.Println("   • Format: username:password")
	fmt.Println("   • Example: ftp://user:pass@ftp.example.com")
	fmt.Println("   • RARELY used (security risk)\n")

	fmt.Println("3. HOST")
	fmt.Println("   • Domain name or IP address")
	fmt.Println("   • Examples: example.com, 192.168.1.1, localhost\n")

	fmt.Println("4. PORT")
	fmt.Println("   • Network port (optional, has defaults)")
	fmt.Println("   • HTTP defaults to :80 (usually omitted)")
	fmt.Println("   • HTTPS defaults to :443 (usually omitted)")
	fmt.Println("   • Example: example.com:8080\n")

	fmt.Println("5. PATH")
	fmt.Println("   • The specific resource location on server")
	fmt.Println("   • Examples: /products/123, /users/profile, /api/v1/users")
	fmt.Println("   • Always starts with /\n")

	fmt.Println("6. QUERY (Query String)")
	fmt.Println("   • Key-value parameters (optional)")
	fmt.Println("   • Format: ?key1=value1&key2=value2")
	fmt.Println("   • Examples: ?search=golang&page=1&sort=date")
	fmt.Println("   • Used for filtering, pagination, search, etc.\n")

	fmt.Println("7. FRAGMENT (Anchor)")
	fmt.Println("   • Jump to specific section within page (optional)")
	fmt.Println("   • Format: #section_name")
	fmt.Println("   • Example: #table-of-contents")
	fmt.Println("   • Not sent to server (browser-side only)\n")

	fmt.Println("COMPLETE EXAMPLE WITH ALL PARTS:")
	fmt.Println("┌─────────────────────────────────────────────────────┐")
	fmt.Println("│ https://user:pass@example.com:8080/api/v1/users?  │")
	fmt.Println("│ id=123&active=true#profile-section                  │")
	fmt.Println("└─────────────────────────────────────────────────────┘\n")

	fmt.Println("│─────────┬──────────┬──────────────┬─────┬─────────────────┬──────────────┬──────────────┐")
	fmt.Println("│ https   │ user:pass│ example.com  │8080 │ /api/v1/users   │ id=123&...   │ profile-... │")
	fmt.Println("│ Scheme  │ Userinfo │ Host         │Port │ Path            │ Query        │ Fragment    │")
	fmt.Println("│─────────┴──────────┴──────────────┴─────┴─────────────────┴──────────────┴──────────────┘\n")

	fmt.Println("REAL-WORLD EXAMPLES:")
	examples := []string{
		"https://github.com/golang/go/issues?labels=bug",
		"https://www.google.com/search?q=golang&tbm=isch",
		"https://api.github.com/users/golang",
		"https://localhost:3000/admin?token=abc123",
	}

	for _, ex := range examples {
		fmt.Printf("  • %s\n", ex)
	}
}

// LESSON 2: Parsing URLs (Reading)
// ================================
func lesson2ParsingURLs() {
	fmt.Println("LESSON 2: PARSING URLs (READING)")
	fmt.Println("--------------------------------\n")

	fmt.Println("TO READ A URL:")
	fmt.Println("  Use url.Parse(urlString)")
	fmt.Println("  Returns: (*url.URL, error)\n")

	fmt.Println("THE PARSED STRUCT HAS THESE FIELDS:")
	fmt.Println("  u.Scheme      - Protocol (https, http, etc.)")
	fmt.Println("  u.Host        - Domain with port (example.com:8080)")
	fmt.Println("  u.Hostname()  - Domain only, no port")
	fmt.Println("  u.Port()      - Port only (as string)")
	fmt.Println("  u.Path        - Resource path (/users/123)")
	fmt.Println("  u.RawQuery    - Raw query string (id=1&name=john)")
	fmt.Println("  u.Fragment    - Anchor fragment (#section)\n")

	fmt.Println("EXAMPLE 1: Basic Parsing")
	urlStr := "https://example.com:8080/api/users?id=123&name=John#profile"
	fmt.Printf("Input URL:\n  %s\n\n", urlStr)

	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("PARSED COMPONENTS:")
	fmt.Printf("  Scheme:     %s\n", u.Scheme)
	fmt.Printf("  Host:       %s\n", u.Host)
	fmt.Printf("  Hostname:   %s\n", u.Hostname())
	fmt.Printf("  Port:       %s\n", u.Port())
	fmt.Printf("  Path:       %s\n", u.Path)
	fmt.Printf("  RawQuery:   %s\n", u.RawQuery)
	fmt.Printf("  Fragment:   %s\n\n", u.Fragment)

	fmt.Println("EXAMPLE 2: Simple URL (no port, no query)")
	simple := "https://golang.org/doc/install"
	fmt.Printf("Input: %s\n\n", simple)

	u2, _ := url.Parse(simple)
	fmt.Printf("  Scheme:     %s\n", u2.Scheme)
	fmt.Printf("  Host:       %s\n", u2.Host)
	fmt.Printf("  Path:       %s\n", u2.Path)
	fmt.Printf("  RawQuery:   %s (empty)\n\n", u2.RawQuery)

	fmt.Println("EXAMPLE 3: Localhost with port")
	localhost := "http://localhost:3000/admin"
	fmt.Printf("Input: %s\n\n", localhost)

	u3, _ := url.Parse(localhost)
	fmt.Printf("  Scheme:     %s\n", u3.Scheme)
	fmt.Printf("  Host:       %s\n", u3.Host)
	fmt.Printf("  Hostname:   %s\n", u3.Hostname())
	fmt.Printf("  Port:       %s\n", u3.Port())
	fmt.Printf("  Path:       %s\n\n", u3.Path)

	fmt.Println("KEY INSIGHT:")
	fmt.Println("  • Host includes port (example.com:8080)")
	fmt.Println("  • Hostname() removes port (example.com)")
	fmt.Println("  • Port() returns port as string")
}

// LESSON 3: Query Parameters (The Important Part)
// ===============================================
func lesson3QueryParameters() {
	fmt.Println("LESSON 3: QUERY PARAMETERS (THE IMPORTANT PART)")
	fmt.Println("----------------------------------------------\n")

	fmt.Println("WHY QUERY PARAMETERS MATTER:")
	fmt.Println("  Query parameters (?key=value&key2=value2) are how:")
	fmt.Println("    • Users search for products")
	fmt.Println("    • Web apps filter and sort data")
	fmt.Println("    • APIs accept configuration\n")

	fmt.Println("THE PROBLEM WITH RawQuery:")
	fmt.Println("  u.RawQuery gives you the raw string: \"id=123&name=John\"")
	fmt.Println("  This is hard to work with - you'd have to parse it manually.\n")

	fmt.Println("THE SOLUTION: u.Query()")
	fmt.Println("  u.Query() parses RawQuery into a map (url.Values)")
	fmt.Println("  url.Values is map[string][]string (values can be lists)\n")

	fmt.Println("EXAMPLE 1: Parsing Query Parameters")
	urlStr := "https://example.com/search?name=John&age=30&hobby=coding&hobby=reading"
	fmt.Printf("URL: %s\n\n", urlStr)

	u, _ := url.Parse(urlStr)

	fmt.Println("RAW (HARD TO USE):")
	fmt.Printf("  u.RawQuery = \"%s\"\n\n", u.RawQuery)

	fmt.Println("PARSED (EASY TO USE):")
	query := u.Query()
	fmt.Printf("  query = %v\n\n", query)

	fmt.Println("ACCESSING VALUES:")
	fmt.Printf("  query.Get(\"name\") = \"%s\"\n", query.Get("name"))
	fmt.Printf("  query.Get(\"age\") = \"%s\"\n", query.Get("age"))
	fmt.Printf("  query.Get(\"hobby\") = \"%s\" (first value)\n", query.Get("hobby"))
	fmt.Printf("  query[\"hobby\"] = %v (all values)\n\n", query["hobby"])

	fmt.Println("EXAMPLE 2: Iterating All Query Parameters")
	fmt.Println("  Code: for key, values := range query {")
	fmt.Println()

	for key, values := range query {
		fmt.Printf("    %s: %v\n", key, values)
	}
	fmt.Println()

	fmt.Println("EXAMPLE 3: Handling Missing Parameters")
	fmt.Printf("  query.Get(\"missing\") = \"%s\" (empty string)\n", query.Get("missing"))
	fmt.Println("    (Get() returns empty string if key not found)\n")

	fmt.Println("KEY INSIGHT:")
	fmt.Println("  • Use .Get() for single values (most common)")
	fmt.Println("  • Use direct map access for multi-value parameters")
	fmt.Println("  • Always check if key exists before using")
}

// LESSON 4: Building URLs (Writing)
// ================================
func lesson4BuildingURLs() {
	fmt.Println("LESSON 4: BUILDING URLs (WRITING)")
	fmt.Println("---------------------------------\n")

	fmt.Println("WHY BUILD URLs DYNAMICALLY?")
	fmt.Println("  • Call third-party APIs with user-specified filters")
	fmt.Println("  • Generate links in your application")
	fmt.Println("  • Redirect to URLs with tracking parameters\n")

	fmt.Println("METHOD 1: THE STRUCT APPROACH (Recommended)")
	fmt.Println("  Build by populating a url.URL struct\n")

	fmt.Println("EXAMPLE:")
	fmt.Println("  Code:")
	fmt.Println("    baseURL := &url.URL{")
	fmt.Println("        Scheme: \"https\",")
	fmt.Println("        Host:   \"api.example.com\",")
	fmt.Println("        Path:   \"/users\",")
	fmt.Println("    }\n")

	baseURL := &url.URL{
		Scheme: "https",
		Host:   "api.example.com",
		Path:   "/users",
	}

	fmt.Println("  Setting query parameters:")
	fmt.Println("    q := baseURL.Query()")
	fmt.Println("    q.Set(\"id\", \"123\")")
	fmt.Println("    q.Set(\"active\", \"true\")")
	fmt.Println("    baseURL.RawQuery = q.Encode()\n")

	q := baseURL.Query()
	q.Set("id", "123")
	q.Set("active", "true")
	baseURL.RawQuery = q.Encode()

	fmt.Printf("  Result: %s\n\n", baseURL.String())

	fmt.Println("METHOD 2: THE url.Values BUILDER (For Query Strings)")
	fmt.Println("  Use url.Values when you just need to build a query string\n")

	fmt.Println("EXAMPLE:")
	fmt.Println("  Code:")
	fmt.Println("    params := url.Values{}")
	fmt.Println("    params.Add(\"search\", \"golang\")")
	fmt.Println("    params.Add(\"sort\", \"date\")")
	fmt.Println("    params.Add(\"limit\", \"10\")\n")

	params := url.Values{}
	params.Add("search", "golang")
	params.Add("sort", "date")
	params.Add("limit", "10")

	fmt.Println("    encoded := params.Encode()")
	encoded := params.Encode()
	fmt.Printf("  Result: {search=golang&sort=date&limit=10}\n")
	fmt.Printf("  Actual: %s\n\n", encoded)

	fmt.Println("COMPLETE URL WITH BUILDER:")
	fullURL := "https://search.example.com/results?" + encoded
	fmt.Printf("  %s\n\n", fullURL)

	fmt.Println("COMPARISON:")
	fmt.Println("  Method 1 (Struct):  Better for complete URLs with multiple parts")
	fmt.Println("  Method 2 (Values):  Better for just building query strings")
}

// LESSON 5: URL Encoding (Safety First)
// ====================================
func lesson5URLEncoding() {
	fmt.Println("LESSON 5: URL ENCODING (SAFETY FIRST)")
	fmt.Println("------------------------------------\n")

	fmt.Println("THE PROBLEM WITH STRING CONCATENATION:")
	fmt.Println("  Never do this: \"?search=\" + userInput")
	fmt.Println("  If userInput contains special chars, the URL breaks!\n")

	fmt.Println("EXAMPLE OF BROKEN URL:")
	badInput := "golang tutorial"
	badURL := "https://example.com/search?q=" + badInput
	fmt.Printf("  Input:   \"%s\"\n", badInput)
	fmt.Printf("  Result:  %s\n", badURL)
	fmt.Println("  Problem: Space breaks the URL!\n")

	fmt.Println("SOLUTION: Use url.Values with .Encode()")
	fmt.Println("  .Encode() converts special characters safely\n")

	fmt.Println("SAFE APPROACH:")
	params := url.Values{}
	params.Add("q", badInput)
	safeQuery := params.Encode()
	safeURL := "https://example.com/search?" + safeQuery

	fmt.Printf("  Input:   \"%s\"\n", badInput)
	fmt.Printf("  Result:  %s\n", safeURL)
	fmt.Println("  ✓ Space converted to %20 or +\n")

	fmt.Println("COMMON CHARACTER ENCODINGS:")
	specialChars := map[string]string{
		"space": " ",
		"&":     "&",
		"=":     "=",
		"+":     "+",
		"#":     "#",
		"%":     "%",
		"?":     "?",
		"@":     "@",
		":":     ":",
		"/":     "/",
	}

	fmt.Println("  Character  →  Encoded As")
	for name, char := range specialChars {
		vals := url.Values{}
		vals.Add("test", char)
		encoded := vals.Encode()
		fmt.Printf("  %-10s →  %s\n", name, encoded)
	}
	fmt.Println()

	fmt.Println("KEY INSIGHT:")
	fmt.Println("  ✓ Always use url.Values.Encode() for user input")
	fmt.Println("  ✓ Never concatenate user input directly into URLs")
	fmt.Println("  ✓ Encoding prevents injection and URL corruption")
}

// LESSON 6: Practical Exercise - API Request Builder
// ==================================================
func lesson6PracticalExercise() {
	fmt.Println("LESSON 6: PRACTICAL EXERCISE - API REQUEST BUILDER")
	fmt.Println("--------------------------------------------------\n")

	fmt.Println("TASK:")
	fmt.Println("  Build a function that constructs a GitHub API URL")
	fmt.Println("  to search for repositories.\n")

	fmt.Println("REQUIREMENTS:")
	fmt.Println("  • Base: https://api.github.com/search/repositories")
	fmt.Println("  • Parameters:")
	fmt.Println("    - q: search query (required)")
	fmt.Println("    - sort: (optional, default: stars)")
	fmt.Println("    - order: asc or desc (optional)\n")

	fmt.Println("IMPLEMENTATION:\n")

	// Function to build GitHub search URL
	buildGitHubSearchURL := func(query string, sort string, order string) string {
		u := &url.URL{
			Scheme: "https",
			Host:   "api.github.com",
			Path:   "/search/repositories",
		}

		q := u.Query()
		q.Set("q", query)
		if sort != "" {
			q.Set("sort", sort)
		}
		if order != "" {
			q.Set("order", order)
		}

		u.RawQuery = q.Encode()
		return u.String()
	}

	fmt.Println("  func buildGitHubSearchURL(query, sort, order string) string {")
	fmt.Println("      u := &url.URL{")
	fmt.Println("          Scheme: \"https\",")
	fmt.Println("          Host:   \"api.github.com\",")
	fmt.Println("          Path:   \"/search/repositories\",")
	fmt.Println("      }")
	fmt.Println("      q := u.Query()")
	fmt.Println("      q.Set(\"q\", query)")
	fmt.Println("      if sort != \"\" { q.Set(\"sort\", sort) }")
	fmt.Println("      if order != \"\" { q.Set(\"order\", order) }")
	fmt.Println("      u.RawQuery = q.Encode()")
	fmt.Println("      return u.String()")
	fmt.Println("  }\n")

	fmt.Println("EXAMPLES:\n")

	// Example 1
	url1 := buildGitHubSearchURL("golang", "stars", "desc")
	fmt.Println("Search for \"golang\" repos, sorted by stars (descending):")
	fmt.Printf("  %s\n\n", url1)

	// Example 2
	url2 := buildGitHubSearchURL("web framework", "forks", "asc")
	fmt.Println("Search for \"web framework\", sorted by forks (ascending):")
	fmt.Printf("  %s\n\n", url2)

	// Example 3 - with special characters
	url3 := buildGitHubSearchURL("machine learning", "", "")
	fmt.Println("Search for \"machine learning\" with no additional sort:")
	fmt.Printf("  %s\n\n", url3)

	fmt.Println("KEY LESSONS:")
	fmt.Println("  ✓ Build URLs using url.URL struct")
	fmt.Println("  ✓ Add query parameters with u.Query()")
	fmt.Println("  ✓ Always use q.Encode() for safe encoding")
	fmt.Println("  ✓ Special characters are handled automatically")
	fmt.Println("  ✓ Use .String() to get the final URL")
}
