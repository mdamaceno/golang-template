package view

import (
    "path/filepath"
    "html/template"
    "fmt"

    "github.com/gin-gonic/gin/render"
)

type Render map[string]*template.Template

type Renderer interface {
	render.HTMLRender
	add(name string, tmpl *template.Template)
	addFromFiles(name string, files ...string) *template.Template
}

var (
	_ render.HTMLRender = Render{}
	_ Renderer          = Render{}
    viewsDir            = filepath.Join("web", "view")
    defaultLayoutFile   = filepath.Join(viewsDir, "layouts", "application.html")
)

func SetViews() Renderer {
    r := make(Render)

    r.addFromFiles("home/index", defaultLayoutFile, viewsDir + "/home/index.html")
    r.addFromFiles("users/index", defaultLayoutFile, viewsDir + "/users/index.html")
    r.addFromFiles("users/show", defaultLayoutFile, viewsDir + "/users/show.html")
    r.addFromFiles("users/new", defaultLayoutFile, viewsDir + "/users/new.html")
    r.addFromFiles("users/edit", defaultLayoutFile, viewsDir + "/users/edit.html")

    return r
}

func (r Render) add(name string, tmpl *template.Template) {
	if tmpl == nil {
		panic("template can not be nil")
	}

	if len(name) == 0 {
		panic("template name cannot be empty")
	}

	if _, ok := r[name]; ok {
		panic(fmt.Sprintf("template %s already exists", name))
	}

	r[name] = tmpl
}

func (r Render) addFromFiles(name string, files ...string) *template.Template {
	tmpl := template.Must(template.ParseFiles(files...))
	r.add(name, tmpl)

	return tmpl
}

func (r Render) Instance(name string, data interface{}) render.Render {
	return render.HTML{
		Template: r[name],
		Data:     data,
	}
}
