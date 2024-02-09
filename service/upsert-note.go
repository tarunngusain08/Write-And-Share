package service

import (
	"Write-And-Share/contracts"
	"Write-And-Share/repo"
	"fmt"
)

type UpsertNoteService struct {
	repo *repo.UpsertNoteRepo
}

func NewUpsertNoteService(notesRepo *repo.UpsertNoteRepo) *UpsertNoteService {
	return &UpsertNoteService{notesRepo}
}

func (g *UpsertNoteService) UpsertNote(request *contracts.UpsertNoteRequest) error {
	switch request.ToUpdate {
	case true:
		return g.repo.UpdateNote(request)
	default:
		id, err := g.repo.CreateNote(request)
		if err != nil {
			return err
		}
		err = g.repo.AddtoNoteToUserMapping(id)
		retries := 1
		for err != nil && retries < 4 {
			fmt.Println("Retrying...")
			retries++
			err = g.repo.AddtoNoteToUserMapping(id)
		}
		return err
	}
}
