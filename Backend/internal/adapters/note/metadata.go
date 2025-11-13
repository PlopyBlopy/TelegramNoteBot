package note

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type MetadataConfig struct {
	Basepath          string `env:"MD_BASE_PATH"`
	Indexpath         string `env:"MD_INDEX_PATH"`
	Notepath          string `env:"MD_NOTE_PATH"`
	MetadataFilename  string `env:"MD_METADATA_FILE_NAME"`
	NoteIndexFilename string `env:"MD_NOTE_INDEX_FILE_NAME"`
	NoteFilename      string `env:"MD_NOTE_FILE_NAME"`
}

type Metadata struct {
	CurrentId  int      `json:"current_id"` // last note id for autoincrement
	IsHaveNote bool     `json:"is_have_note"`
	Themes     []string `json:"themes"`
	Tags       []string `json:"tags"`
	mc         MetadataConfig
}

func NewMetadataService(c MetadataConfig) *Metadata {
	m := &Metadata{mc: c}
	files, _ := os.ReadDir(m.BasePath())

	isentry := false

	for _, file := range files {
		if !file.IsDir() && file.Name() == m.MetadataFileName() {
			isentry = true
			break
		}
	}

	notePath := filepath.Join(m.BasePath(), m.NotePath())
	files, _ = os.ReadDir(notePath)
	if len(files) != 0 {
		m.IsHaveNote = true
	}

	p := filepath.Join(m.BasePath(), m.MetadataFileName())

	if !isentry {
		m.newFile(p)
	} else {
		b, _ := os.ReadFile(p)
		_ = json.Unmarshal(b, &m)

	}

	return m
}

// TODO: из .env если файла нету или его данные изменены -> замена на новые
func (m *Metadata) newFile(path string) {
	m.CurrentId = 0

	b, _ := json.Marshal(m)
	os.Create(path)
	f, _ := os.OpenFile(path, os.O_RDWR, 0666)
	n, err := f.Write(b)
	if n == 0 {

	}
	if err != nil {

	}
}

// need synchronyzed CurrentId Increment, Decrement, Value
func (m *Metadata) GetNoteId() int {
	// var mu sync.Mutex
	// mu.Lock()
	// defer mu.Unlock()

	p := filepath.Join(m.BasePath(), m.MetadataFileName())
	metadate := Metadata{}
	b, _ := os.ReadFile(p)
	json.Unmarshal(b, &metadate)

	id := metadate.CurrentId
	metadate.CurrentId += 1
	m.CurrentId += 1

	b, _ = json.Marshal(metadate)
	os.WriteFile(p, b, 0666)

	return id
}

func (m *Metadata) BasePath() string          { return m.mc.Basepath }
func (m *Metadata) IndexPath() string         { return m.mc.Indexpath }
func (m *Metadata) NotePath() string          { return m.mc.Notepath }
func (m *Metadata) MetadataFileName() string  { return m.mc.MetadataFilename }
func (m *Metadata) NoteIndexFileName() string { return m.mc.NoteIndexFilename }
func (m *Metadata) NoteFileName() string      { return m.mc.NoteFilename }
func (m *Metadata) IsHaveNoteFile() bool      { return m.IsHaveNote }

func (m *Metadata) HaveNote(isHave bool) {
	m.IsHaveNote = isHave

	p := filepath.Join(m.BasePath(), m.MetadataFileName())
	metadate := Metadata{}
	b, _ := os.ReadFile(p)
	json.Unmarshal(b, &metadate)

	metadate.IsHaveNote = isHave
	m.IsHaveNote = isHave

	b, _ = json.Marshal(metadate)
	os.WriteFile(p, b, 0666)
}
func (m *Metadata) AppendTheme(theme string) {
	p := filepath.Join(m.BasePath(), m.MetadataFileName())
	metadate := Metadata{}
	b, _ := os.ReadFile(p)
	json.Unmarshal(b, &metadate)

	for i := 0; i < len(metadate.Themes); i++ {
		if metadate.Themes[i] == theme {
			return
		}
	}

	metadate.Themes = append(metadate.Themes, theme)
	m.Themes = append(m.Themes, theme)

	b, _ = json.Marshal(metadate)
	os.WriteFile(p, b, 0666)
}

func (m *Metadata) AppendTags(tags ...string) {
	if len(tags) == 0 {
		return
	}
	p := filepath.Join(m.BasePath(), m.MetadataFileName())
	metadate := Metadata{}
	b, _ := os.ReadFile(p)
	json.Unmarshal(b, &metadate)

	appendTags := make([]int, len(tags))

	// TODO: на битовую маску
	for i := 0; i < len(metadate.Tags); i++ {
		for j := 0; j < len(tags); j++ {
			if metadate.Tags[i] == tags[j] {
				appendTags[j] = 1
				return
			}
		}
	}

	for i := 0; i < len(appendTags); i++ {
		if appendTags[i] == 0 {
			metadate.Tags = append(metadate.Tags, tags[i])
			m.Tags = append(m.Tags, tags[i])
		}
	}

	b, _ = json.Marshal(metadate)
	os.WriteFile(p, b, 0666)
}
