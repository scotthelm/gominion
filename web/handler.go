package web

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux"
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
		fmt.Println(ts)
		model := reflect.New(reflect.TypeOf(t)).Interface()
		Ctx.Db.Limit(limit).Offset(offset).Order(fmt.Sprintf("%s %s", orderBy, direction)).Find(ts)
		Ctx.Db.Model(model).Count(&count)
		result := reflect.ValueOf(ts).Elem()
		for i := 0; i < result.Len(); i++ {
			for _, prof := range preloads {
				val := reflect.New(result.Index(i).FieldByName(prof).Type()).Interface()
				fmt.Println(reflect.TypeOf(val))
				Ctx.Db.Model(result.Index(i).Interface()).
					Association(prof).
					Find(val)
				result.Index(i).FieldByName(prof).Set(reflect.Indirect(reflect.ValueOf(val)))
			}
		}
		json.NewEncoder(w).Encode(QueryResult{page, perPage, ts, count})
	}
}

func doPreloads(collection []interface{}, preload []string) {
	for i, _ := range collection {
		for _, v := range preload {
			fmt.Println(i, v)
		}
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
