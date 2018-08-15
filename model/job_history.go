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

type JobHistory struct {
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	JobID      int64     `gorm:"column:job_id" json:"job_id"`
	Kind       string    `gorm:"column:kind" json:"kind"`
	Status     string    `gorm:"column:status" json:"status"`
	TargetDate time.Time `gorm:"column:target_date" json:"target_date"`
	ExecutedAt time.Time `gorm:"column:executed_at" json:"executed_at"`
}

// TableName sets the insert table name for this struct type
func (j *JobHistory) TableName() string {
	return "job_histories"
}
