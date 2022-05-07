package gatsby

func BuildSite() {
	blog := Register("src/content/writing/blog/", "templates/page.html")
	blog.CollectFiles()
	blog.ExportFiles()
	blogArticles := blog.GetArticles()

	pages := Register("src/content/pages/", "templates/page.html")
	pages.CollectFiles()
	pages.ExportFiles()
	pageArticles := pages.GetArticles()

	// Create an index page
	ExportIndexPage(blogArticles, pageArticles)
}
