package repo

import (
	"Write-And-Share/contracts"
	"database/sql"
	"errors"
)

type LoginRepo struct {
	db *sql.DB
}

func NewLoginRepo(db *sql.DB) *LoginRepo {
	return &LoginRepo{db: db}
}

const login = `SELECT COUNT(*) FROM Users WHERE username = $1 AND password = $2`

func (l *LoginRepo) Login(details *contracts.UserDetails) error {
	var count int
	err := l.db.QueryRow(login, details.UserName, details.Password).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("wrong username or password")
	}
	return nil
}
