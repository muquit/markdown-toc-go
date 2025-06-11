package main

/////////////////////////////////////////////////////////////////////
// Generate TOC from github README.md. I used to use markdown_helper
// ruby gem but I don't have ruby installed on many systems.
// markdown_help also have some bugs e.g. if some is inside a code
// block and have for ecample #/bin/bash it interprets it as a markdown
// header.
// Developed with Claude AI 3.7/4 Sonnet, working under my guidance and
// instructions.
// Enhanced with glossary support for key-value expansion.
// Mar-28-2025
/////////////////////////////////////////////////////////////////////

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Version information
const (
	appName    = "markdown-toc-go"
	appVersion = "1.0.3"
)

// Command line flags
var (
	inputFile      string
	outputFile     string
	maxDepth       int
	tocTitle       string
	forceWrite     bool
	noCredit       bool
	creditText     string
	preTocContent  string
	preTocFile     string
	insertPosition string
	glossaryFile   string
)

func main() {
	// Parse command line flags
	flag.StringVar(&inputFile, "i", "", "Input markdown file (required)")
	flag.StringVar(&inputFile, "input", "", "Input markdown file (required)")

	flag.StringVar(&outputFile, "o", "", "Output file (default: input-with-toc.md)")
	flag.StringVar(&outputFile, "output", "", "Output file (default: input-with-toc.md)")

	flag.IntVar(&maxDepth, "d", 6, "Maximum heading depth to include (1-6)")
	flag.IntVar(&maxDepth, "depth", 6, "Maximum heading depth to include (1-6)")

	flag.StringVar(&tocTitle, "t", "## Table Of Contents", "Title for table of contents")
	flag.StringVar(&tocTitle, "title", "## Table Of Contents", "Title for table of contents")

	flag.BoolVar(&forceWrite, "f", false, "Overwrite output file if it exists")
	flag.BoolVar(&forceWrite, "force", false, "Overwrite output file if it exists")

	flag.BoolVar(&noCredit, "no-credit", false, "Don't add credit line at the end of the file")

	flag.StringVar(&preTocContent, "pre-toc", "", "Content to insert before TOC (badges, etc.)")
	flag.StringVar(&preTocFile, "pre-toc-file", "", "File containing content to insert before TOC")
	flag.StringVar(&insertPosition, "insert-position", "before-title", "Where to insert pre-TOC content: 'before-title', 'after-title'")

	flag.StringVar(&glossaryFile, "glossary", "", "Glossary file for key-value expansion (format: KEY    value)")

	currentTime := time.Now().Format("Jan-02-2006")
	creditText = "---\n<sub>TOC is created by https://github.com/muquit/markdown-toc-go on " + currentTime + "</sub>"

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s - Generate a table of contents for markdown files\n\n", appName)
		fmt.Fprintf(os.Stderr, "Usage: %s -i input.md [-o output.md] [options]\n\n", appName)
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nGlossary format: KEY<spaces>value (e.g., API_VERSION    v1.2.3)\n")
		fmt.Fprintf(os.Stderr, "Glossary syntax: @KEY@ in markdown will be replaced with value\n")
		fmt.Fprintf(os.Stderr, "\nVersion: %s\n", appVersion)
	}

	flag.Parse()

	// Check if input file is provided
	if inputFile == "" {
		fmt.Fprintln(os.Stderr, "Error: Input file is required")
		flag.Usage()
		os.Exit(1)
	}

	// Check if input file exists
	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: Input file '%s' does not exist\n", inputFile)
		os.Exit(1)
	}

	// Check if glossary file exists (if specified)
	if glossaryFile != "" {
		if _, err := os.Stat(glossaryFile); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error: Glossary file '%s' does not exist\n", glossaryFile)
			os.Exit(1)
		}
	}

	// Set default output file if not provided
	if outputFile == "" {
		ext := filepath.Ext(inputFile)
		baseName := strings.TrimSuffix(inputFile, ext)
		outputFile = baseName + "-with-toc" + ext
	}

	// Check if output file already exists
	if _, err := os.Stat(outputFile); err == nil && !forceWrite {
		fmt.Fprintf(os.Stderr, "Error: Output file '%s' already exists. Use -f to overwrite\n", outputFile)
		os.Exit(1)
	}

	// Validate max depth
	if maxDepth < 1 || maxDepth > 6 {
		fmt.Fprintln(os.Stderr, "Error: Depth must be between 1 and 6")
		os.Exit(1)
	}

	// Validate insert position
	if insertPosition != "before-title" && insertPosition != "after-title" {
		fmt.Fprintln(os.Stderr, "Error: insert-position must be 'before-title' or 'after-title'")
		os.Exit(1)
	}

	// Load glossary if specified
	var glossary map[string]string
	if glossaryFile != "" {
		var err error
		glossary, err = loadGlossary(glossaryFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to load glossary file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Loaded %d glossary entries from '%s'\n", len(glossary), glossaryFile)
	}

	// Read pre-TOC content from file if specified
	if preTocFile != "" {
		if preTocContent != "" {
			fmt.Fprintln(os.Stderr, "Warning: Both pre-toc and pre-toc-file specified. Using pre-toc-file.")
		}

		content, err := os.ReadFile(preTocFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: Failed to read pre-TOC file: %v\n", err)
			os.Exit(1)
		}
		preTocContent = string(content)
	}

	// Read input file content
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to read input file: %v\n", err)
		os.Exit(1)
	}

	markdownContent := string(content)

	// Expand glossary terms if glossary is provided
	if glossary != nil {
		markdownContent = expandGlossary(markdownContent, glossary)
	}

	// Process the markdown content for headings
	headings, err := extractHeadingsFromContent(markdownContent, maxDepth)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to extract headings: %v\n", err)
		os.Exit(1)
	}

	// Generate TOC
	toc := generateTOC(headings)

	// Expand glossary in pre-TOC content if glossary is provided
	if glossary != nil && preTocContent != "" {
		preTocContent = expandGlossary(preTocContent, glossary)
	}

	// Expand glossary in TOC title if glossary is provided
	if glossary != nil {
		tocTitle = expandGlossary(tocTitle, glossary)
	}

	// Create output file with TOC
	err = writeOutputFileWithContent(markdownContent, outputFile, toc, preTocContent)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to write output file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated TOC in '%s'\n", outputFile)
}

// Heading represents a markdown heading
type Heading struct {
	Level  int    // Heading level (1-6)
	Text   string // Heading text
	Anchor string // GitHub-style anchor link
}

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
		indent := strings.Repeat("  ", h.Level-1)

		// Create TOC entry with link
		sb.WriteString(fmt.Sprintf("%s- [%s](#%s)\n", indent, h.Text, h.Anchor))
	}

	return sb.String()
}

// writeOutputFileWithContent creates the output file with TOC and provided content
func writeOutputFileWithContent(content, outputPath, toc, preTocContent string) error {
	// Create output file for writing
	output, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer output.Close()

	writer := bufio.NewWriter(output)

	// Insert pre-TOC content if specified and position is before title
	if preTocContent != "" && insertPosition == "before-title" {
		_, err = writer.WriteString(preTocContent)
		if err != nil {
			return err
		}

		// Add a newline if not already present
		if !strings.HasSuffix(preTocContent, "\n") {
			_, err = writer.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}

	// Write TOC title
	_, err = writer.WriteString(tocTitle + "\n")
	if err != nil {
		return err
	}

	// Insert pre-TOC content if specified and position is after title
	if preTocContent != "" && insertPosition == "after-title" {
		_, err = writer.WriteString(preTocContent)
		if err != nil {
			return err
		}

		// Add a newline if not already present
		if !strings.HasSuffix(preTocContent, "\n") {
			_, err = writer.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}

	// Write TOC content
	_, err = writer.WriteString(toc + "\n")
	if err != nil {
		return err
	}

	// Write original content (now expanded)
	_, err = writer.WriteString(content)
	if err != nil {
		return err
	}

	// Add credit line at the end if not disabled
	if !noCredit {
		// Add a newline before credit if the content doesn't end with one
		if len(content) > 0 && content[len(content)-1] != '\n' {
			_, err = writer.WriteString("\n")
			if err != nil {
				return err
			}
		}

		// Write credit text
		_, err = writer.WriteString("\n" + creditText + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

// writeOutputFile creates the output file with TOC and original content
func writeOutputFile(inputPath, outputPath, toc, preTocContent string) error {
	// Open input file for reading
	input, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer input.Close()

	// Read all content from input file
	content, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	// Create output file for writing
	output, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer output.Close()

	writer := bufio.NewWriter(output)

	// Insert pre-TOC content if specified and position is before title
	if preTocContent != "" && insertPosition == "before-title" {
		_, err = writer.WriteString(preTocContent)
		if err != nil {
			return err
		}

		// Add a newline if not already present
		if !strings.HasSuffix(preTocContent, "\n") {
			_, err = writer.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}

	// Write TOC title
	_, err = writer.WriteString(tocTitle + "\n")
	if err != nil {
		return err
	}

	// Insert pre-TOC content if specified and position is after title
	if preTocContent != "" && insertPosition == "after-title" {
		_, err = writer.WriteString(preTocContent)
		if err != nil {
			return err
		}

		// Add a newline if not already present
		if !strings.HasSuffix(preTocContent, "\n") {
			_, err = writer.WriteString("\n")
			if err != nil {
				return err
			}
		}
	}

	// Write TOC content
	_, err = writer.WriteString(toc + "\n")
	if err != nil {
		return err
	}

	// Write original content
	_, err = writer.Write(content)
	if err != nil {
		return err
	}

	// Add credit line at the end if not disabled
	if !noCredit {
		// Add a newline before credit if the file doesn't end with one
		if len(content) > 0 && content[len(content)-1] != '\n' {
			_, err = writer.WriteString("\n")
			if err != nil {
				return err
			}
		}

		// Write credit text
		_, err = writer.WriteString("\n" + creditText + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
