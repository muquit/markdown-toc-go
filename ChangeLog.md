## Contents
- [v1.0.3](#v103)
- [v1.0.2](#v102)
- [v1.0.1](#v101)

# v1.0.3
* Add glossary support with `--glossary` option
* Supports `@KEY@` syntax for key-value expansion in markdown
* Multi-level nested references in glossary values
* Expansion occurs before TOC generation
* See `glossary.txt`, `test/TEST_GLOSSARY.md` and `TEST_GLOSSARY_EXPANDED.md` for examples

Example:
```
# Create glossary file
echo "APP_NAME    MyApp" > glossary.txt
echo "VERSION     v1.0" >> glossary.txt
echo "REPO_URL    https://github.com/user/MyApp" >> glossary.txt

# Use in markdown: # @APP_NAME@ @VERSION@ Documentation
./markdown-toc-go -i README.md --glossary glossary.txt
```

 (Jun-10-2025)

# v1.0.2
* Add options `--pre-toc`, `--pre-toc-file` and `--insert-position``
Example:
```
# Add badges directly via command line
./markdown-toc-go -i README.md --pre-toc "[![Download](https://img.shields.io/github/downloads/muquit/mailsend/total.svg)](https://github.com/      muquit/mailsend/releases)"
# Or use a file containing your badges
./markdown-toc-go -i README.md --pre-toc-file badges.md
# Control the position (before or after the TOC title)
./markdown-toc-go -i README.md --pre-toc-file badges.md --insert-position after-title
```
# v1.0.1
* Initial Release
  (Mar-29-2025)
