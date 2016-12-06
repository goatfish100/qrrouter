package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

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
}
