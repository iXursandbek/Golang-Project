package homecontroller

import (
	"net/http"
	"text/template"
)

func Index(response http.ResponseWriter, request *http.Request) {
	tmp, _ := template.ParseFiles(
		"views/templates/mytemplates.html",
		"views/home/index.html",
	)
	tmp.ExecuteTemplate(response, "layout", nil)
}
