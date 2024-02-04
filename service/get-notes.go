package service

import (
	"Write-And-Share/contracts"
	"Write-And-Share/repo"
)

type GetNotesService struct {
	repo *repo.GetNotesRepo
}

func NewGetNotesRepo(notesRepo *repo.GetNotesRepo) *GetNotesService {
	return &GetNotesService{notesRepo}
}

func (g *GetNotesService) GetNotes(request *contracts.GetNotesRequest) (*contracts.GetNotesResponse, error) {
	switch request.NoteId {
	case nil:
		return g.repo.GetAllNotes(request)
	default:
		return g.repo.GetNoteById(request)
	}
}
