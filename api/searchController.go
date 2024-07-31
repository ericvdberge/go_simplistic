package api

import (
	"net/http"
)

func HandleGetSearch(w http.ResponseWriter, r *http.Request) {
	RenderPage(w, "web/pages/search.html", nil)
}
