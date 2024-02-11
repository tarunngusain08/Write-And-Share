package repo

import (
	"Write-And-Share/contracts"
	"database/sql"
	"strings"
)

type SearchNoteRepo struct {
	db *sql.DB
}

func NewSearchNoteRepo(db *sql.DB) *SearchNoteRepo {
	return &SearchNoteRepo{db}
}

const searchNotes = `SELECT id, title FROM notes
WHERE to_tsvector('english', title) @@ to_tsquery('english', $1)
ORDER BY timestamp DESC
LIMIT $2 OFFSET $3;`

func (l *SearchNoteRepo) SearchNote(request *contracts.SearchNoteRequest) (*contracts.NotesList, error) {

	page := *request.PageNum
	pageSize := *request.PageSize
	offset := (page - 1) * pageSize

	request.Keywords = strings.Join(strings.Split(request.Keywords, " "), " | ")

	res, err := l.db.Query(searchNotes, request.Keywords, pageSize, offset)
	if err != nil {
		return nil, err
	}

	response := make(contracts.NotesList, 0)
	for res.Next() {
		var note *contracts.Note
		if err := res.Scan(&note.Id, &note.Title); err != nil {
			return nil, err
		}
		response = append(response, note)
	}
	return &response, nil
}
