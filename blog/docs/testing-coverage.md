# Testing Coverage

Testing criteria for a passing coverage requirement:

- Line coverage of 70%
- Cognitive complexity of 0
- Have cognitive complexity < 5, but have any coverage

Low cognitive complexity means there are few conditional branches to cover. Tests with cognitive complexity 0 would be covered by invocation.

## Packages

| Status | Package                                      | Coverage | Cognitive | Lines |
|--------|----------------------------------------------|----------|-----------|-------|
| ✅     | titpetric/platform-example/blog              | 0.00%    | 0         | 0     |
| ✅     | titpetric/platform-example/blog/cmd/blog     | 0.00%    | 0         | 0     |
| ✅     | titpetric/platform-example/blog/cmd/generate | 0.00%    | 0         | 0     |
| ❌     | titpetric/platform-example/blog/markdown     | 72.00%   | 10        | 141   |
| ✅     | titpetric/platform-example/blog/model        | 0.00%    | 0         | 0     |
| ✅     | titpetric/platform-example/blog/schema       | 0.00%    | 0         | 0     |
| ✅     | titpetric/platform-example/blog/storage      | 89.92%   | 8         | 113   |
| ✅     | titpetric/platform-example/blog/template     | 0.00%    | 0         | 0     |

## Functions

| Status | Package                                      | Function                      | Coverage | Cognitive |
|--------|----------------------------------------------|-------------------------------|----------|-----------|
| ❌     | titpetric/platform-example/blog              | ContentNegotiation            | 0.00%    | 4         |
| ❌     | titpetric/platform-example/blog              | Generator.Generate            | 0.00%    | 10        |
| ❌     | titpetric/platform-example/blog              | Generator.copyAssets          | 0.00%    | 9         |
| ❌     | titpetric/platform-example/blog              | Generator.generateArticlePage | 0.00%    | 2         |
| ❌     | titpetric/platform-example/blog              | Generator.generateFeed        | 0.00%    | 2         |
| ❌     | titpetric/platform-example/blog              | Generator.generateIndexPage   | 0.00%    | 2         |
| ✅     | titpetric/platform-example/blog              | Generator.generateStaticPages | 0.00%    | 0         |
| ❌     | titpetric/platform-example/blog              | Generator.walkPages           | 0.00%    | 30        |
| ❌     | titpetric/platform-example/blog              | Handlers.GetArticleHTML       | 0.00%    | 2         |
| ❌     | titpetric/platform-example/blog              | Handlers.GetArticleJSON       | 0.00%    | 2         |
| ❌     | titpetric/platform-example/blog              | Handlers.GetAtomFeed          | 0.00%    | 2         |
| ❌     | titpetric/platform-example/blog              | Handlers.IndexHTML            | 0.00%    | 2         |
| ❌     | titpetric/platform-example/blog              | Handlers.ListArticlesHTML     | 0.00%    | 2         |
| ❌     | titpetric/platform-example/blog              | Handlers.ListArticlesJSON     | 0.00%    | 2         |
| ❌     | titpetric/platform-example/blog              | Handlers.SearchArticlesJSON   | 0.00%    | 3         |
| ❌     | titpetric/platform-example/blog              | Module.Mount                  | 0.00%    | 1         |
| ✅     | titpetric/platform-example/blog              | Module.Name                   | 0.00%    | 0         |
| ✅     | titpetric/platform-example/blog              | Module.ScanMarkdownFiles      | 0.00%    | 0         |
| ✅     | titpetric/platform-example/blog              | Module.SetRepository          | 0.00%    | 0         |
| ❌     | titpetric/platform-example/blog              | Module.Start                  | 0.00%    | 4         |
| ✅     | titpetric/platform-example/blog              | Module.Stop                   | 0.00%    | 0         |
| ❌     | titpetric/platform-example/blog              | Module.parseMarkdownFile      | 0.00%    | 13        |
| ❌     | titpetric/platform-example/blog              | Module.scanMarkdownFiles      | 0.00%    | 9         |
| ✅     | titpetric/platform-example/blog              | NewGenerator                  | 0.00%    | 0         |
| ❌     | titpetric/platform-example/blog              | NewHandlers                   | 0.00%    | 1         |
| ✅     | titpetric/platform-example/blog              | NewModule                     | 0.00%    | 0         |
| ✅     | titpetric/platform-example/blog              | generateID                    | 0.00%    | 0         |
| ❌     | titpetric/platform-example/blog/cmd/blog     | main                          | 0.00%    | 1         |
| ❌     | titpetric/platform-example/blog/cmd/blog     | start                         | 0.00%    | 1         |
| ❌     | titpetric/platform-example/blog/cmd/generate | generate                      | 0.00%    | 4         |
| ❌     | titpetric/platform-example/blog/cmd/generate | main                          | 0.00%    | 2         |
| ✅     | titpetric/platform-example/blog/markdown     | NewRenderer                   | 100.00%  | 0         |
| ✅     | titpetric/platform-example/blog/markdown     | Renderer.Render               | 100.00%  | 0         |
| ✅     | titpetric/platform-example/blog/markdown     | Renderer.highlightCodeBlocks  | 100.00%  | 0         |
| ✅     | titpetric/platform-example/blog/markdown     | Renderer.processCodeBlock     | 87.50%   | 1         |
| ✅     | titpetric/platform-example/blog/markdown     | escapeHTML                    | 0.00%    | 0         |
| ✅     | titpetric/platform-example/blog/markdown     | highlightCode                 | 88.50%   | 8         |
| ✅     | titpetric/platform-example/blog/markdown     | unescapeHTML                  | 100.00%  | 0         |
| ❌     | titpetric/platform-example/blog/markdown     | wrapCodePlain                 | 0.00%    | 1         |
| ✅     | titpetric/platform-example/blog/storage      | CountArticles                 | 100.00%  | 0         |
| ✅     | titpetric/platform-example/blog/storage      | DB                            | 0.00%    | 0         |
| ✅     | titpetric/platform-example/blog/storage      | GetArticleBySlug              | 100.00%  | 1         |
| ✅     | titpetric/platform-example/blog/storage      | GetArticles                   | 83.30%   | 1         |
| ✅     | titpetric/platform-example/blog/storage      | InsertArticle                 | 100.00%  | 0         |
| ✅     | titpetric/platform-example/blog/storage      | NewStorage                    | 100.00%  | 0         |
| ✅     | titpetric/platform-example/blog/storage      | SearchArticles                | 85.70%   | 1         |
| ✅     | titpetric/platform-example/blog/storage      | Storage.CountArticles         | 100.00%  | 1         |
| ✅     | titpetric/platform-example/blog/storage      | Storage.GetArticleBySlug      | 100.00%  | 1         |
| ✅     | titpetric/platform-example/blog/storage      | Storage.GetArticles           | 100.00%  | 1         |
| ✅     | titpetric/platform-example/blog/storage      | Storage.InitSchema            | 100.00%  | 0         |
| ✅     | titpetric/platform-example/blog/storage      | Storage.InsertArticle         | 100.00%  | 1         |
| ✅     | titpetric/platform-example/blog/storage      | Storage.SearchArticles        | 100.00%  | 1         |
| ✅     | titpetric/platform-example/blog/template     | IndexData.Map                 | 0.00%    | 0         |
| ❌     | titpetric/platform-example/blog/template     | NewViews                      | 0.00%    | 1         |
| ✅     | titpetric/platform-example/blog/template     | PostData.Map                  | 0.00%    | 0         |
| ❌     | titpetric/platform-example/blog/template     | StripFrontMatter              | 0.00%    | 2         |
| ❌     | titpetric/platform-example/blog/template     | Views.AtomFeed                | 0.00%    | 11        |
| ❌     | titpetric/platform-example/blog/template     | Views.Blog                    | 0.00%    | 3         |
| ❌     | titpetric/platform-example/blog/template     | Views.Index                   | 0.00%    | 3         |
| ✅     | titpetric/platform-example/blog/template     | Views.IndexFromArticles       | 0.00%    | 0         |
| ❌     | titpetric/platform-example/blog/template     | Views.LoadTemplate            | 0.00%    | 1         |
| ❌     | titpetric/platform-example/blog/template     | Views.Post                    | 0.00%    | 3         |
| ✅     | titpetric/platform-example/blog/template     | Views.PostFromArticle         | 0.00%    | 0         |
| ❌     | titpetric/platform-example/blog/template     | Views.RenderLayout            | 0.00%    | 2         |
| ❌     | titpetric/platform-example/blog/template     | Views.RenderPage              | 0.00%    | 3         |
| ✅     | titpetric/platform-example/blog/template     | escapeXML                     | 0.00%    | 0         |
| ❌     | titpetric/platform-example/blog/template     | fillTemplateData              | 0.00%    | 3         |
| ❌     | titpetric/platform-example/blog/template     | loadFile                      | 0.00%    | 2         |
