# Regex Guide - Complete Documentation

## Summary

I've created **two comprehensive regex guides** for Go following your exact specifications:

### 1. **73_regex_comprehensive.go** (New File - Perfect Format)
- **TL;DR**: One paragraph summary at the top explaining what regex is and key functions
- **6 Major Parts**:
  1. **PART 1: BASICS** - MatchString (checking pattern existence)
  2. **PART 2: FINDING** - FindString / FindAllString (extracting patterns)
  3. **PART 3: REPLACING** - ReplaceAllString / ReplaceAllStringFunc (modifying patterns)
  4. **PART 4: CAPTURING** - FindAllStringSubmatch (extracting groups)
  5. **PART 5: PRACTICAL EXAMPLES** - Real-world use cases
  6. **PART 6: PERFORMANCE & BEST PRACTICES** - Optimization tips

- **Structure for Each Part**:
  - 📌 CONCEPT section (explains what and why)
  - 📝 CODE EXAMPLES (fenced code blocks with language hints)
  - 💡 LINE-BY-LINE EXPLANATION
  - 🔄 LIVE EXECUTION (runnable examples)
  - ✅ KEY TAKEAWAY

- **Copy-Paste Ready Patterns** (at the end)

### 2. **73_regex_detailed.go** (Updated File - Enhanced Format)
- **5 Organized Sections**:
  1. Pattern Matching (MatchString)
  2. Finding Patterns (FindString / FindAllString)
  3. Replacing Patterns (ReplaceAllString / ReplaceAllStringFunc)
  4. Capturing Groups (FindAllStringSubmatch)
  5. Real-World Examples (validation, extraction, censoring)

- **Features**:
  - Clear status indicators (✓ ✗) for test results
  - Organized examples with explanations
  - Real-world use cases (email validation, URL extraction, hashtag extraction, password validation, phone number censoring)
  - Error handling with Compile vs MustCompile
  - Line-by-line code walkthroughs

## Key Takeaways (From the Guides)

1. **MatchString** - Boolean check if pattern exists anywhere in text
2. **FindString** - Returns first match as string (or empty string if none)
3. **FindAllString** - Returns all matches with optional limit parameter
4. **ReplaceAllString** - Static replacement of all matches
5. **ReplaceAllStringFunc** - Dynamic replacement using a function
6. **FindAllStringSubmatch** - Extract capturing groups: `matches[i][0]` = full, `matches[i][j]` = group j

## Copy-Paste Regex Patterns

```
Email (simple):                  [a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z]{2,}
Phone (XXX-XXX-XXXX):           \d{3}-\d{3}-\d{4}
Date (YYYY-MM-DD):              \d{4}-\d{2}-\d{2}
URL (http/https):               https?://[a-zA-Z0-9.-]+
Digits only:                    ^[0-9]+$
Lowercase letters only:         ^[a-z]+$
Uppercase letters only:         ^[A-Z]+$
Word characters:                ^\w+$
Hashtag:                        #\w+
Mention (@username):            @\w+
```

## Best Practices Highlighted

1. **Compile Once, Reuse Many**
   - Don't compile regex inside loops
   - Parse once at startup, use repeatedly

2. **Use Raw Strings**
   - Use backticks to avoid double-escaping
   - `regexp.MustCompile(`\d{3}`)`  instead of `regexp.MustCompile("\\d{3}")`

3. **Use Anchors**
   - `^` and `$` for full-string matching
   - `\b` for word boundaries

4. **Error Handling**
   - Use `MustCompile` for compile-time patterns
   - Use `Compile` for runtime patterns (returns error)

5. **Performance**
   - Parsing is expensive (slow)
   - Matching is cheap (fast)
   - Always compile before loops

## Files Created/Updated

- ✅ `/Users/akarsh/GOTUT/intermediate_topics/73_regex_comprehensive.go` (NEW - 716 lines)
- ✅ `/Users/akarsh/GOTUT/intermediate_topics/73_regex_detailed.go` (UPDATED - 434 lines)
- ✅ Committed and pushed to main branch on GitHub

## Next Steps for Learners

1. **Fix**: Understand why double escaping isn't needed with raw strings
2. **Test**: Run the live examples to see regex patterns in action
3. **Improve**: Extend the practical examples with your own use cases
