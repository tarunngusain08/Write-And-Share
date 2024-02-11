package repo

import (
	"Write-And-Share/contracts"
	"database/sql"
	"fmt"
	"strings"
)

type ShareNoteRepo struct {
	db *sql.DB
}

func NewShareNoteRepo(db *sql.DB) *ShareNoteRepo {
	return &ShareNoteRepo{db: db}
}

const shareNote = `INSERT INTO notes_to_user_mapping (note_id, username) VALUES`

func (g *ShareNoteRepo) ShareNote(request *contracts.ShareNoteRequest) error {

	var queryBuilder strings.Builder
	queryBuilder.WriteString(shareNote)
	valueStrings := make([]string, 0, len(request.ShareUsernames))
	valueArgs := make([]interface{}, 0, len(request.ShareUsernames)*2) // Two placeholders per username (note_id, username)
	for _, username := range request.ShareUsernames {
		valueStrings = append(valueStrings, "(?, ?)")
		valueArgs = append(valueArgs, request.NoteId, username)
	}
	queryBuilder.WriteString(strings.Join(valueStrings, ","))
	res, err := g.db.Exec(queryBuilder.String(), request.NoteId, request.Username)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil || int(rows) == len(request.ShareUsernames) {
		return fmt.Errorf("failed to share note with all the users: %v", err)
	}
	return nil
}
