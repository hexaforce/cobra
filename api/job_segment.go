package api

import (
	"net/http"

	"github.com/hexaforce/cobra/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configJobSegmentsRouter(router *httprouter.Router) {
	router.GET("/jobsegments", GetAllJobSegments)
	router.POST("/jobsegments", AddJobSegment)
	router.GET("/jobsegments/:id", GetJobSegment)
	router.PUT("/jobsegments/:id", UpdateJobSegment)
	router.DELETE("/jobsegments/:id", DeleteJobSegment)
}

func GetAllJobSegments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jobsegments := []model.JobSegment{}
	DB.Find(&jobsegments)
	writeJSON(w, &jobsegments)

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

	jobsegments := []*model.JobSegment{}

	if order != "" {
		err = DB.Model(&model.JobSegment{}).Order(order).Offset(offset).Limit(pagesize).Find(&jobsegments).Error
	} else {
		err = DB.Model(&model.JobSegment{}).Offset(offset).Limit(pagesize).Find(&jobsegments).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetJobSegment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	jobsegment := &model.JobSegment{}
	if DB.First(jobsegment, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, jobsegment)
}

func AddJobSegment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jobsegment := &model.JobSegment{}

	if err := readJSON(r, jobsegment); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(jobsegment).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, jobsegment)
}

func UpdateJobSegment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	jobsegment := &model.JobSegment{}
	if DB.First(jobsegment, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.JobSegment{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(jobsegment, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(jobsegment).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, jobsegment)
}

func DeleteJobSegment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	jobsegment := &model.JobSegment{}

	if DB.First(jobsegment, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(jobsegment).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
