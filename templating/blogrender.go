package blogrender

import (
	"embed"
	"html/template"
	"io"
)

var (
	//go:embed "templates/*"
	posTemplates embed.FS
)

// const (
// 	postTemplate = "<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags:<ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>"
// )

type Post struct {
	Title, Body, Description string
	Tags                     []string
}

func Render(w io.Writer, p *Post) error {
	// templ, err := template.New("blog").Parse(postTemplate)
	templ, err := template.ParseFS(posTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil {
		return err
	}

	return nil
	// _, err := fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>Tags:<ul>", p.Title, p.Description)
	// if err != nil {
	// 	return err
	// }
	// for _, t := range p.Tags {
	// 	_, err = fmt.Fprintf(w, "<li>%s</li>", t)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	// _, err = fmt.Fprint(w, "</ul>")
	// return err
}
