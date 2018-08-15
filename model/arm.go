package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type Arm struct {
	ID      int64  `gorm:"column:id;primary_key" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	Vos     string `gorm:"column:vos" json:"vos"`
	PsnlzID string `gorm:"column:psnlz_id" json:"psnlz_id"`
	LpTitle string `gorm:"column:lp_title" json:"lp_title"`
	URL     string `gorm:"column:url" json:"url"`
	Alpha   int    `gorm:"column:alpha" json:"alpha"`
	Beta    int    `gorm:"column:beta" json:"beta"`
}

// TableName sets the insert table name for this struct type
func (a *Arm) TableName() string {
	return "arms"
}

// Connect to the database
func Connect() (*gorm.DB, error) {
	mysqlConnect := "root:password@tcp(127.0.0.1:13306)/sakila"
	db, err := gorm.Open("mysql", mysqlConnect)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}
