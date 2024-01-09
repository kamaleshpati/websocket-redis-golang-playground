package route

import (
	"net/http"

	"github.com/bmizerany/pat"

	"github.com/kamaleshpati/wsredisPlayground/internal/route/v0/handler"
)

func Routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handler.WSEndpointHandler))

	return mux
}
