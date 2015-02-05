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
	Route{"CampaignShow", "GET", "/campaigns/{id}", CampaignShow},
	Route{"CampaignCreate", "POST", "/campaigns", CampaignPost},
	Route{"CampaignDelete", "DELETE", "/campaigns/{id}", CampaignDelete},
}
