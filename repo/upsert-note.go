package repo

import (
	"Write-And-Share/contracts"
	"database/sql"
	"fmt"
)

type UpsertNoteRepo struct {
	db *sql.DB
}

func NewUpsertNoteRepo(db *sql.DB) *UpsertNoteRepo {
	return &UpsertNoteRepo{db: db}
}

const createNote = `INSERT INTO notes VALUES ($1, $2, $3, $4);`

const addtoNoteToUserMapping = `INSERT INTO notes_to_user_mapping VALUES ($1, $2, $3, $4)`

const updateNote = `UPDATE notes SET content = $1, time = $2, heading = $3 WHERE note_id = $4 AND username = $5;`

func (g *UpsertNoteRepo) CreateNote(request *contracts.UpsertNoteRequest) (int64, error) {

	res, err := g.db.Exec(createNote, request.Username, request.Content, request.Heading, request.Time)
	if err != nil {
		return 0, fmt.Errorf("failed to create note: %v", err)
	}

	rows, err := res.RowsAffected()
	if err != nil || rows == 0 {
		return 0, fmt.Errorf("failed to create note: %v", err)
	}
	return res.LastInsertId()
}

func (g *UpsertNoteRepo) AddtoNoteToUserMapping(noteId int64) error {

	res, err := g.db.Exec(addtoNoteToUserMapping, noteId)
	if err != nil {
		return fmt.Errorf("failed to create note: %v", err)
	}

	rows, err := res.RowsAffected()
	if err != nil || rows == 0 {
		return fmt.Errorf("failed to create note: %v", err)
	}
	return nil
}

func (g *UpsertNoteRepo) UpdateNote(request *contracts.UpsertNoteRequest) error {

	res, err := g.db.Exec(updateNote, request.Content, request.Time, request.Heading, request.NoteId, request.Username)
	if err != nil {
		return fmt.Errorf("failed to update note: %v", err)
	}

	rows, err := res.RowsAffected()
	if err != nil || rows == 0 {
		return fmt.Errorf("failed to create note: %v", err)
	}
	return nil
}
