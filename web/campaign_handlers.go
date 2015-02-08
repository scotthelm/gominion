package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/scotthelm/gominion/db"
	m "github.com/scotthelm/gominion/models"
	"net/http"
	"strconv"
)

var Ctx db.Context

func CampaignIndex(w http.ResponseWriter, r *http.Request) {
	var campaigns []m.Campaign
	Ctx.Db.Find(&campaigns)
	json.NewEncoder(w).Encode(campaigns)
}

func CampaignShow(w http.ResponseWriter, r *http.Request) {
	var campaign m.Campaign
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err == nil {
		Ctx.Db.Find(&campaign, id)
		json.NewEncoder(w).Encode(campaign)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{ \"ok\" : false, \"error\" : \"failed to parse id\"}"))
	}
}

func CampaignPost(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var c m.Campaign
	decoder.Decode(&c)
	Ctx.Db.Save(&c)
	json.NewEncoder(w).Encode(c)

}

func CampaignDelete(w http.ResponseWriter, r *http.Request) {
	var campaign m.Campaign
	id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	if err == nil {
		Ctx.Db.Find(&campaign, id)
		Ctx.Db.Delete(&campaign)
		json.NewEncoder(w).Encode(campaign)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("{ \"ok\" : false, \"error\" : \"failed to parse id\"}"))
	}

}
