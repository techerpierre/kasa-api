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

type CommentFiltersDTO struct {
	Content         *string `form:"content"`
	AccommodationID *string `form:"accommodationId"`
	UserID          *string `form:"userId"`
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

func MakeCommentFilters(filterDTO CommentFiltersDTO) []entities.Filter {
	return []entities.Filter{
		{Field: "content", Value: filterDTO.Content},
		{Field: "accommodationId", Value: filterDTO.AccommodationID},
		{Field: "userId", Value: filterDTO.UserID},
	}
}
