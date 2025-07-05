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

The syntax is similar to @MARKDOWN_HELPER@ ruby gem. 

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
