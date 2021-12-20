package repository

import "github.com/jmoiron/sqlx"

type TodoList interface {

}

type Repository struct {
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}