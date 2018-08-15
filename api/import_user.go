package api

import (
	"net/http"

	"github.com/hexaforce/cobra/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configImportUsersRouter(router *httprouter.Router) {
	router.GET("/importusers", GetAllImportUsers)
	router.POST("/importusers", AddImportUser)
	router.GET("/importusers/:id", GetImportUser)
	router.PUT("/importusers/:id", UpdateImportUser)
	router.DELETE("/importusers/:id", DeleteImportUser)
}

func GetAllImportUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	importusers := []model.ImportUser{}
	DB.Find(&importusers)
	writeJSON(w, &importusers)

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

	importusers := []*model.ImportUser{}

	if order != "" {
		err = DB.Model(&model.ImportUser{}).Order(order).Offset(offset).Limit(pagesize).Find(&importusers).Error
	} else {
		err = DB.Model(&model.ImportUser{}).Offset(offset).Limit(pagesize).Find(&importusers).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetImportUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	importuser := &model.ImportUser{}
	if DB.First(importuser, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, importuser)
}

func AddImportUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	importuser := &model.ImportUser{}

	if err := readJSON(r, importuser); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(importuser).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, importuser)
}

func UpdateImportUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	importuser := &model.ImportUser{}
	if DB.First(importuser, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.ImportUser{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(importuser, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(importuser).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, importuser)
}

func DeleteImportUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	importuser := &model.ImportUser{}

	if DB.First(importuser, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(importuser).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
