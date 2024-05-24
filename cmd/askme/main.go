package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/fajaralfa/askme/internal/utils"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", homePage)
	router.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("dist"))))

	http.ListenAndServe(":8080", router)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	manifest, err := utils.ParseManifestJSON("./dist/.vite/manifest.json")
	if err != nil {
		log.Fatal(err)
	}

	t, err := template.ParseFiles("./web/index.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, map[string]any{
		"Development": true,
		"CssFiles":    manifest["web/main.js"]["css"],
		"JsFile":      manifest["web/main.js"]["file"],
	})
}
