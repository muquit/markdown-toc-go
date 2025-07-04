## Table Of Contents
- [Introduction](#introduction)
- [History](#history)
- [Synopsis](#synopsis)
- [Features](#features)
  - [Glossary Expansion](#glossary-expansion)
  - [File Inclusion](#file-inclusion)
- [Version 1.0.4](#version-104)
- [Quick Start](#quick-start)
- [Building from source](#building-from-source)
- [How I use it](#how-i-use-it)
- [Examples](#examples)
- [Authors](#authors)
- [License](#license)

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
It also has a bug, for example, say if inside a code block has a comment with #, it interprets 
it as header and adds in TOC. Also, project does not seem to be active. I like
the file inclusion feature of [markdown_helper](https://github.com/BurdetteLamar/markdown_helper), `markdown-doc-go` v1.0.4 adds
this feature.

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
- **File inlcusion**: supports embedding content from other Markdown or plain text files.
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



## File Inclusion
`markdown-toc-go` now supports embedding content from other Markdown or 
plain text files directly into your main input document. This feature is 
invaluable for breaking down large documents, reusing content snippets, or 
maintaining distinct sections separately. 

**Syntax**

To include a file, use the following directive:

```
@[:markdown](path/to/your/file.md)
```

The syntax is similar to [markdown_helper](https://github.com/BurdetteLamar/markdown_helper) ruby gem. 

- The `path/to/your/file.md` is the path to the file you wish to include. 
This path is resolved relative to your main input file (specified by the `-i` flag).

- The content of the included file will be inserted directly at the position of the directive.

- Glossary expansion will be applied to the content of the included file, 
    allowing you to use within your included snippets.

**Example**
Let's say you have a main.md which includes many other files as sections:

`main.md`:

```
@[:markdown](intro.md)
@[:markdown](history.md)
@[:markdown](features.md)
@[:markdown](glossary_expansion.md)
@[:markdown](file_inclusion.md)
@[:markdown](version.md)
@[:markdown](quick_start.md)
@[:markdown](build.md)
@[:markdown](how_I_use.md)
@[:markdown](examples.md)
@[:markdown](authors.md)
@[:markdown](license.md)
```

When you run `markdown-toc-go -i main.md --glossary my_glossary.txt`, the 
content of the included files will replace the directives in `main.md` before the 
TOC is generated.

# Version 1.0.4
The current version is 1.0.4

Please look at [ChangeLog](ChangeLog.md) for what has changed in the current version.

# Quick Start

Download pre-compiled binaries from
[Releases](https://github.com/muquit/markdown-toc-go/releases) page


# Building from source
Install [go](https://go.dev/) first

```
git clone https://github.com/yourusername/markdown-toc-go.git
cd markdown-toc-go
go build
```
Look at `Makefile` for more info.
```
make build
or
make build_all
```

(Optional) Move the binary to your PATH:
```
sudo mv markdown-toc-go /usr/local/bin/
```


# How I use it

I keep my markdown files in `./docs/` directory and create a separate markdown
files for each section and include them from [docs/main.d](docs/main.md) and
and use [docs/main.d](docs/main.md) as the input file with `-i`. Example:

```bash
./markdown-toc-go -i docs/main.md -o ./README.md --glossary docs/glossary.txt -f
```
Here is how the [docs/main.d](docs/main.md) files look like:

```
@[:markdown](intro.md)
@[:markdown](history.md)
@[:markdown](features.md)
@[:markdown](glossary_expansion.md)
@[:markdown](version.md)
@[:markdown](quick_start.md)
@[:markdown](build.md)
@[:markdown](how_I_use.md)
@[:markdown](examples.md)
@[:markdown](authors.md)
@[:markdown](license.md)
```
**Note:** you do not have to use it that way, you can have one README.md as input
file.

# Examples

- Basic usage. This will generate a new file named `README-with-toc.md` 
containing the original content plus a table of contents at the top.
```
markdown-toc-go -i README.md -glossary glossary.txt
```

- Keep README.md in say `./docs` directory and generate `README.md` with TOC in 
the current working directory and overwrite README.md forcefully. Please have
a look at the `Makefile`'s `doc` target.

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
- Assemble many files with sections in a separate file
```
markdown-toc-go -i docs/main.md -o ./README.md
```
`docs/main.md` might look like:
```
@[:markdown](intro.md)
@[:markdown](history.md)
@[:markdown](features.md)
@[:markdown](glossary_expansion.md)
@[:markdown](version.md)
@[:markdown](quick_start.md)
@[:markdown](build.md)
@[:markdown](how_I_use.md)
@[:markdown](examples.md)
@[:markdown](authors.md)
@[:markdown](license.md)
```
`./docs` directory contains all the files with sections

- By default, a credit line is added at the end of the document. If you do not want to 
show the credit line, run with the option
```
-no-credit
```




etc.


# Authors

- First cut: developed with [Claude AI 3.7,4 Sonnet](https://claude.ai), working under my guidance and instructions.
- Modularized with [Google Gemini 2.5 Flash](https://gemini.google.com/app)
- File inclusion code was done by [Google Gemini 2.5 Flash](https://gemini.google.com/app) without much hand holding


# License

This project is licensed under the MIT License - see the LICENSE.txt file for details.



---
<sub>TOC is created by https://github.com/muquit/markdown-toc-go on Jul-04-2025</sub>
