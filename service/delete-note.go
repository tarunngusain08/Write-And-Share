package service

import (
	"Write-And-Share/contracts"
	"Write-And-Share/repo"
)

type DeleteNoteService struct {
	repo *repo.DeleteNoteRepo
}

func NewDeleteNoteService(notesRepo *repo.DeleteNoteRepo) *DeleteNoteService {
	return &DeleteNoteService{notesRepo}
}

func (g *DeleteNoteService) DeleteNote(request *contracts.DeleteNoteRequest) error {
	return g.repo.DeleteNote(request)
}
