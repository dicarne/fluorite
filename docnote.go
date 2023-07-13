package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/google/uuid"
)

type FrontMatter struct {
	DateCreated  string `yaml:"date created"`
	DateModified string `yaml:"date modified"`
}

func parserHook(data []byte) (ast.Node, []byte, int) {
	if node, d, n := ParseDocLink(data); node != nil {
		return node, d, n
	}
	if node, d, n := ParseDocLinkEmb(data); node != nil {
		return node, d, n
	}
	if node, d, n := ParseCallout(data); node != nil {
		return node, d, n
	}
	return nil, nil, 0
}

func shortURL(url string) string {
	sp := strings.Split(url, "|")
	if len(sp) == 1 {
		return url
	}
	return sp[1]
}

func renderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if d, ok := node.(*DocLink); ok {
		if entering {
			file, err := findFileWithShortNamePath(d.URL)
			if err != nil {
				fmt.Printf("Cant find url: %s\n", d.URL)
				return ast.GoToNext, true
			}
			if d.Inline {
				io.WriteString(w, fmt.Sprintf("<a href=\"%s\"  target=\"_blank\">%s</a>", file.ShortWebPath, shortURL(d.URL)))
			} else {
				io.WriteString(w, fmt.Sprintf("<p><a href=\"%s\"  target=\"_blank\">%s</a></p>", file.ShortWebPath, shortURL(d.URL)))
			}

		}
		return ast.GoToNext, true
	}

	if d, ok := node.(*DocLinkEmb); ok {
		if entering {
			file, err := findFileWithShortNamePath(d.URL)
			if err != nil {
				fmt.Printf("Cant find url: %s\n", d.URL)
				return ast.GoToNext, true
			}
			if IsImage(file.Ext) {
				io.WriteString(w, fmt.Sprintf("<img src=\"%s\" alt=\"%s\"/>", file.ShortWebPath, shortURL(d.URL)))
			} else {
				io.WriteString(w, fmt.Sprintf("<a href=\"%s\"  target=\"_blank\">%s</a>", file.ShortWebPath, shortURL(d.URL)))
			}

		}
		return ast.GoToNext, true
	}

	if d, ok := node.(*Callout); ok {
		if entering {
			io.WriteString(w, fmt.Sprintf("\n<div class=\"callout-%s\">", d.Tag))
			if d.Title != "" {
				io.WriteString(w, fmt.Sprintf("<p class=\"callout-title\">%s</p>", d.Title))
			}
		} else {
			io.WriteString(w, "\n</div>")
		}
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}

func mdToHTML(md []byte, frontMatter FrontMatter) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock | parser.EmptyLinesBreakList | parser.HardLineBreak
	parserer := parser.NewWithExtensions(extensions)
	parserer.Opts.ParserHook = parserHook

	doc := parserer.Parse(md)
	doc = modifyAst(doc)

	// create HTML renderer with extensions
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{
		Flags:          htmlFlags,
		RenderNodeHook: renderHook,
		CSS:            "../main.css",
	}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func readAndParseMD(filedata FileData) {
	md, err := os.ReadFile(filedata.AbsPath)
	if err != nil {
		log.Fatal(err)
	}

	frontMatterBin := []byte("")
	frontMatter := FrontMatter{}

	// has yaml front matter
	if bytes.HasPrefix(md, []byte("---")) {
		frontMatterPrefixLen := len("---\n")
		end := bytes.Index(md[frontMatterPrefixLen:], []byte("---\n"))
		if end != -1 {
			frontMatterBin = md[frontMatterPrefixLen : end+frontMatterPrefixLen]
			md = md[end+frontMatterPrefixLen*2:]
		}
	}
	if len(frontMatterBin) > 0 {
		err = yaml.Unmarshal(frontMatterBin, &frontMatter)
		if err != nil {
			fmt.Println("error front matter")
		}
	}

	html := mdToHTML(md, frontMatter)
	buf := bytes.Buffer{}
	buf.WriteString("<h1>")
	titles := strings.Split(filedata.Path, "/")
	title := strings.Split(titles[len(titles)-1], ".")[0]
	buf.WriteString(title)
	buf.WriteString("</h1>")
	buf.WriteString("<hr/>")
	buf.Write(html)
	html = wrapHTML(buf.Bytes(), filedata.Path, "../main.css")
	os.WriteFile(filedata.WebPath, html, 0666)
}

type ObsidianAppConfig struct {
	AttachmentFolderPath string `json:"attachmentFolderPath"`
}

type FileData struct {
	Path         string
	UUID         string
	AbsPath      string
	Ext          string
	WebPath      string
	ShortWebPath string
}

var files []FileData

func findFileWithShortNamePath(namePath string) (FileData, error) {
	removeOther := strings.Split(namePath, "#")[0]
	removeOther = strings.Split(removeOther, "|")[0]
	namePath = strings.Split(removeOther, "^")[0]
	dots := strings.Split(namePath, ".")
	if len(dots) == 1 {
		// no dot means is markdown
		namePath += ".md"
	}
	for _, v := range files {
		if strings.HasSuffix(v.Path, namePath) {
			return v, nil
		}
	}
	return FileData{}, fmt.Errorf("can't find this file: %s", namePath)
}

var obsidianConfigFolder = ".obsidian"

func generateObsidianValt(obsidianRoot string, outputFolder string, themeName string) {
	config := ObsidianAppConfig{}
	appconfBin, err := os.ReadFile(path.Join(obsidianRoot, obsidianConfigFolder, "app.json"))
	IfFatal(err, "Obsidian config load error!")

	err = json.Unmarshal(appconfBin, &config)
	IfFatal(err, "Obsidian app.json load error!")

	filePaths, err := GetAllFiles(obsidianRoot)
	IfFatal(err, "Unable to get files.")

	ClearFolder(outputFolder)
	os.Mkdir(path.Join(outputFolder, "notes"), 0666)
	CopyFile("theme/"+themeName+"/main.css", path.Join(outputFolder, "main.css"))

	prefixLen := len(obsidianRoot)
	for _, filePath := range filePaths {
		filePath = filepath.ToSlash(filePath) // change "\"" to "/"
		uuid := uuid.New()
		dots := strings.Split(filePath, ".")
		if len(dots) == 1 {
			dots = append(dots, "unknow")
		}
		ext := dots[len(dots)-1]
		ext2 := ext
		if ext2 == "md" {
			ext2 = "html"
		}
		files = append(files, FileData{
			Path:         filePath[prefixLen+1:],
			UUID:         uuid.String(),
			AbsPath:      filePath,
			Ext:          ext,
			WebPath:      path.Join(outputFolder, "notes", fmt.Sprintf("%x.%s", md5.Sum([]byte(filePath[prefixLen+1:])), ext2)),
			ShortWebPath: fmt.Sprintf("%x.%s", md5.Sum([]byte(filePath[prefixLen+1:])), ext2),
		})
	}

	for _, file := range files {
		if file.Ext != "md" {
			// is resource file
			CopyFile(file.AbsPath, file.WebPath)
			continue
		}
		readAndParseMD(file)
	}
	buildIndex(outputFolder)
}

func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := os.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)

	for _, fi := range dir {
		if fi.IsDir() {
			if fi.Name() == obsidianConfigFolder {
				continue
			}
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}

	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		files = append(files, temp...)
	}

	return files, nil
}

func buildIndex(outputFolder string) {
	buf := bytes.Buffer{}
	buf.WriteString("<h1>Index</h1>")
	buf.WriteString("<div>")
	for _, v := range files {
		if v.Ext != "md" {
			continue
		}
		buf.WriteString("<p><a href=\"")
		buf.WriteString("notes/" + v.ShortWebPath)
		buf.WriteString("\">")
		buf.WriteString(v.Path)
		buf.WriteString("</a></p>\n")
	}
	buf.WriteString("</div>")
	html := wrapHTML(buf.Bytes(), "Home", "main.css")
	os.WriteFile(path.Join(outputFolder, "index.html"), html, 0666)
}
