package service

import (
	"forum/internal/repository"
	"forum/models"
)

type CommentService struct {
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (s *CommentService) CreateComment(comment models.Comment) error {
	return s.repo.CreateComment(comment)
}

func (s *CommentService) GetCommentByPostID(id int) ([]models.Comment, error) {
	return s.repo.GetCommentByPostID(id)
}
