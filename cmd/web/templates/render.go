package main

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
)

// Custom data type
type templateData struct {
	StringMap       map[string]string
	IntMap          map[string]int
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Flash           string
	Warning         string
	Error           string
	IsAuthenticated int
	API             string
	CSSVersion      string
}

// Functions to pass to the template
var functions = template.FuncMap{}

// Go embedding: allows to compile the application and it's associated templates into a single binary

// go:embed templates
var templateFS embed.FS

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	return td
}

// Render function
func (app *application) renderTemplates(w http.ResponseWriter, r *http.Request, page string, td *templateData, partials ...string) error {

	var t *template.Template
	var err error
	// Conditional template rendering
	templateToRender := fmt.Sprintf("templates/%s.page.tmpl", page)

	// Checking if the variable exists in the template cache
	_, templateInMap := app.templateCache[templateToRender]

	if app.config.env == "production" && templateInMap {

	}
}
