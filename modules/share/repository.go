package share

import "github.com/go-pg/pg/v10"

// Repository is a DB Connection to postgres
type Repository struct {
	DBRead, DBWrite *pg.DB
}

// NewRepository create new repository with read & write DB
func NewRepository(dbRead, dbWrite *pg.DB) *Repository {
	return &Repository{DBRead: dbRead, DBWrite: dbWrite}
}
