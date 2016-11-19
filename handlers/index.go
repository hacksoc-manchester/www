package handlers

import (
	"github.com/hacksoc-manchester/www/services/eventService"
	"net/http"
)

type indexModel struct {
	HasUpcomingEvent bool
	Event            *eventService.Event
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	var model indexModel

	model.HasUpcomingEvent = model.Event != nil
	templates["index"].ExecuteTemplate(w, "layout", &model)
}
