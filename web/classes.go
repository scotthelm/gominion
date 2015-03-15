package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/scotthelm/gominion/models"
)

func ClassesIndexHandler(w http.ResponseWriter, r *http.Request) {

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

	var classes []models.Class
	var model models.Class
	Ctx.Db.Limit(limit).Offset(offset).Order(fmt.Sprintf("%s %s", orderBy, direction)).Find(&classes)
	for i, _ := range classes {
		Ctx.Db.Model(&classes[i]).Association("Proficiencies").Find(&classes[i].Proficiencies)
		Ctx.Db.Model(&classes[i]).Association("Skills").Find(&classes[i].Skills)
	}

	fmt.Println(classes)
	Ctx.Db.Model(model).Count(&count)
	json.NewEncoder(w).Encode(QueryResult{page, perPage, classes, count})
}
