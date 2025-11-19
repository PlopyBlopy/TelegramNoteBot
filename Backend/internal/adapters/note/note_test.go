package note

import (
	"fmt"
	"testing"
	"time"
)

func TestGetFilteredNotes(t *testing.T) {

	noteLen := 5
	completedNotes := make([]Note, 0, noteLen)

	for i := 0; i < noteLen; i++ {
		completedNotes = append(completedNotes, Note{
			Id:          i,
			Title:       fmt.Sprintf("Title %d", i),
			Description: "Some Description for this note",
		})
	}

	themes := map[int][]int{0: {0, 1, 2, 3, 4}, 1: {2, 3}, 2: {4}}
	tags := map[int][]int{0: {0, 1, 2}, 1: {2, 3}, 2: {4}}

	noteIndexes := make([]NoteIndex, 0, noteLen)

	noteIndexes = append(noteIndexes, NoteIndex{
		Id:        0,
		Completed: false,
		Deleted:   false,
		ThemeId:   0,
		TagIds:    []int{0},
		CreatedAt: time.Now(),
	})

	noteIndexes = append(noteIndexes, NoteIndex{
		Id:        1,
		Completed: false,
		Deleted:   false,
		ThemeId:   0,
		TagIds:    []int{0},
		CreatedAt: time.Now(),
	})

	noteIndexes = append(noteIndexes, NoteIndex{
		Id:        2,
		Completed: false,
		Deleted:   false,
		ThemeId:   1,
		TagIds:    []int{0, 1},
		CreatedAt: time.Now(),
	})

	noteIndexes = append(noteIndexes, NoteIndex{
		Id:        3,
		Completed: false,
		Deleted:   false,
		ThemeId:   1,
		TagIds:    []int{1},
		CreatedAt: time.Now(),
	})

	noteIndexes = append(noteIndexes, NoteIndex{
		Id:        4,
		Completed: false,
		Deleted:   false,
		ThemeId:   2,
		TagIds:    []int{2},
		CreatedAt: time.Now(),
	})

	indexManager := IndexManager{
		i: Index{
			NoteTitles: make(map[string]int, len(completedNotes)),
		},
	}

	noteCards := make([]NoteCard, 0, noteLen)

	for i := 0; i < noteLen; i++ {
		noteCards = append(noteCards, NoteCard{
			Note:        completedNotes[i],
			Completed:   noteIndexes[i].Completed,
			ThemeId:     noteIndexes[i].ThemeId,
			TagsId:      noteIndexes[i].TagIds,
			NoteColorId: noteIndexes[i].NoteColorId,
			CreatedAt:   noteIndexes[i].CreatedAt,
		})
	}

	indexManager.i.CompletedNotes = completedNotes
	indexManager.i.NoteIndexes = noteIndexes
	indexManager.i.Tags = tags
	indexManager.i.Themes = themes

	indexManager.scanNoteTitle()

	// limit not used now
	tests := []struct {
		name     string
		search   string
		themeId  int
		tagIds   []int
		limit    int
		expected []NoteCard
	}{
		{"empty", "", -1, []int{}, 0, []NoteCard{noteCards[4]}},
		{"theme", "", 2, []int{}, 0, []NoteCard{noteCards[4]}},
		{"theme", "", 1, []int{}, 0, []NoteCard{noteCards[2], noteCards[3]}},
		{"search, theme", "Tit", 0, []int{}, 0, []NoteCard{noteCards[0], noteCards[1], noteCards[2], noteCards[3], noteCards[4]}},
		{"search, theme", "Tit", 1, []int{}, 0, []NoteCard{noteCards[2], noteCards[3]}},
		{"search, theme", "Title 2", 1, []int{}, 0, []NoteCard{noteCards[2]}},
		{"theme, tag", "", 0, []int{0}, 0, []NoteCard{noteCards[0], noteCards[1], noteCards[2]}},
		{"theme, tags", "", 0, []int{0, 1}, 0, []NoteCard{noteCards[2]}},
		{"theme, tags", "", 1, []int{0, 1}, 0, []NoteCard{noteCards[2]}},
	}

	noteManager := NoteManager{indexManager: &indexManager}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := noteManager.GetFilteredNoteCards(tt.search, tt.limit, tt.themeId, tt.tagIds...)
			if err != nil {
				t.Errorf("failed test: %s", err)
			}

			if len(res) == 0 {
				return
			}

			resCount := 0

			for i := 0; i < len(res); i++ {
				equal := true
				card := res[i]
				expect := tt.expected[i]

				resCount++

				if card.Note != expect.Note || card.Completed != expect.Completed || card.ThemeId != expect.ThemeId || card.NoteColorId != expect.NoteColorId || card.CreatedAt != expect.CreatedAt {
					equal = false
				}

				if len(card.TagsId) != len(expect.TagsId) {
					equal = false
				} else {
					for j := range card.TagsId {
						if card.TagsId[j] != expect.TagsId[j] {
							equal = false
							break
						}
					}
				}

				if !equal {
					t.Errorf("GetFilteredNotes(%s, %d, %d, %d)=%+v; expected %+v", tt.search, tt.limit, tt.themeId, tt.tagIds, res, tt.expected[i])
				}

			}

			if resCount < len(tt.expected) {
				t.Errorf("less values passed the check than they should have passed")
			}
		})
	}
}
