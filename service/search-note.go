package service

import (
	"Write-And-Share/contracts"
	"Write-And-Share/repo"
)

type SearchNoteService struct {
	repo *repo.SearchNoteRepo
}

func NewSearchNoteService(notesRepo *repo.SearchNoteRepo) *SearchNoteService {
	return &SearchNoteService{notesRepo}
}

func (g *SearchNoteService) SearchNote(request *contracts.SearchNoteRequest) (*contracts.NotesList, error) {
	return g.repo.SearchNote(request)
}
