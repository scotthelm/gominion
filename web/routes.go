package web

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", Index},
	Route{"CampaignIndex", "GET", "/campaigns", CampaignIndex},
	Route{"CampaignShow", "GET", "/campaigns/{campaignId}", CampaignShow},
}