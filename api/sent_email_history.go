package api

import (
	"net/http"

	"github.com/hexaforce/cobra/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configSentEmailHistoriesRouter(router *httprouter.Router) {
	router.GET("/sentemailhistories", GetAllSentEmailHistories)
	router.POST("/sentemailhistories", AddSentEmailHistory)
	router.GET("/sentemailhistories/:id", GetSentEmailHistory)
	router.PUT("/sentemailhistories/:id", UpdateSentEmailHistory)
	router.DELETE("/sentemailhistories/:id", DeleteSentEmailHistory)
}

func GetAllSentEmailHistories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sentemailhistories := []model.SentEmailHistory{}
	DB.Find(&sentemailhistories)
	writeJSON(w, &sentemailhistories)

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

	sentemailhistories := []*model.SentEmailHistory{}

	if order != "" {
		err = DB.Model(&model.SentEmailHistory{}).Order(order).Offset(offset).Limit(pagesize).Find(&sentemailhistories).Error
	} else {
		err = DB.Model(&model.SentEmailHistory{}).Offset(offset).Limit(pagesize).Find(&sentemailhistories).Error
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetSentEmailHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	sentemailhistory := &model.SentEmailHistory{}
	if DB.First(sentemailhistory, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, sentemailhistory)
}

func AddSentEmailHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sentemailhistory := &model.SentEmailHistory{}

	if err := readJSON(r, sentemailhistory); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(sentemailhistory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, sentemailhistory)
}

func UpdateSentEmailHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	sentemailhistory := &model.SentEmailHistory{}
	if DB.First(sentemailhistory, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.SentEmailHistory{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(sentemailhistory, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(sentemailhistory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, sentemailhistory)
}

func DeleteSentEmailHistory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	sentemailhistory := &model.SentEmailHistory{}

	if DB.First(sentemailhistory, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(sentemailhistory).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
