package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type JobCode struct {
	ID           int64  `gorm:"column:id;primary_key" json:"id"`
	LargeJobCode string `gorm:"column:large_job_code" json:"large_job_code"`
	LargeJobName string `gorm:"column:large_job_name" json:"large_job_name"`
	SmallJobCode string `gorm:"column:small_job_code" json:"small_job_code"`
	SmallJobName string `gorm:"column:small_job_name" json:"small_job_name"`
}

// TableName sets the insert table name for this struct type
func (j *JobCode) TableName() string {
	return "job_codes"
}
