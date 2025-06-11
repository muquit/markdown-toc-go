# Introduction

# Introduction
`markdown-toc-go` is a simple multi-platform program to generate table of 
contents for markdown files and expand glossary placeholders for 
consistent documentation.  I use it to generate TOC for README.md for my 
projects in github.  I know github shows TOC of README.md but you 
have to click on the list icon to show. I want to see TOC at the top of my 
README.

It's also a lot of typing to add link and image tags while creating README.
md files. The glossary feature can cut down a lot of repetitive and 
painful typing by allowing you to define reusable content once and 
reference it throughout your documentation.  A single global glossary file 
can be shared across multiple projects for consistent branding, URLs, and 
common references.

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

Version: 1.0.3
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

# Version
The current version is 1.0.3

Please look at [ChangeLog](ChangeLog.md) for what has changed in the current version.

# Quick Start
Install [go](https://go.dev/) first

```bash
go install github.com/muquit/markdown-toc-go@latest
```
# Download

Download pre-compiled binaries from
[Releases](https://github.com/muquit/markdown-toc-go/releases) page

# Bulding from source
Install [go](https://go.dev/) first

```
git clone https://github.com/yourusername/markdown-toc-go.git
cd markdown-toc-go
go build
```
Look at `Makefile` for more info.

(Optional) Move the binary to your PATH:
```
sudo mv markdown-toc-go /usr/local/bin/
```

# How I use it

I keep my README.md in `./docs/` directory and generate the README.md with TOC from it. 

# Examples

- Basic usage. his will generate a new file named `README-with-toc.md` containing the original content plus a table of contents at the top.
```
markdown-toc-go -i README.md -glossary glossary.txt
```

- Keep README.md in say ./docs directory and generate README.md with TOC in 
the current working directory and overwrite README.md forcefully.

```
markdown-toc-go -i docs/README.md -o ./README.md -glossary glossary.txt -f
```
- Generate TOC with custom output file:
```
markdown-toc-go -i README.md -o README_with_toc.md
```
- Generate TOC including only level 1 and 2 headings:
```
markdown-toc-go -i README.md -d 2
```

- Use a custom title for the TOC:
```
markdown-toc-go -i README.md -t "## Contents"
```

- Force overwrite existing output file:
```
markdown-toc-go -i README.md -o README_with_toc.md -f
```

- Generate TOC of the test document:
```
markdown-toc-go -i test/TEST.md -o ./TEST.md
```
- Generate TOC of the test document by expanding glossary key value
```
markdown-toc-go -i test/TEST_GLOSSARY.md -o ./TEST_GLOSSARY_EXPANDED.md
```

- By default, a credit line is added at the end of the document. If you do not want to 
show the credit line, run with the option
```
-no-credit
```


etc.

# Authors

Developed with @CLAUDE@, working under my guidance and instructions.

# License

This project is licensed under the MIT License - see the LICENSE file for details.
