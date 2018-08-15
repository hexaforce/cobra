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

type ImportUser struct {
	ID               int64       `gorm:"column:id;primary_key" json:"id"`
	ArmID            int64       `gorm:"column:arm_id" json:"arm_id"`
	ArmName          string      `gorm:"column:arm_name" json:"arm_name"`
	LeadID           string      `gorm:"column:lead_id" json:"lead_id"`
	PsnlzID          string      `gorm:"column:psnlz_id" json:"psnlz_id"`
	LpTitle          string      `gorm:"column:lp_title" json:"lp_title"`
	TargetDate       time.Time   `gorm:"column:target_date" json:"target_date"`
	CalculatedArms   null.String `gorm:"column:calculated_arms" json:"calculated_arms"`
	ImportExecutedAt null.Time   `gorm:"column:import_executed_at" json:"import_executed_at"`
}

// TableName sets the insert table name for this struct type
func (i *ImportUser) TableName() string {
	return "import_users"
}
