package repository

import (
	"database/sql"
	"github.com/altheaix/CleanArchitecture/internal/domain"
)

type PostSQLRepository struct {
	db *sql.DB
}

func (r *PostSQLRepository) FindByID(id int) (*domain.Post, error) {
	return nil, nil
}

func (r *PostSQLRepository) FindAll() (*[]domain.Post, error) {
	return nil, nil
}

func (r *PostSQLRepository) Create(post *domain.Post) error {
	return nil
}
