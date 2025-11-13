package note

import "time"

// TODO: bit mask for tags
type NoteIndex struct {
	Id        int       `json:"id"`
	Status    int       `json:"status"` // 0 - не выполнено, 1 - выполнено, 2 - удалена
	Theme     string    `json:"theme"`  // есть всегда, default = без темы
	Tags      []string  `json:"tags"`
	Off       int64     `json:"off"`  // offset
	Size      int       `json:"size"` // writed bytes
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

const noteIndexFileName = "note_index.json"
