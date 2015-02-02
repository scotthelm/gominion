package web

import (
	"encoding/json"
	"github.com/scotthelm/gominion/db"
	m "github.com/scotthelm/gominion/models"
	"net/http"
)

var Ctx db.Context

func CampaignIndex(w http.ResponseWriter, r *http.Request) {
	campaigns := m.Campaigns{
		m.Campaign{1, "The Moors Of Tel", ""},
		m.Campaign{1, "The Legend of Malganis", ""},
	}
	json.NewEncoder(w).Encode(campaigns)
}

func CampaignShow(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(m.Campaign{1, "The Moors of Tel", ""})
}

func CampaignPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var c m.Campaign
	decoder.Decode(&c)
	Ctx.DB.Save(&c)

}
