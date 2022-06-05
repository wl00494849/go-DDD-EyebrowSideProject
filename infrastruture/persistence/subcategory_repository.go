package persistence

import "database/sql"

type SubcategoryRepo struct {
	db *sql.DB
}

func NewSubcategoryRepo(db *sql.DB) *SubcategoryRepo {
	return &SubcategoryRepo{db}
}
