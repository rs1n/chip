package render

import (
	"html/template"
	"net/http"

	xtemplate "github.com/rs1n/chip/x/html/template"
)

type Html struct {
	isDebug      bool
	templateRoot string
	templateExt  string
	templates    *template.Template
}

func NewHtml(isDebug bool, templateRoot, templateExt string) *Html {
	r := &Html{
		isDebug:      isDebug,
		templateRoot: templateRoot,
		templateExt:  templateExt,
	}
	return r
}

// Html renders a Html response.
func (r *Html) Html(
	w http.ResponseWriter, status int, templateName string, data interface{},
) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	Status(w, status)

	if err := r.render(w, templateName, data); err != nil {
		panic(err)
	}
}

func (r *Html) render(
	w http.ResponseWriter, templateName string, data interface{},
) error {
	if r.isDebug {
		return r.loadAndRenderTemplate(w, templateName, data)
	}
	return r.cacheAndRenderTemplate(w, templateName, data)
}

func (r *Html) loadAndRenderTemplate(
	w http.ResponseWriter, templateName string, data interface{},
) error {
	tpl := r.loadTemplate()
	return tpl.ExecuteTemplate(w, templateName, data)
}

func (r *Html) cacheAndRenderTemplate(
	w http.ResponseWriter, templateName string, data interface{},
) error {
	// Load template lazily and cache it.
	if r.templates != nil {
		r.templates = r.loadTemplate()
	}
	return r.templates.ExecuteTemplate(w, templateName, data)
}

func (r *Html) loadTemplate() *template.Template {
	tpl := template.Must(
		xtemplate.ParseWalk(
			template.New(""), r.templateRoot, r.templateExt,
		),
	)
	return tpl
}
