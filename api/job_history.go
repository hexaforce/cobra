package api

import (
	"net/http"

	"github.com/hexaforce/cobra/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configJobHistoriesRouter(router *httprouter.Router) {
	router.GET("/jobhistories", GetAllJobHistories)
	router.POST("/jobhistories", AddJobHistory)
	router.GET("/jobhistories/:id", GetJobHistory)
	router.PUT("/jobhistories/:id", UpdateJobHistory)
	router.DELETE("/jobhistories/:id", DeleteJobHistory)
}

func GetAllJobHistories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jobhistories := []model.JobHistory{}
	DB.Find(&jobhistories)
	writeJSON(w, &jobhistories)

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

	jobhistories := []*model.JobHistory{}

	if order != "" {
		err = DB.Model(&model.JobHistory{}).Order(order).Offset(offset).Limit(pagesize).Find(&jobhistories).Error
	} else {
		err = DB.Model(&model.JobHistory{}).Offset(offset).Limit(pagesize).Find(&jobhistories).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetJobHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	jobhistory := &model.JobHistory{}
	if DB.First(jobhistory, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, jobhistory)
}

func AddJobHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jobhistory := &model.JobHistory{}

	if err := readJSON(r, jobhistory); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(jobhistory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, jobhistory)
}

func UpdateJobHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	jobhistory := &model.JobHistory{}
	if DB.First(jobhistory, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.JobHistory{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(jobhistory, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(jobhistory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, jobhistory)
}

func DeleteJobHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	jobhistory := &model.JobHistory{}

	if DB.First(jobhistory, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(jobhistory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
