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

type JobSegment struct {
	ID                  int64  `gorm:"column:id;primary_key" json:"id"`
	JobCode             string `gorm:"column:job_code" json:"job_code"`
	AggregatePeriodDays int    `gorm:"column:aggregate_period_days" json:"aggregate_period_days"`
}

// TableName sets the insert table name for this struct type
func (j *JobSegment) TableName() string {
	return "job_segments"
}
