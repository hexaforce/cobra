package api

import (
	"net/http"

	"github.com/hexaforce/cobra/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configJobCodesRouter(router *httprouter.Router) {
	router.GET("/jobcodes", GetAllJobCodes)
	router.POST("/jobcodes", AddJobCode)
	router.GET("/jobcodes/:id", GetJobCode)
	router.PUT("/jobcodes/:id", UpdateJobCode)
	router.DELETE("/jobcodes/:id", DeleteJobCode)
}

func GetAllJobCodes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jobcodes := []model.JobCode{}
	DB.Find(&jobcodes)
	writeJSON(w, &jobcodes)

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

	// jobcodes := []*model.JobCode{}

	// if order != "" {
	// 	err = DB.Model(&model.JobCode{}).Order(order).Offset(offset).Limit(pagesize).Find(&jobcodes).Error
	// } else {
	// 	err = DB.Model(&model.JobCode{}).Offset(offset).Limit(pagesize).Find(&jobcodes).Error
	// }

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

func GetJobCode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	jobcode := &model.JobCode{}
	if DB.First(jobcode, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, jobcode)
}

func AddJobCode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	jobcode := &model.JobCode{}

	if err := readJSON(r, jobcode); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(jobcode).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, jobcode)
}

func UpdateJobCode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	jobcode := &model.JobCode{}
	if DB.First(jobcode, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.JobCode{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(jobcode, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(jobcode).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, jobcode)
}

func DeleteJobCode(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	jobcode := &model.JobCode{}

	if DB.First(jobcode, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(jobcode).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
