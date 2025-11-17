package main

import (
	"fmt"
	"net/http"
	"text/template"
	"time"
)

const pathToTempletes = "./cmd/web/templates"

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int
	FloatMap      map[string]float64
	Data          map[string]interface{}
	Flash         string
	Warning       string
	Error         string
	Authenticated bool
	Now           time.Time
	// User
}

func (app *Config) render(w http.ResponseWriter, r *http.Request, temp string, tempData *TemplateData) {
	partials := []string{
		fmt.Sprintf("%s/base.layout.gohtml", pathToTempletes),
		fmt.Sprintf("%s/header.partial.gohtml", pathToTempletes),
		fmt.Sprintf("%s/navbar.partial.gohtml", pathToTempletes),
		fmt.Sprintf("%s/footer.partial.gohtml", pathToTempletes),
		fmt.Sprintf("%s/alerts.partial.gohtml", pathToTempletes),
	}
	var templeteSlice []string
	templeteSlice = append(templeteSlice, fmt.Sprintf("%s/%s", pathToTempletes, temp))

	templeteSlice = append(templeteSlice, partials...)

	if tempData == nil {
		tempData = &TemplateData{}
	}

	tmpl, err := template.ParseFiles(templeteSlice...)
	if err != nil {
		app.Errorlog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, app.AddDefaultData(r, tempData)); err != nil {
		app.Errorlog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (app *Config) AddDefaultData(r *http.Request, tempData *TemplateData) *TemplateData {
	tempData.Flash = app.Session.PopString(r.Context(), "flash")
	tempData.Warning = app.Session.PopString(r.Context(), "warning")
	tempData.Error = app.Session.PopString(r.Context(), "error")
	tempData.Authenticated = app.isAuthenticated(r)
	tempData.Now = time.Now()
	return tempData
}

func (app *Config) isAuthenticated(r *http.Request) bool {
	return app.Session.Exists(r.Context(), "userID")
}
