package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/jinzhu/gorm"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type JobSegmentArm struct {
	ID           int64 `gorm:"column:id;primary_key" json:"id"`
	ArmID        int64 `gorm:"column:arm_id" json:"arm_id"`
	JobSegmentID int64 `gorm:"column:job_segment_id" json:"job_segment_id"`
}

// TableName sets the insert table name for this struct type
func (j *JobSegmentArm) TableName() string {
	return "job_segment_arms"
}
