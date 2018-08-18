package api

import (
	"net/http"

	"github.com/hexaforce/cobra/model"
	"github.com/julienschmidt/httprouter"
	"github.com/smallnest/gen/dbmeta"
)

func configSchemaMigrationsRouter(router *httprouter.Router) {
	router.GET("/schemamigrations", GetAllSchemaMigrations)
	router.POST("/schemamigrations", AddSchemaMigration)
	router.GET("/schemamigrations/:id", GetSchemaMigration)
	router.PUT("/schemamigrations/:id", UpdateSchemaMigration)
	router.DELETE("/schemamigrations/:id", DeleteSchemaMigration)
}

func GetAllSchemaMigrations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	schemamigrations := []model.SchemaMigration{}
	DB.Find(&schemamigrations)
	writeJSON(w, &schemamigrations)

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

	// schemamigrations := []*model.SchemaMigration{}

	// if order != "" {
	// 	err = DB.Model(&model.SchemaMigration{}).Order(order).Offset(offset).Limit(pagesize).Find(&schemamigrations).Error
	// } else {
	// 	err = DB.Model(&model.SchemaMigration{}).Offset(offset).Limit(pagesize).Find(&schemamigrations).Error
	// }

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
}

func GetSchemaMigration(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	schemamigration := &model.SchemaMigration{}
	if DB.First(schemamigration, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	writeJSON(w, schemamigration)
}

func AddSchemaMigration(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	schemamigration := &model.SchemaMigration{}

	if err := readJSON(r, schemamigration); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := DB.Save(schemamigration).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, schemamigration)
}

func UpdateSchemaMigration(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	schemamigration := &model.SchemaMigration{}
	if DB.First(schemamigration, id).Error != nil {
		http.NotFound(w, r)
		return
	}

	updated := &model.SchemaMigration{}
	if err := readJSON(r, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := dbmeta.Copy(schemamigration, updated); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := DB.Save(schemamigration).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, schemamigration)
}

func DeleteSchemaMigration(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	schemamigration := &model.SchemaMigration{}

	if DB.First(schemamigration, id).Error != nil {
		http.NotFound(w, r)
		return
	}
	if err := DB.Delete(schemamigration).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
