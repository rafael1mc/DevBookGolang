package routes

import (
	"api/src/controllers"
	"net/http"
)

var basicRoutes = []Route{
	{
		URI:         "/",
		Method:      http.MethodGet,
		Function:    controllers.Root,
		RequireAuth: false,
	},
}
