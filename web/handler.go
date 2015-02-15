package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/scotthelm/gominion/db"
	// "log"
	"net/http"
	"reflect"
	"strconv"
)

var Ctx db.Context

func IndexHandler(t interface{}) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ts := reflect.New(reflect.TypeOf(t)).Interface()
		Ctx.Db.Find(ts)
		json.NewEncoder(w).Encode(ts)
	}
}

func ShowHandler(t interface{}) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		st := reflect.New(reflect.TypeOf(t)).Interface()
		id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		if err == nil {
			Ctx.Db.Find(st, id)
			json.NewEncoder(w).Encode(st)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{ \"ok\" : false, \"error\" : \"failed to parse id\"}"))
		}
	}
}

func CreateHandler(t interface{}) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		item := reflect.New(reflect.TypeOf(t)).Interface()
		decoder.Decode(item)
		Ctx.Db.Save(item)
		json.NewEncoder(w).Encode(item)
	}
}

func DeleteHandler(t interface{}) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		dt := reflect.New(reflect.TypeOf(t)).Interface()
		id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		if err == nil {
			Ctx.Db.Find(dt, id)
			Ctx.Db.Delete(dt)
			json.NewEncoder(w).Encode(dt)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{ \"ok\" : false, \"error\" : \"failed to parse id\"}"))
		}
	}
}

func UpdateHandler(t interface{}) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		ut := reflect.New(reflect.TypeOf(t)).Interface()
		id, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
		if err == nil {
			decoder.Decode(ut)
			reflect.ValueOf(ut).Elem().FieldByName("Id").SetInt(id)
			Ctx.Db.Save(ut)
			json.NewEncoder(w).Encode(ut)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("{ \"ok\" : false, \"error\" : \"failed to parse id\"}"))
		}
	}
}
