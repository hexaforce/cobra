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

type Click struct {
	ID         int64     `gorm:"column:id;primary_key" json:"id"`
	LeadID     string    `gorm:"column:lead_id" json:"lead_id"`
	Vos        string    `gorm:"column:vos" json:"vos"`
	TargetDate time.Time `gorm:"column:target_date" json:"target_date"`
	ImportedAt time.Time `gorm:"column:imported_at" json:"imported_at"`
}

// TableName sets the insert table name for this struct type
func (c *Click) TableName() string {
	return "clicks"
}
