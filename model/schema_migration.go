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

type SchemaMigration struct {
	Version int64 `gorm:"column:version;primary_key" json:"version"`
	Dirty   int   `gorm:"column:dirty" json:"dirty"`
}

// TableName sets the insert table name for this struct type
func (s *SchemaMigration) TableName() string {
	return "schema_migrations"
}
