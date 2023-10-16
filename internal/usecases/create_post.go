package usecases

import (
	"github.com/altheaix/CleanArchitecture/internal/domain"
	"time"
)

type CreatePostInput struct {
	Title    string
	Content  string
	AuthorID int
}

type CreatePostOutput struct {
	ID        int
	Title     string
	Content   string
	AuthorID  int
	CreatedAt time.Time
}

type CreatePostUseCase interface {
	Execute(input CreatePostInput) (*CreatePostOutput, error)
}

func NewCreatePostUseCase(postRepo domain.Post) CreatePostUseCase {
	return &createPostUseCase{postRepo: postRepo}
}

type createPostUseCase struct {
	postRepo domain.Post
}

func (c *createPostUseCase) Execute(input CreatePostInput) (*CreatePostOutput, error) {
	return nil, nil
}
