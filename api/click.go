package api

import (
	"net/http"

	"github.com/hexaforce/cobra/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configClicksRouter(router *httprouter.Router) {
	router.GET("/clicks", GetAllClicks)
	router.POST("/clicks", AddClick)
	router.GET("/clicks/:id", GetClick)
	router.PUT("/clicks/:id", UpdateClick)
	router.DELETE("/clicks/:id", DeleteClick)
}

func GetAllClicks(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	clicks := []model.Click{}
	DB.Find(&clicks)
	writeJSON(w, &clicks)

	// page, err := readInt(r, "page", 1)
	// if err != nil || page < 1 {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// }
	// pagesize, err := readInt(r, "pagesize", 20)
	// if err != nil || pagesize <= 0 {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// }
	// offset := (page - 1) * pagesize

	// order := r.FormValue("order")

	// clicks := []*model.Click{}

	// if order != "" {
	// 	err = DB.Model(&model.Click{}).Order(order).Offset(offset).Limit(pagesize).Find(&clicks).Error
	// } else {
	// 	err = DB.Model(&model.Click{}).Offset(offset).Limit(pagesize).Find(&clicks).Error
	// }

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

func GetClick(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	click := &model.Click{}
	if DB.First(click, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, click)
}

func AddClick(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	click := &model.Click{}

	if err := readJSON(r, click); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(click).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, click)
}

func UpdateClick(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	click := &model.Click{}
	if DB.First(click, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Click{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(click, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(click).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, click)
}

func DeleteClick(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	click := &model.Click{}

	if DB.First(click, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(click).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
