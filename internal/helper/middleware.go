package helper

import "net/http"

func Middlewares(h http.Handler, mws ...func(http.Handler) http.Handler) http.Handler {
	if len(mws) < 1 {
		return h
	}
	wrappedhandler := h

	for i := len(mws) - 1; i >= 0; i-- {
		wrappedhandler = mws[i](wrappedhandler)
	}

	return wrappedhandler
}
