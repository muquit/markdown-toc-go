# Introduction

`markdown-toc-go` is a simple multi-platform program to generate table of contents for markdown files. 
I use it to generate TOC for README.md for my projects in github. 
I know github shows TOC of README.md but you have
to click on the list icon to show. I want to see TOC at the top of my README.

# History
I used to use
[markdown_helper](https://github.com/BurdetteLamar/markdown_helper), a ruby gem to generate TOC
for my markdown files but at times I don't have ruby installed in the systems I'm working on. 
It also has a bug, for example, say if inside a code block has a commennt with #, it interprets 
it as header and adds in TOC. Also, project does not seem to be active.

Hope you find this program useful.

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
  -i string
    	Input markdown file (required)
  -input string
    	Input markdown file (required)
  -no-credit
    	Don't add credit line at the end of the file
  -o string
    	Output file (default: input-with-toc.md)
  -output string
    	Output file (default: input-with-toc.md)
  -t string
    	Title for table of contents (default "## Table Of Contents")
  -title string
    	Title for table of contents (default "## Table Of Contents")

Version: 1.0.1
```

# Features

- Automatically extracts headings from markdown files
- Generates GitHub-style anchor links
- Creates a properly indented TOC with bullet points
- Produces a new file with the TOC added at the top

# Version
The current version is 1.0.1

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

# Examples

- Basic usage. his will generate a new file named `README-with-toc.md` containing the original content plus a table of contents at the top.
```
markdown-toc-go -i README.md
```

- Keep README.md in say ./docs directory and generate README.md with TOC in 
the current working directory and overwrite README.md forefully.

```
markdown-toc-go -i docs/README.md -o ./README.md -f
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

- By default, a credit line is added at the end of the document. If you do not want to 
show the credit line, run with the option
```
-no-credit
```


etc.

# Authors

Developed with Claude AI 3.7 Sonnet, working under my guidance and instructions.

# License

This project is licensed under the MIT License - see the LICENSE file for details.
