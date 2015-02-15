package web

import (
	m "github.com/scotthelm/gominion/models"
	"net/http"
	// "reflect"
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
	Route{"CampaignIndex", "GET", "/api/campaigns", IndexHandler([]m.Campaign{})},
	Route{"CampaignShow", "GET", "/api/campaigns/{id}", ShowHandler(m.Campaign{})},
	Route{"CampaignCreate", "POST", "/api/campaigns", CreateHandler(m.Campaign{})},
	Route{"CampaignDelete", "DELETE", "/api/campaigns/{id}", DeleteHandler(m.Campaign{})},
	Route{"CampaignUpdate", "PUT", "/api/campaigns/{id}", UpdateHandler(m.Campaign{})},
	Route{"RaceIndex", "GET", "/api/races", IndexHandler([]m.Race{})},
	Route{"RaceShow", "GET", "/api/races/{id}", ShowHandler(m.Race{})},
	Route{"RaceCreate", "POST", "/api/races", CreateHandler(m.Race{})},
	Route{"RaceDelete", "DELETE", "/api/races/{id}", DeleteHandler(m.Race{})},
	Route{"RaceUpdate", "PUT", "/api/races/{id}", UpdateHandler(m.Race{})},
}
