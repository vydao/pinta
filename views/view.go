package views

import (
	"html/template"
	"net/http"
)

type View struct {
	Template *template.Template
	Layout   string
}

var (
	LayoutDir   string = "views/layouts/"
	TemplateExt string = ".gohtml"
)

func NewView(files ...string) *View {
	// files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}

// func layoutFiles() []string {
// 	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return files
// }

func (v *View) Render(w http.ResponseWriter,
	data interface{}) error {
	return v.Template.Execute(w, data)
}
