package dto

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type CommentDTO struct {
	ID              string `json:"id"`
	Content         string `json:"content"`
	AccommodationID string `json:"accomodationId"`
	UserID          string `json:"userId"`
}

type CommentInputDTO struct {
	Content         string `json:"content"`
	AccommodationID string `json:"accomodationId"`
	UserID          string `json:"userId"`
}

func PipeCommentInDTO(source *entities.Comment, target *CommentDTO) {
	target.ID = source.ID
	target.Content = source.Content
	target.AccommodationID = source.AccommodationID
	target.UserID = source.UserID
}

func PipeInputDTOInComment(source *CommentInputDTO, target *entities.Comment) {
	target.Content = source.Content
	target.AccommodationID = source.AccommodationID
	target.UserID = source.UserID
}
