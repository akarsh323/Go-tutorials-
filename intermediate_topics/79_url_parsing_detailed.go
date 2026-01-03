package intermediate

import (
	"fmt"
	"net/url"
)

// Topic 79: URL Parsing
// Parsing and building URLs programmatically

func main() {

	// Parsing URLs
	urlStr := "https://example.com:8080/path?name=John&age=30#section"
	u, err := url.Parse(urlStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("Host:", u.Host)
	fmt.Println("Hostname:", u.Hostname())
	fmt.Println("Port:", u.Port())
	fmt.Println("Path:", u.Path)
	fmt.Println("Query:", u.RawQuery)
	fmt.Println("Fragment:", u.Fragment)

	// Parsing query parameters
	query := u.Query()
	fmt.Println("\nQuery Parameters:")
	for key, values := range query {
		fmt.Printf("%s: %v\n", key, values)
	}

	// Building URLs
	baseURL := &url.URL{
		Scheme: "https",
		Host:   "example.com",
		Path:   "/path",
	}

	query2 := baseURL.Query()
	query2.Set("name", "John")
	query2.Set("age", "30")
	baseURL.RawQuery = query2.Encode()

	fmt.Println("\nBuilt URL:", baseURL.String())

}
