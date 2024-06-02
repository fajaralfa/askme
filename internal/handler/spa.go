package handler

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/fajaralfa/askme/internal/helper"
)

func SpaHandler(w http.ResponseWriter, r *http.Request) {
	manifest := helper.ParseManifestJSON("./dist/.vite/manifest.json")

	t, err := template.ParseFiles("./web/index.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, map[string]any{
		"ENVIRONMENT": os.Getenv("ENVIRONMENT"),
		"CssFiles":    manifest["web/main.js"]["css"],
		"JsFile":      manifest["web/main.js"]["file"],
	})
}
