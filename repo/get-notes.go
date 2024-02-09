package repo

import (
	"Write-And-Share/contracts"
	"database/sql"
	"fmt"
)

type GetNotesRepo struct {
	db *sql.DB
}

func NewGetNotesRepo(db *sql.DB) *GetNotesRepo {
	return &GetNotesRepo{db: db}
}

const getAllNotes = `SELECT title, content FROM notes_to_user_mapping n WHERE user_name = $1;`

const getNoteById = `SELECT title, content FROM notes_to_user_mapping WHERE user_name = $1 AND note_id = $2;`

func (g *GetNotesRepo) GetAllNotes(request *contracts.GetNotesRequest) (*contracts.GetNotesResponse, error) {

	response := make(contracts.GetNotesResponse, 0)
	rows, err := g.db.Query(getAllNotes, request.UserName)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch notes: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var note contracts.Note
		if err := rows.Scan(&note.Id, &note.Title, &note.Content); err != nil {
			return nil, fmt.Errorf("failed to scan note: %v", err)
		}
		response = append(response, &note)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating rows: %v", err)
	}

	return &response, nil
}

func (g *GetNotesRepo) GetNoteById(request *contracts.GetNotesRequest) (*contracts.GetNotesResponse, error) {

	response := make(contracts.GetNotesResponse, 0)
	rows, err := g.db.Query(getNoteById, request.UserName, request.NoteId)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch notes: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var note contracts.Note
		if err := rows.Scan(&note.Id, &note.Title, &note.Content); err != nil {
			return nil, fmt.Errorf("failed to scan note: %v", err)
		}
		response = append(response, &note)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while iterating rows: %v", err)
	}

	if len(response) == 0 {
		err = fmt.Errorf("either the note = %v doesn't exist or user = %v doesn't have access to the note", request.NoteId, request.UserName)
	}

	return &response, nil
}
