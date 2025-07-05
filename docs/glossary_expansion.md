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


