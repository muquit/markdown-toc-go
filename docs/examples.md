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

