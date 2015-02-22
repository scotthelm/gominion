package web

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/scotthelm/gominion/db"
	// "log"
	"net/http"
	"reflect"
	"strconv"
)

var Ctx db.Context

type QueryResult struct {
	Page         string      `json:"page"`
	PerPage      string      `json:"per_page"`
	Result       interface{} `json:"result"`
	TotalRecords int64       `json:"total_records"`
}

func IndexHandler(t interface{}, preloads ...string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var orderBy, direction, page, perPage string
		var offset, limit, count int64

		count = 4

		orderBy = r.FormValue("order_by")
		if orderBy == "" {
			orderBy = "id"
		}

		direction = r.FormValue("direction")
		if direction == "" {
			direction = "asc"
		}

		page = r.FormValue("page")
		if page == "" {
			page = "1"
		}

		perPage = r.FormValue("per_page")
		if perPage == "" {
			perPage = "10"
		}

		ipage, _ := strconv.ParseInt(page, 10, 64)
		iperPage, _ := strconv.ParseInt(perPage, 10, 64)
		offset = (ipage * iperPage) - iperPage
		limit = iperPage

		ts := reflect.New(reflect.SliceOf(reflect.TypeOf(t))).Interface()
		model := reflect.New(reflect.TypeOf(t)).Interface()
		pdb := doPreloads(&Ctx.Db, preloads)
		pdb.Limit(limit).Offset(offset).Order(fmt.Sprintf("%s %s", orderBy, direction)).Find(ts)
		Ctx.Db.Model(model).Count(&count)
		json.NewEncoder(w).Encode(QueryResult{page, perPage, ts, count})
	}
}

func doPreloads(db *gorm.DB, preload []string) *gorm.DB {
	for _, v := range preload {
		db = db.Preload(v)
	}
	return db
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
