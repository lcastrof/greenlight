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
	Tokens TokenModel // Add a new Tokens field.
	Users  UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Movies: MovieModel{DB: db},
		Tokens: TokenModel{DB: db}, // Initialize a new TokenModel instance.
		Users:  UserModel{DB: db},
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
