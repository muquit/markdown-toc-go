## Table Of Contents
- [Introduction](#introduction)
- [History](#history)
- [Synopsis](#synopsis)
- [Features](#features)
  - [Glossary Expansion](#glossary-expansion)

# Introduction
`markdown-toc-go` is a simple multi-platform program to generate table of 
contents for markdown files and expand glossary placeholders for 
consistent documentation.  I use it to generate TOC for `README.md` for my 
projects in github.  I know github shows TOC of README.md but you 
have to click on the list icon to show. I want to see TOC at the top of my 
README.

It's also a lot of typing to add link and image tags while creating README.
md files. The glossary feature can cut down a lot of repetitive and 
painful typing by allowing you to define reusable content once and 
reference it throughout your documentation.  A single global glossary file 
can be shared across multiple projects for consistent branding, URLs, and 
common references.

[AsciiDoc](https://asciidoc.org/) has similar feature but it comes with the whole toolchain baggage -
I do use [AsciiDoc](https://asciidoc.org/) documents.  `markdown-toc-go` is just one 
static binary and you're done.

Hope you find this program useful.

# History
I used to use
[markdown_helper](https://github.com/BurdetteLamar/markdown_helper), a ruby gem to generate TOC
for my markdown files but at times I don't have ruby installed in the systems I'm working on. 
It also has a bug, for example, say if inside a code block has a commennt with #, it interprets 
it as header and adds in TOC. Also, project does not seem to be active.

# Synopsis
```
markdown-toc-go - Generate a table of contents for markdown files

Usage: markdown-toc-go -i input.md [-o output.md] [options]

Options:
  -d int
    	Maximum heading depth to include (1-6) (default 6)
  -depth int
    	Maximum heading depth to include (1-6) (default 6)
  -f	Overwrite output file if it exists
  -force
    	Overwrite output file if it exists
  -glossary string
    	Glossary file for key-value expansion (format: KEY    value)
  -i string
    	Input markdown file (required)
  -input string
    	Input markdown file (required)
  -insert-position string
    	Where to insert pre-TOC content: 'before-title', 'after-title' (default "before-title")
  -no-credit
    	Don't add credit line at the end of the file
  -o string
    	Output file (default: input-with-toc.md)
  -output string
    	Output file (default: input-with-toc.md)
  -pre-toc string
    	Content to insert before TOC (badges, etc.)
  -pre-toc-file string
    	File containing content to insert before TOC
  -t string
    	Title for table of contents (default "## Table Of Contents")
  -title string
    	Title for table of contents (default "## Table Of Contents")

Glossary format: KEY<spaces>value (e.g., API_VERSION    v1.2.3)
Glossary syntax: @KEY@ in markdown will be replaced with value

Version: 1.0.4

```
# Features
- Automatically extracts headings from markdown files
- Generates GitHub-style anchor links
- Creates a properly indented TOC with bullet points
- Produces a new file with the TOC added at the top
- **Glossary support**: Expand `@KEY@` placeholders with key-value pairs from external file
- **Nested references**: Glossary values can reference other glossary keys
- **Rich content**: Supports links, images, formatting, and code in glossary values
- Configurable TOC depth (1-6 heading levels)
- Pre-TOC content insertion with positioning options
- Code block detection to avoid parsing headers inside code blocks
- Overwrite protection with force option

## Glossary Expansion

Writing documentation often involves typing the same content repeatedly - 
company names, URLs, version numbers, installation commands, and formatted 
links. This becomes tedious and error-prone in large README files.

**Common pain points:**
- Copying the same long URLs multiple times
- Inconsistent formatting of company names or product names
- Updating version numbers scattered throughout the document
- Retyping complex installation commands or code snippets
- Managing badges, links, and images that appear multiple times

**Glossary expansion solves this by:**
- **Write once, use everywhere** - Define content once in a glossary file, reference it with `@KEY@` syntax
- **Consistent formatting** - Company names, links, and formatting stay uniform across the entire document
- **Easy updates** - Change a URL or version number in one place, it updates everywhere automatically
- **Reduced errors** - No more typos from retyping the same content
- **Faster writing** - Focus on content instead of repetitive typing

A single global glossary file can be shared across multiple projects for 
consistent branding and common references.

**Note:** Glossary keys are not expanded inside code blocks (``` or `` ` ``) 
to preserve literal code content and examples.





---
<sub>TOC is created by https://github.com/muquit/markdown-toc-go on Jul-04-2025</sub>
