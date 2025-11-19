package note

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Index struct {
	NoteTitles       map[string]int // title: noteId
	Themes           map[int][]int  // themeId: noteId
	Tags             map[int][]int  // tagId: noteId
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

type IndexManager struct {
	i               Index
	metadataManager IMetadataManager
}

type IIndexManager interface {
	AddNote(note Note) error
	AddNoteIndex(noteIndex NoteIndex) error

	GetCompletedNotesFilteredNoteIds(noteIds ...int) ([]Note, error)
	GetNoteIndexesFilteredNoteIds(noteIds ...int) ([]NoteIndex, error)

	GetFilteredTitleNoteIds(search string) ([]int, error)
	GetFilteredTagNoteIds(tagIds ...int) ([]int, error)
	GetFilteredThemeNoteIds(themeId int) ([]int, error)
}

func NewIndexManager(mm IMetadataManager) (*IndexManager, error) {
	im := &IndexManager{
		i: Index{
			NoteTitles:       map[string]int{},
			Themes:           map[int][]int{},
			Tags:             map[int][]int{},
			NoteIndexes:      []NoteIndex{},
			OffSize:          []OffSize{},
			CompletedNotes:   []Note{},
			UnCompletedNotes: []Note{},
		},
		metadataManager: mm,
	}

	return im, nil
}

func (im *IndexManager) Scan() error {

	im.scanNoteIndex()

	scans := []func() error{
		im.scanNote,
		im.scanOffSize,
		im.scanNoteTheme,
		im.scanNoteTag,
	}

	im.scanNoteTitle()

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
	p := filepath.Join(im.metadataManager.BasePath(), im.metadataManager.IndexPath(), im.metadataManager.NoteIndexFileName())
	b, _ := os.ReadFile(p)
	ni := []NoteIndex{}
	json.Unmarshal(b, &ni)

	im.i.NoteIndexes = ni
}

func (im *IndexManager) scanNote() error {
	p := filepath.Join(im.metadataManager.BasePath(), im.metadataManager.NotePath(), im.metadataManager.NoteFileName())
	b, _ := os.ReadFile(p)
	n := []Note{}
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

// scan only by completed notes
func (im *IndexManager) scanNoteTitle() error {
	noteTitles := make(map[string]int, len(im.i.CompletedNotes))
	completedNotes := im.i.CompletedNotes

	for _, v := range completedNotes {
		key := strings.ToLower(v.Title)
		noteTitles[key] = v.Id
	}

	im.i.NoteTitles = noteTitles

	return nil
}
func (im *IndexManager) scanNoteTheme() error {
	ThemeIds, err := im.metadataManager.GetThemeIds()
	if err != nil {

	}

	themes := make(map[int][]int, len(ThemeIds))
	noteIndexes := im.i.NoteIndexes

	for _, themeId := range ThemeIds {
		for _, noteIndex := range noteIndexes {
			themes[themeId] = append(themes[themeId], noteIndex.Id)
		}
	}

	im.i.Themes = themes

	return nil
}
func (im *IndexManager) scanNoteTag() error {
	TagsIds, err := im.metadataManager.GetTagIds()
	if err != nil {

	}

	tags := make(map[int][]int, len(TagsIds))
	noteIndexes := im.i.NoteIndexes

	for _, tagId := range TagsIds {
		for _, noteIndex := range noteIndexes {
			tags[tagId] = append(tags[tagId], noteIndex.Id)
		}
	}

	im.i.Tags = tags

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

// GetCompletedNoteIds(cursor, limit int) ([]int, int, error)
func (im *IndexManager) GetCompletedNotesFilteredNoteIds(noteIds ...int) ([]Note, error) {
	notes := make([]Note, 0, len(noteIds)) // problem
	completedNotes := im.i.CompletedNotes

	for _, n := range completedNotes {
		for _, id := range noteIds {
			if n.Id == id {
				notes = append(notes, n)
			}
		}
	}

	return notes, nil
}

func (im *IndexManager) GetUncompletedNotes(cursor, limit int) ([]Note, int, error) {
	start, end := im.getCursorIndex(cursor, limit, len(im.i.UnCompletedNotes))
	if start == -1 || end == -1 {
		return nil, -1, nil
	}
	n := im.i.UnCompletedNotes[start:end]

	return n, end, nil
}

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

func (im *IndexManager) GetNoteIndexesFilteredNoteIds(noteIds ...int) ([]NoteIndex, error) {
	notes := make([]NoteIndex, 0, len(noteIds))
	noteIndexes := im.i.NoteIndexes

	for _, n := range noteIndexes {
		for _, id := range noteIds {
			if n.Id == id {
				notes = append(notes, n)
			}
		}
	}

	return notes, nil
}

func (im *IndexManager) GetFilteredTitleNoteIds(search string) ([]int, error) {
	ids := []int{}
	titles := im.i.NoteTitles
	title := strings.ToLower(search)

	for k, v := range titles {
		if strings.Contains(k, title) {
			ids = append(ids, v)
		}
	}

	return ids, nil
}

func (im *IndexManager) GetFilteredTagNoteIds(tagIds ...int) ([]int, error) {
	tags := im.i.Tags

	// срезы что содержатся у тегов
	noteIds := [][]int{}

	for k, v := range tags {
		for _, id := range tagIds {
			if k == id {
				noteIds = append(noteIds, v)
			}
		}

	}

	// карта с k=noteId, v=количеству упоминаний
	temp := map[int]int{}

	for i := 0; i < len(noteIds); i++ {
		for _, id := range noteIds[i] {
			temp[id] = temp[id] + 1
		}
	}

	// результирующий срез с noteId, содержащиеся в искомом/искомых теге/тегах
	res := []int{}

	for k, v := range temp {
		if v == len(tagIds) {
			res = append(res, k)
		}
	}

	return res, nil
}

func (im *IndexManager) GetFilteredThemeNoteIds(themeId int) ([]int, error) {
	ids := []int{}
	themes := im.i.Themes

	if themeId < 0 {
		return ids, nil
	}

	for k, v := range themes {
		if k == themeId {
			ids = append(ids, v...)
		}
	}

	return ids, nil
}
