package database

import (
	"gorm.io/gorm"
)

type option struct {
	tx          *gorm.DB
	debug       bool
	withMaster  bool
	withDeleted bool
	forUpdate   bool
}

type provider struct {
	db *gorm.DB
}

