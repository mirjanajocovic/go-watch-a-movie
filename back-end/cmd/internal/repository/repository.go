package repository

import (
	"backend/cmd/internal/models"
	"database/sql"
)

type DatabaseRepo interface {
	Connection () *sql.DB
	AllMovies() ([]*models.Movie, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
	OneMovie(id int) (*models.Movie, error)  // for public, be cause it is not connected with edit, for static display
	OneMovieForEdit(id int) (*models.Movie, []*models.Genre, error) 
	AllGenres() ([]*models.Genre, error)
	InsertMovie(movie models.Movie) (int, error)
	UpdateMovieGenres(id int, genreIDs []int) error
	UpdateMovie(movie models.Movie) error
	DeleteMovie(id int) error
}