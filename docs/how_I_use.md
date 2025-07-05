# How I use it

I keep my markdown files in `./docs/` directory and create a separate markdown
files for each section and include them from @MAIN_MD@ and
and use @MAIN_MD@ as the input file with `-i`. Example:

```bash
./markdown-toc-go -i docs/main.md -o ./README.md --glossary docs/glossary.txt -f
```
Here is how the @MAIN_MD@ files look like:

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
