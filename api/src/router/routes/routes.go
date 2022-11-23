package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI         string
	Method      string
	Function    func(http.ResponseWriter, *http.Request)
	RequireAuth bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := basicRoutes
	routes = append(routes, userRoutes...)
	routes = append(routes, loginRoutes)
	routes = append(routes, publicationRoutes...)

	for _, route := range routes {
		if route.RequireAuth {
			r.HandleFunc(
				route.URI,
				middlewares.Logger(
					middlewares.Auth(route.Function),
				),
			).Methods(route.Method)
		} else {
			r.HandleFunc(
				route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	return r
}
