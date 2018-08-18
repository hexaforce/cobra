package api

import (
	"net/http"

	"github.com/hexaforce/cobra/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configArmsRouter(router *httprouter.Router) {
	router.GET("/arms", GetAllArms)
	router.POST("/arms", AddArm)
	router.GET("/arms/:id", GetArm)
	router.PUT("/arms/:id", UpdateArm)
	router.DELETE("/arms/:id", DeleteArm)
}

func GetAllArms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	arms := []model.Arm{}
	DB.Find(&arms)
	writeJSON(w, &arms)

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

	// arms := []*model.Arm{}

	// if order != "" {
	// 	err = DB.Model(&model.Arm{}).Order(order).Offset(offset).Limit(pagesize).Find(&arms).Error
	// } else {
	// 	err = DB.Model(&model.Arm{}).Offset(offset).Limit(pagesize).Find(&arms).Error
	// }

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

func GetArm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	arm := &model.Arm{}
	if DB.First(arm, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, arm)
}

func AddArm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	arm := &model.Arm{}

	if err := readJSON(r, arm); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(arm).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, arm)
}

func UpdateArm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	arm := &model.Arm{}
	if DB.First(arm, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.Arm{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(arm, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(arm).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, arm)
}

func DeleteArm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	arm := &model.Arm{}

	if DB.First(arm, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(arm).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
