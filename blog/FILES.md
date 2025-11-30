# Complete File Listing

## Module Core Files

### blog.go (280 lines)
- Module struct implementing platform.Module
- NewModule() constructor
- Name() method
- Mount(ctx, router) - Route registration
- Start(ctx, platform) - Initialization with markdown scanning
- Stop(ctx) - Cleanup
- scanMarkdownFiles() - Walk data directory
- parseMarkdownFile() - YAML + markdown parsing
- generateID() - Create unique article IDs
- init() - Auto-register with platform

### handlers.go (200 lines)
- Handlers struct with storage + helpers
- NewHandlers() constructor
- ListArticlesJSON - GET /api/blog/articles
- GetArticleJSON - GET /api/blog/articles/{slug}
- SearchArticlesJSON - GET /api/blog/search
- ListArticlesHTML - GET /blog/
- GetArticleHTML - GET /blog/{slug}
- Helper functions: ContentNegotiation, LogArticleView

## Storage Layer

### storage/db.go (14 lines)
- DB(ctx) - Get database connection from platform
- Uses platform.Database.Connect(ctx, "blog")
- Named connection: PLATFORM_DB_BLOG

### storage/storage.go (70 lines)
- Storage struct wrapping sqlx.DB
- NewStorage(db) constructor
- GetArticleBySlug(ctx, slug)
- GetAllArticles(ctx)
- SearchArticles(ctx, query)
- InsertArticle(ctx, article)
- CountArticles(ctx)
- InitSchema(ctx) - Create SQLite schema

### storage/articles.go (100 lines)
- Package-level functions for SQL operations
- GetArticleBySlug() - SELECT by slug
- GetAllArticles() - ORDER BY date DESC
- SearchArticles() - LIKE queries on title/description
- InsertArticle() - INSERT OR REPLACE
- CountArticles() - SELECT COUNT

## Template Layer

### template/base.go (50 lines)
- BaseData struct with layout fields
- Base(ctx, data) - Render base.vuego
- BaseFromArticle() - Create BaseData from Article

### template/post.go (55 lines)
- PostData struct with article-specific fields
- Post(ctx, data) - Render post.vuego
- PostFromArticle() - Create PostData from Article

### template/helpers.go (75 lines)
- TemplateHelpers struct with formatting functions
- DefaultHelpers() - Create default helpers
- MetaTitle() - Format page title
- MetaDescription() - Default description
- PostDate() - Format date (Jan 2, 2006)
- ReadingTime() - Calculate from word count
- GetCSS(), GetBundle(), GetJs() - Asset loaders

### template/component.go (40 lines)
- Component interface with Render(ctx, w)
- BaseComponent implementation
- ArticleListComponent implementation
- NewArticleListComponent() constructor

## Models

### model/article.go (35 lines)
- Article struct (12 fields)
  - ID, Slug, Title, Description, Content
  - Date, OGImage, Layout, Source
  - URL, CreatedAt, UpdatedAt
- Metadata struct (6 fields)
  - YAML front matter parsing
- ArticleList struct (4 fields)
  - Pagination metadata

## VueGo Templates

### theme/layouts/base.vuego (80 lines)
- Root HTML layout
- Meta tags for SEO
- CSS/JS bundle insertion
- Content slot with :html binding

### theme/layouts/post.vuego (60 lines)
- Article detail layout
- Title and metadata display
- Date and reading time
- Back link

### theme/components/
- **site-header.vuego** (50 lines)
  - Navigation, theme switcher integration
- **site-footer.vuego** (40 lines)
  - Copyright, social links
- **article-list.vuego** (40 lines)
  - List of articles with v-for
- **theme-machine.vuego** (150 lines)
  - Appearance/theme switcher with JS
- **page-timer.vuego** (170 lines)
  - Scroll-driven reading progress
- **lite-youtube.vuego** (170 lines)
  - Lazy-loaded YouTube embeds
- **inline-svg.vuego** (5 lines)
  - SVG loader component
- **target-toggler.vuego** (45 lines)
  - Toggle visibility of elements
- **info-cta.vuego** (15 lines)
  - Info callout box
- Plus pages: blog.vuego, index.vuego, 404.vuego

## Database Schema

### schema/001_articles.sql (50 lines)
- CREATE TABLE articles (11 columns)
- Indexes: date, slug, layout, created_at
- CREATE VIRTUAL TABLE articles_fts
- FTS5 triggers for auto-sync

## Data Files

### data/ (38 markdown files)
Sample articles with YAML front matter:
- 50-50-overflow.md
- blog-questions-challenge-2025.md
- click-spark.md
- ... (35 more)
- we-can-has-it-all.md
- x-scrolling-centered-max-width-container.md

Each includes:
- YAML front matter (title, date, description, etc.)
- Markdown content

## Configuration

### config/ (JSON files)
- navigation.json - Menu items, themes, appearances
- meta.js - Site metadata
- articles.json - External articles
- themes.json - Color themes
- Other config files

## Documentation

### ARCHITECTURE.md (450 lines)
- Complete system architecture
- Component descriptions
- Data flow diagrams (ASCII)
- Design patterns
- Database schema details
- API contracts
- Performance characteristics
- Security considerations
- Testing strategy
- Future extensions

### IMPLEMENTATION.md (300 lines)
- Architecture overview
- Database access pattern
- Module lifecycle
- Handlers description
- Storage pattern
- Templates section
- Markdown processing
- Schema definition
- Configuration
- Usage examples
- API endpoint documentation
- Design decisions

### SETUP.md (250 lines)
- Prerequisites and installation
- Configuration via environment
- Creating articles (markdown format)
- Module lifecycle explanation
- Database initialization
- Testing examples
- Troubleshooting guide
- Performance tuning
- Next steps

### PORTING.md (200 lines)
- WebC to VueGo migration guide
- Syntax changes table
- Example conversions
- Go module implementation details
- Model types
- Registration pattern
- API endpoints
- Template variables
- Migration checklist

### FILES.md (This file)
- Complete file listing with descriptions
- Line counts
- Contents of each file

### README.md
- Original blog README

## Summary

Total: 70+ files

### Code Files (17 files)
- 3 core module files (blog.go, handlers.go)
- 3 storage files
- 4 template files
- 1 model file
- 4 schema files
- Total: ~900 lines of Go code

### Template Files (15 files)
- 2 layout templates
- 9 component templates
- 3 page templates
- Plus assets (CSS, JS, images)
- Total: ~750 lines of VueGo

### Data Files (38 markdown files)
- Real blog content
- YAML front matter
- Markdown formatted

### Documentation (5 files)
- Complete implementation guide
- Setup instructions
- Architecture overview
- Migration guide
- This file listing

## Dependencies

### Go Packages
- github.com/jmoiron/sqlx - Database access
- github.com/mattn/go-sqlite3 - SQLite driver
- github.com/go-chi/chi/v5 - Router (via platform)
- github.com/russross/blackfriday/v2 - Markdown rendering
- github.com/titpetric/vuego - Template engine
- github.com/titpetric/platform - Platform framework
- gopkg.in/yaml.v3 - YAML parsing

### External Tools
- SQLite 3.x
- Go 1.20+

## Getting Started

1. Read SETUP.md for configuration
2. Read IMPLEMENTATION.md for architecture
3. Check ARCHITECTURE.md for design details
4. Look at blog.go for module implementation
5. Review handlers.go for HTTP handling
6. Check template/ for rendering logic
7. Browse theme/ for UI components
