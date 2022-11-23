package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.GetUsers,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodGet,
		Function:    controllers.GetUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}/follow",
		Method:      http.MethodPost,
		Function:    controllers.FollowUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}/stop-follow",
		Method:      http.MethodPost,
		Function:    controllers.UnfollowUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}/followers",
		Method:      http.MethodGet,
		Function:    controllers.GetFollowers,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}/following",
		Method:      http.MethodGet,
		Function:    controllers.GetFollowing,
		RequireAuth: true,
	},
	{
		URI:         "/users/{userId}/update-password",
		Method:      http.MethodPost,
		Function:    controllers.UpdatePassword,
		RequireAuth: true,
	},
}
