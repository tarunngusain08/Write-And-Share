package service

import (
	"Write-And-Share/contracts"
	"Write-And-Share/repo"
)

type ShareNoteService struct {
	repo *repo.ShareNoteRepo
}

func NewShareNoteService(notesRepo *repo.ShareNoteRepo) *ShareNoteService {
	return &ShareNoteService{notesRepo}
}

func (g *ShareNoteService) ShareNote(request *contracts.ShareNoteRequest) error {
	return g.repo.ShareNote(request)
}
