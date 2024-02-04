package contracts

type LoginRequest struct {
	UserName string
	Password string
}

type GetNotesResponse []*Note

type Note struct {
	Id      int
	Title   string
	Content string
}

type GetNotesRequest struct {
	UserName string
	NoteId   *string
}
