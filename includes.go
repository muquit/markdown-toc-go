// includes.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

// expandFileInclusions processes the main markdown content to expand file inclusion directives.
// It takes the initial content, the directory of the original input file (for path resolution),
// and the glossary map.
func expandFileInclusions(content, baseDir string, glossary map[string]string) (string, error) {
	// Regular expression to find @[:markdown](filename) directives
	// It captures the filename inside the parentheses.
	includeRegex := regexp.MustCompile(`@\[:markdown\]\(([^)]+)\)`)

	// Use ReplaceAllStringFunc to process each match
	expandedContent := includeRegex.ReplaceAllStringFunc(content, func(match string) string {
		// Extract the filename from the captured group
		submatches := includeRegex.FindStringSubmatch(match)
		if len(submatches) < 2 {
			// This should not happen if the regex matches, but as a safeguard
			fmt.Fprintf(os.Stderr, "Warning: Could not extract filename from inclusion directive: %s\n", match)
			return match // Return original match on error
		}
		includedFilename := submatches[1]

		// Resolve the full path of the included file
		includedFilePath := filepath.Join(baseDir, includedFilename)

		// Read the content of the included file
		includedFileBytes, err := os.ReadFile(includedFilePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to read included file '%s': %v\n", includedFilePath, err)
			return fmt.Sprintf("", includedFilename, err)
		}
		includedFileContent := string(includedFileBytes)

		// Expand glossary terms within the included file's content
		if glossary != nil {
			includedFileContent = expandGlossary(includedFileContent, glossary)
		}

		return includedFileContent
	})

	return expandedContent, nil
}
