package dto

type ResponseDTO[T any] struct {
	StatusCode int  `json:"statusCode"`
	Data       T    `json:"data"`
	Count      *int `json:"count"`
}

func CreateResponse[T any](statusCode int, data T, count *int) ResponseDTO[T] {
	return ResponseDTO[T]{
		StatusCode: statusCode,
		Data:       data,
		Count:      count,
	}
}
