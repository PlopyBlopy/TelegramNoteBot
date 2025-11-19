package note

import (
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

type CreateNote struct {
	Title       string
	Description string
	ThemeId     int
	TagIds      []int
	NoteColorId int
}
