package repo

import (
	"Write-And-Share/contracts"
	"database/sql"
	"fmt"
)

type DeleteNoteRepo struct {
	db *sql.DB
}

func NewDeleteNoteRepo(db *sql.DB) *DeleteNoteRepo {
	return &DeleteNoteRepo{db: db}
}

const deleteNote = `DELETE FROM notes WHERE note_id = $1 AND username = $2;`

func (g *DeleteNoteRepo) DeleteNote(request *contracts.DeleteNoteRequest) error {

	res, err := g.db.Exec(deleteNote, request.NoteId, request.Username)
	if err != nil {
		return fmt.Errorf("failed to create note: %v", err)
	}

	rows, err := res.RowsAffected()
	if err != nil || rows == 0 {
		return fmt.Errorf("failed to create note: %v", err)
	}
	return nil
}
