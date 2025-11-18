package note

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Note struct {
	Id          int
	Title       string
	Description string
}

type NoteCard struct {
	Note        Note
	Completed   bool
	ThemeId     int
	TagsId      []int
	NoteColorId int
	CreatedAt   time.Time
}

type NoteManager struct {
	metadataManager  IMetadataManager
	noteIndexManager INoteIndexManager
	indexManager     IIndexManager
}

func NewNoteManager(mm IMetadataManager, ni INoteIndexManager, im IIndexManager) (*NoteManager, error) {
	nm := &NoteManager{
		metadataManager:  mm,
		noteIndexManager: ni,
		indexManager:     im,
	}

	err := nm.existOrCreate()
	if err != nil {
		return nil, fmt.Errorf("failed new note: %w", err)
	}

	nm.metadataManager = mm

	return nm, nil
}

// Проверка есть ли note file если нет создать пустым
func (nm NoteManager) existOrCreate() error {
	isentry := false

	files, _ := os.ReadDir(nm.metadataManager.NotePath())

	noteFileName := nm.metadataManager.NoteFileName()
	for _, file := range files {
		if !file.IsDir() && file.Name() == noteFileName {
			isentry = true
			break
		}
	}

	p := filepath.Join(nm.metadataManager.BasePath(), nm.metadataManager.NotePath(), nm.metadataManager.NoteFileName())

	empty := []interface{}{}

	if !isentry {
		os.Create(p)
		b, _ := json.Marshal(empty)
		f, _ := os.OpenFile(p, os.O_RDWR, 0666)
		n, err := f.Write(b)
		if n == 0 {

		}
		if err != nil {

		}
	}

	return nil
}

func (nm *NoteManager) AddNote(title, description string, themeId, noteColorId int, tagIds ...int) error {
	noteId := nm.metadataManager.GetNoteId()

	note := Note{
		Id:          noteId,
		Title:       title,
		Description: description,
	}

	notePath := filepath.Join(nm.metadataManager.BasePath(), nm.metadataManager.NotePath(), nm.metadataManager.NoteFileName())
	noteFile, _ := os.OpenFile(notePath, os.O_RDWR, 0666)
	defer noteFile.Close()
	b, _ := json.Marshal(note)
	off, size := WriteAt(b, noteFile)

	err := nm.noteIndexManager.AddNoteIndex(noteId, themeId, noteColorId, size, off, tagIds...)
	if err != nil {
		// nm.RemoveLastNote()
		return err
	}

	if err := nm.indexManager.AddNote(note); err != nil {

	}

	return nil
}

func (nm *NoteManager) RemoveLastNote() {

}
