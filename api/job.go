package api

import (
	"net/http"

	"github.com/hexaforce/cobra/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configJobsRouter(router *httprouter.Router) {
	router.GET("/jobs", GetAllJobs)
	router.POST("/jobs", AddJob)
	router.GET("/jobs/:id", GetJob)
	router.PUT("/jobs/:id", UpdateJob)
	router.DELETE("/jobs/:id", DeleteJob)
}

func GetAllJobs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jobs := []model.Job{}
	DB.Find(&jobs)
	writeJSON(w, &jobs)

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

	jobs := []*model.Job{}

	if order != "" {
		err = DB.Model(&model.Job{}).Order(order).Offset(offset).Limit(pagesize).Find(&jobs).Error
	} else {
		err = DB.Model(&model.Job{}).Offset(offset).Limit(pagesize).Find(&jobs).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	job := &model.Job{}
	if DB.First(job, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, job)
}

func AddJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	job := &model.Job{}

	if err := readJSON(r, job); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(job).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, job)
}

func UpdateJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	job := &model.Job{}
	if DB.First(job, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Job{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(job, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(job).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, job)
}

func DeleteJob(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	job := &model.Job{}

	if DB.First(job, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(job).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
