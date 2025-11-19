package service

import (
	"fmt"

	"github.com/PlopyBlopy/notebot/internal/adapters/note"
)

type INoteManager interface {
	AddNote(title, description string, themeId, noteColorId int, tagIds ...int) error
	GetFilteredNoteCards(search string, limit, themeId int, tags ...int) ([]note.NoteCard, error)
}

type NoteService struct {
	noteManager INoteManager
}

func NewNoteService(nm INoteManager) (*NoteService, error) {
	return &NoteService{noteManager: nm}, nil
}

func (ns NoteService) AddNote(note note.CreateNote) error {
	ns.noteManager.AddNote(note.Title, note.Description, note.ThemeId, note.NoteColorId, note.TagIds...)
	return nil
}

func (ns NoteService) GetFilteredNoteCards(search string, limit, themeId int, tagIds ...int) ([]note.NoteCard, error) {
	noteCards, err := ns.noteManager.GetFilteredNoteCards(search, limit, themeId, tagIds...)
	if err != nil {
		return nil, fmt.Errorf("failed get filtered note cards")
	}
	return noteCards, nil
}
