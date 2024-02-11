package service

import (
	"Write-And-Share/contracts"
	"Write-And-Share/repo"
	"strings"
)

type SearchNoteService struct {
	repo *repo.SearchNoteRepo
}

func NewSearchNoteService(notesRepo *repo.SearchNoteRepo) *SearchNoteService {
	return &SearchNoteService{notesRepo}
}

func (g *SearchNoteService) SearchNote(request *contracts.SearchNoteRequest) (*contracts.NotesList, error) {

	request.Keywords = strings.Join(strings.Split(request.Keywords, " "), " | ")
	return g.repo.SearchNote(request)
}
