// output.go
package main

import (
	"bufio"
	//"io"
	"os"
	"strings"
)

// writeOutputFileWithContent creates the output file with TOC and provided content
// This function takes the already processed markdownContent string.
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
// NOTE: This function is largely redundant with writeOutputFileWithContent
// after the main content is read into a string. We will consolidate this
// in a later cleanup. For now, it's moved here to keep original behavior.
/*
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
*/
