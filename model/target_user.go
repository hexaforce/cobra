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

type TargetUser struct {
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	LeadID     string    `gorm:"column:lead_id" json:"lead_id"`
	JobCode    string    `gorm:"column:job_code" json:"job_code"`
	AddedAt    time.Time `gorm:"column:added_at" json:"added_at"`
	TargetDate time.Time `gorm:"column:target_date" json:"target_date"`
	ImportedAt time.Time `gorm:"column:imported_at" json:"imported_at"`
}

// TableName sets the insert table name for this struct type
func (t *TargetUser) TableName() string {
	return "target_users"
}
