package domain

import "time"

type Post struct {
	ID        int
	Title     string
	Content   string
	AuthorID  int
	CreatedAt time.Time
}

type PostRepository interface {
	FindById(id int) (*Post, error)
	FindAll() (*[]Post, error)
	Create(post *Post) error
}
