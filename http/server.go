package http

import (
	"github.com/gin-gonic/gin"
)

type Route struct {
	Path    string
	Method  string
	Handler func(*gin.Context)
}

var routes = []Route{
	{
		Path:    "/.well-known/healthcheck",
		Method:  "get",
		Handler: Healthcheck,
	},
	{
		Path:    "/users",
		Method:  "post",
		Handler: CreateUser,
	},
	{
		Path:    "/users/:id",
		Method:  "get",
		Handler: GetUser,
	},
	{
		Path:    "/users/:id/communities",
		Method:  "get",
		Handler: GetUserCommunities,
	},
	{
		Path:    "/communities",
		Method:  "post",
		Handler: CreateCommunity,
	},
	{
		Path:    "/communities",
		Method:  "get",
		Handler: GetCommunities,
	},
	{
		Path:    "/communities/:id",
		Method:  "get",
		Handler: GetCommunity,
	},
	{
		Path:    "/communities/:id/users",
		Method:  "post",
		Handler: AddUserToCommunity,
	},
	{
		Path:    "/communities/:id/users",
		Method:  "get",
		Handler: GetUsersInCommunity,
	},
	{
		Path:    "/communities/:id/users/:user_id",
		Method:  "delete",
		Handler: RemoveUserFromCommunity,
	},
}

func New() *gin.Engine {
	router := gin.Default()

	for _, route := range routes {
		switch route.Method {
		case "get":
			router.GET(route.Path, route.Handler)
		case "post":
			router.POST(route.Path, route.Handler)
		case "put":
			router.PUT(route.Path, route.Handler)
		case "patch":
			router.PATCH(route.Path, route.Handler)
		case "delete":
			router.DELETE(route.Path, route.Handler)
		case "head":
			router.HEAD(route.Path, route.Handler)
		default:
			router.GET(route.Path, route.Handler)
		}
	}

	return router
}
