package note

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

type Index struct {
	OffSize          []OffSize
	NoteIndexes      []NoteIndex
	CompletedNotes   []Note
	UnCompletedNotes []Note
}

type OffSize struct {
	Off  int64
	Id   int
	Size int
}

type IMetadataPathProvider interface {
	BasePath() string
	IndexPath() string
	NotePath() string
	NoteFileName() string
	NoteIndexFileName() string
}

type IndexManager struct {
	i            Index
	metadataPath IMetadataPathProvider
}

type IIndexManager interface {
	AddNote(note Note) error
	AddNoteIndex(noteIndex NoteIndex) error
}

func NewIndexManager(metadataPath IMetadataPathProvider) (*IndexManager, error) {
	im := &IndexManager{
		i: Index{
			NoteIndexes:      []NoteIndex{},
			OffSize:          []OffSize{},
			CompletedNotes:   []Note{},
			UnCompletedNotes: []Note{},
		},
		metadataPath: metadataPath,
	}

	return im, nil
}

func (im *IndexManager) Scan() error {

	im.scanNoteIndex()

	scans := []func() error{
		im.scanNote,
		im.scanOffSize,
	}

	doneChan := make(chan bool, len(scans))
	errChan := make(chan error, len(scans))

	for _, scan := range scans {
		go func(f func() error) {
			if err := f(); err != nil {
				errChan <- err
			}
			doneChan <- true
		}(scan)
	}

	for i := 0; i < len(scans); i++ {
		select {
		case <-doneChan:
		case err := <-errChan:
			return err
		case <-time.After(5 * time.Second):
			return errors.New("scan time end")
		}
	}

	return nil
}

func (im *IndexManager) scanNoteIndex() {
	p := filepath.Join(im.metadataPath.BasePath(), im.metadataPath.IndexPath(), im.metadataPath.NoteIndexFileName())
	b, _ := os.ReadFile(p)
	ni := make([]NoteIndex, 0)
	json.Unmarshal(b, &ni)

	im.i.NoteIndexes = ni
}

func (im *IndexManager) scanNote() error {
	p := filepath.Join(im.metadataPath.BasePath(), im.metadataPath.NotePath(), im.metadataPath.NoteFileName())
	b, _ := os.ReadFile(p)
	n := make([]Note, 0)
	json.Unmarshal(b, &n)

	for i := 0; i < len(n); i++ {

		if im.i.NoteIndexes[i].Id != n[i].Id {
			return errors.New("failed scanNote: NoteIndexes.Id not equal note.Id")
		}

		switch im.i.NoteIndexes[i].Completed {
		case false:
			im.i.UnCompletedNotes = append(im.i.UnCompletedNotes, n[i])
		case true:
			im.i.CompletedNotes = append(im.i.CompletedNotes, n[i])
		}
	}
	return nil
}

func (im *IndexManager) scanOffSize() error {
	offSize := make([]OffSize, 0, len(im.i.NoteIndexes))

	for _, ni := range im.i.NoteIndexes {
		offSize = append(offSize, OffSize{
			Id:   ni.Id,
			Off:  ni.Off,
			Size: ni.Size,
		})
	}

	im.i.OffSize = offSize

	return nil
}

func (im *IndexManager) AddNote(note Note) error {
	im.i.CompletedNotes = append(im.i.CompletedNotes, note)

	return nil
}

func (im *IndexManager) AddNoteIndex(noteIndex NoteIndex) error {
	im.i.NoteIndexes = append(im.i.NoteIndexes, noteIndex)
	im.i.OffSize = append(im.i.OffSize, OffSize{
		Id:   noteIndex.Id,
		Off:  noteIndex.Off,
		Size: noteIndex.Size,
	})

	return nil
}

// return cursor, error
func (im *IndexManager) GetCompletedNotes(cursor, limit int) ([]Note, int, error) {
	start, end := im.getCursorIndex(cursor, limit, len(im.i.CompletedNotes))
	if start == -1 || end == -1 {
		return nil, -1, nil
	}
	n := im.i.CompletedNotes[start:end]

	return n, end, nil
}

func (im *IndexManager) GetUncompletedNotes(cursor, limit int) ([]Note, int, error) {
	start, end := im.getCursorIndex(cursor, limit, len(im.i.UnCompletedNotes))
	if start == -1 || end == -1 {
		return nil, -1, nil
	}
	n := im.i.UnCompletedNotes[start:end]

	return n, end, nil
}

// func (im *IndexManager) GetDeletedNotes(cursor, limit int) ([]Note, int, error) {
// 	start, end := in.getCursorIndex(cursor, limit, len(in.DeletedNotes))
// 	if start == -1 || end == -1 {
// 		return nil, -1, nil
// 	}
// 	n := im.DeletedNotes[start:end]

// 	return n, end, nil
// }

func (im *IndexManager) getCursorIndex(cursor, limit, notesLim int) (start, end int) {
	cl := cursor + 1 + limit

	if cl > notesLim {
		if notesLim > cursor {
			return cursor, notesLim
		}
	} else {
		return cursor, cursor + limit
	}

	return -1, -1
}
