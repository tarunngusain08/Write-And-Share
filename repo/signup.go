package repo

import (
	"Write-And-Share/contracts"
	"database/sql"
	"errors"
)

type SignUpRepo struct {
	db *sql.DB
}

func NewSignUpRepo(db *sql.DB) *SignUpRepo {
	return &SignUpRepo{db: db}
}

const signUp = `INSERT INTO Users VALUES ($1, $2)`

func (s *SignUpRepo) SignUp(details *contracts.UserDetails) error {
	res, err := s.db.Exec(signUp, details.UserName, details.Password)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("signUp was not successful")
	}
	return nil
}
