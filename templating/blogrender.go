package blogrender

import (
	"embed"
	"html/template"
	"io"

	"github.com/gomarkdown/markdown"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

// const (
// 	postTemplate = "<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags:<ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>"
// )

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p *Post) error {
	if err := r.templ.Execute(w, &struct {
		Title, Body, Description string
		Tags                     []string
	}{
		Title:       p.Title,
		Body:        string(markdown.ToHTML([]byte(p.Body), nil, nil)), // we could remove scaped html
		Description: p.Description,
		Tags:        p.Tags,
	}); err != nil {
		return err
	}
	return nil
}
