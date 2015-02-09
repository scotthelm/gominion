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
	Route{"CampaignIndex", "GET", "/api/campaigns", CampaignIndex},
	Route{"CampaignShow", "GET", "/api/campaigns/{id}", CampaignShow},
	Route{"CampaignCreate", "POST", "/api/campaigns", CampaignPost},
	Route{"CampaignDelete", "DELETE", "/api/campaigns/{id}", CampaignDelete},
	Route{"CampaignUpdate", "PUT", "/api/campaigns/{id}", CampaignUpdate},
}
