package routes

import (
	"api/src/controllers"
	"net/http"
)

var publicationRoutes = []Route{
	{
		URI:         "/publications",
		Method:      http.MethodPost,
		Function:    controllers.CreatePublication,
		RequireAuth: true,
	},
	{
		URI:         "/publications",
		Method:      http.MethodGet,
		Function:    controllers.GetPublications,
		RequireAuth: true,
	},
	{
		URI:         "/publications/{publicationId}",
		Method:      http.MethodGet,
		Function:    controllers.GetPublication,
		RequireAuth: true,
	},
	{
		URI:         "/publications/{publicationId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdatePublication,
		RequireAuth: true,
	},
	{
		URI:         "/publications/{userId}/publications",
		Method:      http.MethodGet,
		Function:    controllers.GetPublicationByUser,
		RequireAuth: true,
	},
	{
		URI:         "/publications/{publicationId}/like",
		Method:      http.MethodPost,
		Function:    controllers.LikePublication,
		RequireAuth: true,
	},
	{
		URI:         "/publications/{publicationId}/unlike",
		Method:      http.MethodPost,
		Function:    controllers.UnlikePublication,
		RequireAuth: true,
	},
}
