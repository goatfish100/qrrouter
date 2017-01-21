package main

import "net/http"

//Route - the route structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes - Create he routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetResource",
		"GET",
		"/qrresource/{resourceId}",
		GetResource,
	},
	Route{
		"ResourceCreate",
		"POST",
		"/resourcecreate",
		ResourceCreate,
	},
	Route{
		"GetOrgs",
		"GET",
		"/getorgs",
		GetOrgs,
	},
	Route{
		"GetOrg",
		"GET",
		"/getorg/{orgId}",
		GetOrg,
	},
	Route{
		"PostCreateOrg",
		"post",
		"/createorg",
		PostCreateOrg,
	},
}
