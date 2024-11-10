package view

import "web_lab/internal/models"

type MoviesResponse struct {
	Movies []Movie `json:"movies"`
	Total  int64   `json:"total"`
}

type Movie struct {
	ID          int64  `json:"ID"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func ConvertMoviesResponse(movies []models.Movie, count int64) MoviesResponse {
	result := make([]Movie, 0, len(movies))

	for _, movie := range movies {
		result = append(result, Movie{
			ID:          movie.ID,
			Title:       movie.Title,
			Description: movie.Description,
			Image:       movie.Image,
		})
	}

	return MoviesResponse{
		Movies: result,
		Total:  count,
	}
}
