# Main Title of the Document

This is a test document to verify that markdown-toc-go correctly handles various markdown elements.

## Introduction

This document contains various headings and code blocks to test the TOC generation.

### Purpose of Testing

We need to ensure that the TOC generator works properly, especially with code blocks.

## Code Blocks

Let's test how the program handles code blocks with comments that look like headings:

```bash
# This is a comment in bash, not a heading
echo "Hello World"
# Another comment that should not appear in TOC
```

```python
# Python comment that looks like a heading
def hello():
    # Indented comment
    print("Hello from Python")
```

## Configurations

### Go Configuration

Here's a Go configuration file:

```go
// This is a Go comment
package main

// # This line has a hash in a comment
func main() {
    // Print something
    fmt.Println("Hello, Go!")
}
```

### Python Configuration

```python
# CONFIGURATION SETTINGS
# Set DEBUG=True for development
DEBUG = True
# DATABASE settings below
DATABASE = {
    "host": "localhost"
}
```

## Edge Cases

### Inline Code with Hash

What about inline code like `# not a heading` or text with # symbols in it?

### Fenced Code Block with Different Marker

We should also test with different fence markers:

```
# This is inside a code block without language specification
function test() {
    # Another comment
    return true;
}
```

## Tables

| Heading 1 | Heading 2 | Heading 3 |
|-----------|-----------|-----------|
| # Not a heading | Value 2 | Value 3 |
| Value 4 | # Also not a heading | Value 6 |

## Nested Lists

1. First item
   - # Not a heading inside a list
   - Sub item 2
2. Second item
   ```
   # Comment in code block inside a list
   ```

## Other Elements

- Bullet list
- With multiple items

### Final Section

This is the end of the test document.

#### A Level 4 Heading

Testing deeper levels.

##### Level 5

Even deeper.

###### Level 6

The deepest level.

# Another Main Heading

Just to make sure we capture all headings.

## Conclusion

If the program works correctly, the TOC should only include the actual headings from this document, not any text that looks like headings inside code blocks, tables, or lists.
