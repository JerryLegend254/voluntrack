package main

import (
	"net/http"

	"github.com/JerryLegend254/voluntrack/ui/html"
)

type Stat struct {
	Description string
	Value       string
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	stats := []Stat{
		Stat{"Incredible  volunteers", "15K"},
		Stat{"Successful campaigns", "100+"},
		Stat{"Monthly donors", "600+"},
	}

	app.render(w, r, "home.page.tmpl", &templateData{Stats: stats})
}

func (app *application) base(w http.ResponseWriter, r *http.Request) error {
	return html.BaseLayout("Home").Render(r.Context(), w)
}
