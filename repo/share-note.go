package repo

import (
	"Write-And-Share/contracts"
	"database/sql"
	"fmt"
)

type ShareNoteRepo struct {
	db *sql.DB
}

func NewShareNoteRepo(db *sql.DB) *ShareNoteRepo {
	return &ShareNoteRepo{db: db}
}

const shareNote = `DELETE FROM notes WHERE note_id = $1 AND username = $2;`

func (g *ShareNoteRepo) ShareNote(request *contracts.ShareNoteRequest) error {

	res, err := g.db.Exec(deleteNote, request.NoteId, request.Username)
	if err != nil {
		return fmt.Errorf("failed to create note: %v", err)
	}

	rows, err := res.RowsAffected()
	if err != nil || int(rows) == len(request.ShareUsernames) {
		return fmt.Errorf("failed to create note: %v", err)
	}
	return nil
}
