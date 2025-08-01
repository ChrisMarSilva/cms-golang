package handlers

import (
	"net/http"

	"github.com/chrismarsilva/cms.project.todo/views"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) error {
	return render(w, r, views.HomePage())
}
