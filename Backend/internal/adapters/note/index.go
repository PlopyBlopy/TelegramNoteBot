package note

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

type OffSize struct {
	Id   int
	Off  int64
	Size int
}

type Index struct {
	NoteIndexes      []NoteIndex
	CompletedNotes   []Note
	UnCompletedNotes []Note
	DeletedNotes     []Note
	OffSize          []OffSize
	ms               MetadataPathProvider
}

type MetadataPathProvider interface {
	BasePath() string
	IndexPath() string
	NotePath() string
	NoteFileName() string
}

func NewIndex(ms MetadataPathProvider) *Index {
	in := &Index{
		NoteIndexes:      []NoteIndex{},
		CompletedNotes:   make([]Note, 0),
		UnCompletedNotes: make([]Note, 0),
		DeletedNotes:     make([]Note, 0),
		OffSize:          make([]OffSize, 0),
		ms:               ms,
	}

	in.scan()

	return in
}

func (in *Index) scan() error {

	in.scanNoteIndex()

	scans := []func() error{
		in.scanNote,
		in.scanOffSize,
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

func (in *Index) scanNote() error {
	p := filepath.Join(in.ms.BasePath(), in.ms.NotePath(), in.ms.NoteFileName())
	b, _ := os.ReadFile(p)
	n := make([]Note, 0)
	json.Unmarshal(b, &n)

	for i := 0; i < len(n); i++ {

		if in.NoteIndexes[i].Id != n[i].Id {
			return errors.New("failed scanNote: NoteIndexes.Id not equal note.Id")
		}

		switch in.NoteIndexes[i].Status {
		case 0:
			in.UnCompletedNotes = append(in.UnCompletedNotes, n[i])
		case 1:
			in.CompletedNotes = append(in.CompletedNotes, n[i])
		case 2:
			in.DeletedNotes = append(in.DeletedNotes, n[i])
		}
	}
	return nil
}
func (in *Index) scanNoteIndex() {
	p := filepath.Join(in.ms.BasePath(), in.ms.IndexPath(), noteIndexFileName)
	b, _ := os.ReadFile(p)
	ni := make([]NoteIndex, 0)
	json.Unmarshal(b, &ni)

	in.NoteIndexes = ni
}

func (in *Index) scanOffSize() error {
	offSize := make([]OffSize, 0, len(in.NoteIndexes))

	for _, ni := range in.NoteIndexes {
		offSize = append(offSize, OffSize{
			Id:   ni.Id,
			Off:  ni.Off,
			Size: ni.Size,
		})
	}

	in.OffSize = offSize

	return nil
}

// return cursor, error
func (in *Index) GetCompletedNotes(cursor, limit int) ([]Note, int, error) {
	start, end := in.getCursorIndex(cursor, limit, len(in.CompletedNotes))
	if start == -1 || end == -1 {
		return nil, -1, nil
	}
	n := in.CompletedNotes[start:end]

	return n, end, nil
}

func (in *Index) GetUncompletedNotes(cursor, limit int) ([]Note, int, error) {
	start, end := in.getCursorIndex(cursor, limit, len(in.UnCompletedNotes))
	if start == -1 || end == -1 {
		return nil, -1, nil
	}
	n := in.UnCompletedNotes[start:end]

	return n, end, nil
}

func (in *Index) GetDeletedNotes(cursor, limit int) ([]Note, int, error) {
	start, end := in.getCursorIndex(cursor, limit, len(in.DeletedNotes))
	if start == -1 || end == -1 {
		return nil, -1, nil
	}
	n := in.DeletedNotes[start:end]

	return n, end, nil
}

func (in *Index) getCursorIndex(cursor, limit, notesLim int) (start, end int) {
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
