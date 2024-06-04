package handler

import "net/http"

func ApiNotFound(w http.ResponseWriter, r *http.Request) {
	ApiNotFoundErr(w, "route not found")
}
