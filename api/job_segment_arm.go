package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configJobSegmentArmsRouter(router *httprouter.Router) {
	router.GET("/jobsegmentarms", GetAllJobSegmentArms)
	router.POST("/jobsegmentarms", AddJobSegmentArm)
	router.GET("/jobsegmentarms/:id", GetJobSegmentArm)
	router.PUT("/jobsegmentarms/:id", UpdateJobSegmentArm)
	router.DELETE("/jobsegmentarms/:id", DeleteJobSegmentArm)
}

func GetAllJobSegmentArms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jobsegmentarms := []model.JobSegmentArm{}
	DB.Find(&jobsegmentarms)
	writeJSON(w, &jobsegmentarms)

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

	// jobsegmentarms := []*model.JobSegmentArm{}

	// if order != "" {
	// 	err = DB.Model(&model.JobSegmentArm{}).Order(order).Offset(offset).Limit(pagesize).Find(&jobsegmentarms).Error
	// } else {
	// 	err = DB.Model(&model.JobSegmentArm{}).Offset(offset).Limit(pagesize).Find(&jobsegmentarms).Error
	// }

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

func GetJobSegmentArm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	jobsegmentarm := &model.JobSegmentArm{}
	if DB.First(jobsegmentarm, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, jobsegmentarm)
}

func AddJobSegmentArm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jobsegmentarm := &model.JobSegmentArm{}

	if err := readJSON(r, jobsegmentarm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(jobsegmentarm).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, jobsegmentarm)
}

func UpdateJobSegmentArm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	jobsegmentarm := &model.JobSegmentArm{}
	if DB.First(jobsegmentarm, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.JobSegmentArm{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(jobsegmentarm, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(jobsegmentarm).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, jobsegmentarm)
}

func DeleteJobSegmentArm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	jobsegmentarm := &model.JobSegmentArm{}

	if DB.First(jobsegmentarm, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(jobsegmentarm).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
