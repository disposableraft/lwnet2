package gatsby

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/gomarkdown/markdown"
)

const (
	SOURCE_CONTENT = "src/content/"
	TARGET_CONTENT = "docs/"
)

type file struct {
	name     string
	path     string
	markdown templateData
	data     []byte
	headline string
}

type content struct {
	Template         string
	ContentDirectory string
	PublicDirectory  string
	Files            []file
}

type templateData struct {
	HTML string
}

func Register(contentDirectory string, template string) *content {
	// Check that there's a trailing slash
	if endsWithSlash := strings.HasSuffix(contentDirectory, "/"); !endsWithSlash {
		contentDirectory = fmt.Sprintf("%s/", contentDirectory)
	}

	c := new(content)
	c.Template = template
	c.ContentDirectory = contentDirectory
	c.PublicDirectory = strings.Replace(contentDirectory, SOURCE_CONTENT, TARGET_CONTENT, 1)

	return c
}

// CollectFiles collects files from `gatsyb.ContentDirectory` and adds markdown
// data to `gatsby.Files`.
func (c *content) CollectFiles() error {
	entries, err := os.ReadDir(c.ContentDirectory)

	if err != nil {
		log.Fatalf("error reading directory (%s): %s", c.ContentDirectory, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		file, err := c.collectFile(entry)
		if err != nil {
			log.Fatalf("error collecting file %v", entry)
		}
		c.Files = append(c.Files, file)
	}

	return nil
}

func (c *content) collectFile(entry fs.DirEntry) (file, error) {
	name := strings.Replace(entry.Name(), ".md", ".html", 1)
	path := fmt.Sprintf("%s%s", c.PublicDirectory, name)

	mdPath := fmt.Sprintf("%s%s", c.ContentDirectory, entry.Name())

	file := file{
		name: name,
		path: path,
		data: readFile(mdPath),
	}

	file.markdown = templateData{
		// Error could come from here
		HTML: string(markdown.ToHTML(file.data, nil, nil)),
	}

	file.headline = scanFile(file.data)

	return file, nil
}

func (c *content) ExportFiles() error {
	for _, file := range c.Files {
		tmpl := template.Must(template.ParseFiles(c.Template))
		out := createFile(c.PublicDirectory, file)
		tmpl.Execute(out, file.markdown)
	}
	return nil
}

func readFile(path string) []byte {
	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading file (%s): %s", path, err)
	}
	return bytes
}

func createFile(dir string, f file) *os.File {
	if isDirNotExist(dir) {
		if err := os.MkdirAll(dir, 0754); err != nil {
			fmt.Printf("failed to make directories: %v\n", err)
		}
	}

	file, err := os.Create(f.path)

	if err != nil {
		log.Fatalf("error creating file (%s): %s", f.path, err)
	}

	return file
}

func isDirNotExist(dir string) bool {
	fileInfo, err := os.Stat(dir)
	if os.IsNotExist(err) || !fileInfo.IsDir() {
		return true
	}
	return false
}

func (c *content) GetArticles() []file {
	return c.Files
}

func createIndexPageLinks(files []file) string {
	var data string

	// Sort the slice of files desc.
	sort.Slice(files, func(i, j int) bool { return files[i].path > files[j].path })

	for _, f := range files {
		path := strings.TrimPrefix(f.path, "docs")
		data += fmt.Sprintf("<li><a href=\"%s\">%s</a></li>", path, f.headline)
	}

	return data
}

func ExportIndexPage(fileCollections ...[]file) {
	var data string

	for _, files := range fileCollections {
		data += createIndexPageLinks(files)
	}

	file, err := os.Create(fmt.Sprintf("%sindex.html", TARGET_CONTENT))

	if err != nil {
		log.Fatalf("error creating file")
	}

	indexTemplate := "templates/index.html"
	tmpl := template.Must(template.ParseFiles(indexTemplate))

	tmpl.Execute(file, templateData{HTML: data})
}
