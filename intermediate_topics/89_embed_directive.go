package main

import "fmt"

// 89 Embed Directive
// Note: using go:embed requires a package-level directive and an existing file.
// This demo explains usage only (comment) to avoid needing extra files.
func Demo89EmbedDirective() {
	fmt.Println("-- 89 Embed Directive --")
	fmt.Println("Use //go:embed <file> at package level and an embed.FS or string variable.")
}
