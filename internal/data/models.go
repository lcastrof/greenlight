package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrEditConflict   = errors.New("edit conflict")
)

type Models struct {
	Movies MovieModel
	Users  UserModel // Add a new Users field.
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
		Users:  UserModel{DB: db}, // Initialize a new UserModel instance.
	}
}

// Create a helper function which returns a Models instance containing the mock models
// only.
// func NewMockModels() Models {
// 	movies := make([]Movie, 0)
// 	return Models{
// 		Movies: MockMovieModel{
// 			DB: &movies,
// 		},
// 	}
// }
