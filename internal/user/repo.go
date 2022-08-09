package user

import (
	"go-wire-architecture/pkg"
	"gorm.io/gorm"
)

// Repository database structure
type Repository struct {
	pkg.Database
	logger pkg.Logger
}

// NewRepository creates a new user repository
func NewRepository(db pkg.Database, logger pkg.Logger) Repository {
	return Repository{db, logger}
}

// WithTrx delegate transaction from user repository
func (r Repository) WithTrx(trxHandle *gorm.DB) Repository {
	if trxHandle != nil {
		r.logger.Debug("using WithTrx as trxHandle is not nil")
		r.Database.DB = trxHandle
	}
	return r
}
