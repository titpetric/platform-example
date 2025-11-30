# Blog

This is a blog application. It uses the `user` module to provide session functionality,
and defines multiple roles with access to a blog. The roles are:

- Owner: can manage articles and users
- Editor: can only manage articles, can't manage users
- Reviewer: can suggest edits pending approval by owner/editor

Owners can do everything, including the review flow, suggesting changes
/ editing, without necessarily go into the editor. Approval can be made
separately. Think of it as GitHub comments with code change suggestions.

## Theme

The blog theme is a variant of `ryan-mulligan-dev`. The original templates
were ported to [vuego](https://github.com/titpetric/vuego). The template
structure is comprised of two folders:

- `theme` - the page layouts (index, post)
- `theme/components` - components and data files for inclusion
- `theme/assets` - static file assets (css, js)

## Storage

The blog system scans a source folder for markdown files. They are read,
their metadata is decoded to decide on which page layout to render, and
as needed the markdown is rendered to html.

After indexing the files, the application provides search functionality.
The API will return HTML if `text/html` is requested. It will return JSON
if requested with the appropriate Content-Type header.

The storage in use is a sqlite `:memory:` db. All the markdown files on
disk contain the source of truth, so the application only builds this as
a working index to ease searching and listing articles.
