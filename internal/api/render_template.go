package api

import (
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func NewTemplate() *Template {
	return &Template{
		templates: template.Must(template.ParseGlob("./web/templates/*.tmpl")),
	}
}

func (t *Template) Render(w io.Writer, name string, data any) error {
	return t.RenderFragment(w, name, name, data)
}

func (t *Template) RenderFragment(w io.Writer, fileName string, name string, data any) error {
	tmpl := template.Must(t.templates.Clone())
	tmpl = template.Must(tmpl.ParseGlob("./web/templates/" + fileName + ".tmpl"))

	return tmpl.ExecuteTemplate(w, name, data)
}
