package main

import "lancewakeling.net/website/src/gatsby"

func main() {
	// TODO Use command line flags/args to export individual directories
	// https://gobyexample.com/command-line-arguments
	blog := gatsby.Register("src/content/writing/blog/", "templates/page.html")
	blog.CollectFiles()
	blog.ExportFiles()
	blogArticles := blog.GetArticles()

	pages := gatsby.Register("src/content/pages/", "templates/page.html")
	pages.CollectFiles()
	pages.ExportFiles()
	pageArticles := pages.GetArticles()

	// Create an index page
	gatsby.ExportIndexPage(blogArticles, pageArticles)

	// TODO Use command line flags to create a new git branch and publish it to GH
}
