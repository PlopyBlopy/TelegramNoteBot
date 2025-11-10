package note

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Note struct {
	Id          int
	Title       string
	Description string
}

type NoteManager struct {
	metadata MetadataService
}

type MetadataService interface {
	GetNoteId() int
	BasePath() string
	IndexPath() string
	NotePath() string
	NoteFileName() string
	NoteIndexFileName() string

	IsHaveNoteFile() bool
	HaveNote(isHave bool)
	AppendTheme(theme string)
	AppendTags(tags ...string)
}

func NewNoteService(ms MetadataService) *NoteManager {
	return &NoteManager{
		metadata: ms,
	}
}

func (nm *NoteManager) AddNote(title, description, theme string, tags ...string) {
	var err error
	if !nm.metadata.IsHaveNoteFile() {
		err = nm.newNoteFile()
		if err != nil {

		}
	}

	noteId := nm.metadata.GetNoteId()

	note := Note{
		Id:          noteId,
		Title:       title,
		Description: description,
	}

	notePath := filepath.Join(nm.metadata.BasePath(), nm.metadata.NotePath(), nm.metadata.NoteFileName())
	noteFile, _ := os.OpenFile(notePath, os.O_RDWR, 0666)
	defer noteFile.Close()
	b, _ := json.Marshal(note)
	off, size := WriteAt(b, noteFile)

	createdAt := time.Now()

	index := NoteIndex{
		Id:        noteId,
		Status:    0,
		Theme:     theme,
		Tags:      tags,
		Off:       off,
		Size:      size,
		UpdatedAt: createdAt,
		CreatedAt: createdAt,
	}

	indexPath := filepath.Join(nm.metadata.BasePath(), nm.metadata.IndexPath(), nm.metadata.NoteIndexFileName())
	indexFile, _ := os.OpenFile(indexPath, os.O_RDWR, 0666)
	defer indexFile.Close()
	b, _ = json.Marshal(index)
	WriteAt(b, indexFile)

	nm.metadata.AppendTheme(theme)
	nm.metadata.AppendTags(tags...)
}

// func (nm *NoteManager) GetNotes(limit int) (pointer int) {

// }

func (nm *NoteManager) newNoteFile() (err error) {
	notePath := filepath.Join(nm.metadata.BasePath(), nm.metadata.NotePath(), nm.metadata.NoteFileName())
	indexPath := filepath.Join(nm.metadata.BasePath(), nm.metadata.IndexPath(), nm.metadata.NoteIndexFileName())

	data := []byte("[]")
	os.WriteFile(notePath, data, 0666)
	os.WriteFile(indexPath, data, 0666)

	nm.metadata.HaveNote(true)

	return nil
}
