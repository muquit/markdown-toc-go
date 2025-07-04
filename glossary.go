// glossary.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// loadGlossary reads a glossary file and returns a map of key-value pairs
func loadGlossary(filePath string) (map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	glossary := make(map[string]string)
	scanner := bufio.NewScanner(file)
	lineNum := 0

	// Regular expression to match key-value pairs: KEY<spaces>VALUE
	keyValueRegex := regexp.MustCompile(`^([A-Za-z_][A-Za-z0-9_]*)\s+(.+)$`)

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comments
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		match := keyValueRegex.FindStringSubmatch(line)
		if match != nil {
			key := match[1]
			value := match[2]
			glossary[key] = value
		} else {
			fmt.Fprintf(os.Stderr, "Warning: Invalid glossary format at line %d: %s\n", lineNum, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Expand glossary values that reference other keys
	expandGlossaryValues(glossary)

	return glossary, nil
}

// expandGlossaryValues expands @KEY@ references within glossary values
func expandGlossaryValues(glossary map[string]string) {
	const maxPasses = 5
	keyRegex := regexp.MustCompile(`@([A-Za-z_][A-Za-z0-9_]*)@`)

	for pass := 0; pass < maxPasses; pass++ {
		expanded := false

		for key, value := range glossary {
			newValue := keyRegex.ReplaceAllStringFunc(value, func(match string) string {
				// Extract key name from @KEY@
				refKey := match[1 : len(match)-1]

				// Prevent self-reference
				if refKey == key {
					fmt.Fprintf(os.Stderr, "Warning: Circular reference detected for key '%s'\n", key)
					return match
				}

				// Look up the referenced key
				if refValue, exists := glossary[refKey]; exists {
					expanded = true
					return refValue
				} else {
					fmt.Fprintf(os.Stderr, "Warning: Unknown key '@%s@' referenced in glossary value for '%s'\n", refKey, key)
					return match
				}
			})

			if newValue != value {
				glossary[key] = newValue
			}
		}

		// If no expansions were made in this pass, we're done
		if !expanded {
			break
		}
	}
}

// expandGlossary replaces @KEY@ placeholders in content with glossary values
func expandGlossary(content string, glossary map[string]string) string {
	keyRegex := regexp.MustCompile(`@([A-Za-z_][A-Za-z0-9_]*)@`)

	return keyRegex.ReplaceAllStringFunc(content, func(match string) string {
		// Extract key name from @KEY@
		key := match[1 : len(match)-1]

		// Look up the key in glossary
		if value, exists := glossary[key]; exists {
			return value
		} else {
			fmt.Fprintf(os.Stderr, "Warning: Unknown key '@%s@' found in content\n", key)
			return match
		}
	})
}