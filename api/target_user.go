package api

import (
	"net/http"

	"github.com/hexaforce/cobra/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configTargetUsersRouter(router *httprouter.Router) {
	router.GET("/targetusers", GetAllTargetUsers)
	router.POST("/targetusers", AddTargetUser)
	router.GET("/targetusers/:id", GetTargetUser)
	router.PUT("/targetusers/:id", UpdateTargetUser)
	router.DELETE("/targetusers/:id", DeleteTargetUser)
}

func GetAllTargetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	targetusers := []model.TargetUser{}
	DB.Find(&targetusers)
	writeJSON(w, &targetusers)

	page, err := readInt(r, "page", 1)
	if err != nil || page < 1 {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	pagesize, err := readInt(r, "pagesize", 20)
	if err != nil || pagesize <= 0 {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	offset := (page - 1) * pagesize

	order := r.FormValue("order")

	targetusers := []*model.TargetUser{}

	if order != "" {
		err = DB.Model(&model.TargetUser{}).Order(order).Offset(offset).Limit(pagesize).Find(&targetusers).Error
	} else {
		err = DB.Model(&model.TargetUser{}).Offset(offset).Limit(pagesize).Find(&targetusers).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetTargetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	targetuser := &model.TargetUser{}
	if DB.First(targetuser, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, targetuser)
}

func AddTargetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	targetuser := &model.TargetUser{}

	if err := readJSON(r, targetuser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(targetuser).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, targetuser)
}

func UpdateTargetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	targetuser := &model.TargetUser{}
	if DB.First(targetuser, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.TargetUser{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(targetuser, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(targetuser).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, targetuser)
}

func DeleteTargetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	targetuser := &model.TargetUser{}

	if DB.First(targetuser, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(targetuser).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
