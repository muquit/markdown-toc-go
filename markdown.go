// markdown.go
package main

import (
	"fmt"
	//"os"
	"bufio"
	"regexp"
	"strings"
)

// Heading represents a markdown heading
type Heading struct {
	Level  int    // Heading level (1-6)
	Text   string // Heading text
	Anchor string // GitHub-style anchor link
}

// extractHeadingsFromContent extracts headings from markdown content string
func extractHeadingsFromContent(content string, maxDepth int) ([]Heading, error) {
	var headings []Heading
	scanner := bufio.NewScanner(strings.NewReader(content))

	// Regular expression to match markdown headings
	headingRegex := regexp.MustCompile(`^(#{1,6})\s+(.+)$`)

	// Flag to track if we're inside a code block
	insideCodeBlock := false

	// Pattern to detect code block markers (three backticks)
	codeBlockMarker := "```"

	for scanner.Scan() {
		line := scanner.Text()

		// Check if this line is a code block marker
		if strings.HasPrefix(strings.TrimSpace(line), codeBlockMarker) {
			insideCodeBlock = !insideCodeBlock
			continue
		}

		// Skip heading detection if inside code block
		if insideCodeBlock {
			continue
		}

		match := headingRegex.FindStringSubmatch(line)

		if match != nil {
			level := len(match[1])

			// Skip if heading level is greater than maxDepth
			if level > maxDepth {
				continue
			}

			text := strings.TrimSpace(match[2])

			// Generate GitHub-style anchor
			anchor := generateAnchor(text)

			headings = append(headings, Heading{
				Level:  level,
				Text:   text,
				Anchor: anchor,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return headings, nil
}

// extractHeadings reads the input file and extracts all headings up to maxDepth
// NOTE: This function is currently redundant as extractHeadingsFromContent is used
// after reading the entire file. We will remove this in a later cleanup step.
// For now, it's moved here to keep the original behavior intact.
/*
func extractHeadings(filePath string, maxDepth int) ([]Heading, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read lines and extract headings
	var headings []Heading
	scanner := bufio.NewScanner(file)

	// Regular expression to match markdown headings
	headingRegex := regexp.MustCompile(`^(#{1,6})\s+(.+)$`)

	// Flag to track if we're inside a code block
	insideCodeBlock := false

	// Pattern to detect code block markers (three backticks)
	codeBlockMarker := "```"

	for scanner.Scan() {
		line := scanner.Text()

		// Check if this line is a code block marker
		if strings.HasPrefix(strings.TrimSpace(line), codeBlockMarker) {
			insideCodeBlock = !insideCodeBlock
			continue
		}

		// Skip heading detection if inside code block
		if insideCodeBlock {
			continue
		}

		match := headingRegex.FindStringSubmatch(line)

		if match != nil {
			level := len(match[1])

			// Skip if heading level is greater than maxDepth
			if level > maxDepth {
				continue
			}

			text := strings.TrimSpace(match[2])

			// Generate GitHub-style anchor
			anchor := generateAnchor(text)

			headings = append(headings, Heading{
				Level:  level,
				Text:   text,
				Anchor: anchor,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return headings, nil
}
*/

// generateAnchor creates a GitHub-style anchor link from heading text
func generateAnchor(text string) string {
	// Convert to lowercase
	anchor := strings.ToLower(text)

	// Replace spaces with hyphens
	anchor = strings.ReplaceAll(anchor, " ", "-")

	// Remove characters that are not alphanumeric, dash, or underscore
	re := regexp.MustCompile(`[^a-z0-9\-_]`)
	anchor = re.ReplaceAllString(anchor, "")

	// Replace multiple consecutive dashes with a single dash
	re = regexp.MustCompile(`-+`)
	anchor = re.ReplaceAllString(anchor, "-")

	// Remove leading and trailing dashes
	anchor = strings.Trim(anchor, "-")

	return anchor
}

// generateTOC creates a markdown TOC from the list of headings
func generateTOC(headings []Heading) string {
	if len(headings) == 0 {
		return "No headings found in the document."
	}

	var sb strings.Builder

	// Add TOC entries
	for _, h := range headings {
		// Create proper indentation
		indent := strings.Repeat("  ", h.Level-1) // Use 2 spaces for consistency

		// Create TOC entry with link
		sb.WriteString(fmt.Sprintf("%s- [%s](#%s)\n", indent, h.Text, h.Anchor))
	}

	return sb.String()
}
