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

// TODO: goroutine для записей в файлы NoteId, NoteIndex, Theme, Tags
// Вариант для NoteId, Theme, Tags -> обьединить в одну операцию и открывать и записывать 1 раз а не 3
// для записи в файл для NoteIndex и в Metadata: 2 goroutine, синхронизация с AddNote через 2 канала для Metadata и NoteIndex - возможно и 1 канал
// Возможне не надо ждать результат записи - риск отката не будет
// Вывод определить долго ли запись в файл
// !!!!!!!!! передавать канал в анонимные функции и если ошибка добавлять значение в errChan
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
	// syncNote(note) - добавить в срез Note
	// syncNoteIndex(noteIndex) - добавить в срез NoteIndex
	// syncIndex - пересчитать индексы для Note, NoteIndex, и по сути другие индексы, то есть нужно универсальная система для обновления индекса в фоне но с синхронизацией
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
