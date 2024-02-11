package contracts

import "time"

type LoginRequest struct {
	UserName string
	Password string
}

type NotesList []*Note

type Note struct {
	Id      int
	Title   string
	Content string
}

type GetNotesRequest struct {
	UserName string
	NoteId   *string
	PageNum  *int
	PageSize *int
}

type UpsertNoteRequest struct {
	Username string
	Content  string
	Heading  string
	Time     time.Time
	ToUpdate bool
	NoteId   string
}

type DeleteNoteRequest struct {
	Username string
	NoteId   string
}

type ShareNoteRequest struct {
	Username       string
	NoteId         string
	ShareUsernames []string
}

type SearchNoteRequest struct {
	PageNum  *int
	PageSize *int
	Keywords string
}
