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
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// Version information
const (
	appName    = "markdown-toc-go"
	appVersion = "1.0.4"
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

    // --- NEW CODE STARTS HERE ---
    // Get the directory of the input file for resolving included paths
    inputDir := filepath.Dir(inputFile)

    // Expand file inclusions
    markdownContent, err = expandFileInclusions(markdownContent, inputDir, glossary)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: Failed to expand file inclusions: %v\n", err)
        os.Exit(1)
    }
    // --- NEW CODE ENDS HERE ---

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

