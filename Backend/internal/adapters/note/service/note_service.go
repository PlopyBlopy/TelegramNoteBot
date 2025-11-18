package service

import (
	"time"

	"github.com/PlopyBlopy/notebot/internal/adapters/note"
)

type INoteManager interface {
	AddNote(title, description string, themeId, noteColorId int, tagIds ...int) error
}

type CreateNote struct {
	Title       string
	Description string
	ThemeId     int
	TagIds      []int
	NoteColorId int
}

type NoteCard struct {
	Note        note.Note
	Completed   bool
	ThemeId     int
	TagsId      []int
	NoteColorId int
	CreatedAt   time.Time
}

type NoteService struct {
	noteManager INoteManager
}

func NewNoteService(nm INoteManager) (*NoteService, error) {
	return &NoteService{noteManager: nm}, nil
}

func (ns NoteService) AddNote(note CreateNote) error {
	ns.noteManager.AddNote(note.Title, note.Description, note.ThemeId, note.NoteColorId, note.TagIds...)
	return nil
}
