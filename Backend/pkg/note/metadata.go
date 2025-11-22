package note

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
	CurrentId      int     `json:"current_id"` // last note id for autoincrement
	Themes         []Theme `json:"themes"`
	Tags           []Tag   `json:"tags"`
	TagColors      []Color `json:"tag_colors"`
	NoteCardColors []Color `json:"note_card_colors"`
}

type MetadataManager struct {
	m              *Metadata
	metadataConfig *MetadataConfig
}

type IMetadataManager interface {
	GetNoteId() int
	BasePath() string
	IndexPath() string
	NotePath() string
	NoteFileName() string
	NoteIndexFileName() string

	AddTheme(title string) error
	AddTag(title string, colorId int) error
	AddTagColor(name, variable string) error

	GetTags() ([]Tag, error)
	GetTagIds() ([]int, error)
	GetThemes() ([]Theme, error)
	GetThemeIds() ([]int, error)
}

func NewMetadataManager(c *MetadataConfig) (*MetadataManager, error) {
	mm := &MetadataManager{
		m: &Metadata{
			CurrentId:      0,
			Themes:         []Theme{},
			Tags:           []Tag{},
			TagColors:      []Color{},
			NoteCardColors: []Color{},
		},
		metadataConfig: c,
	}

	err := mm.getMetadataOrCreate()
	if err != nil {
		return nil, fmt.Errorf("failed new metadata: %w", err)
	}

	mm.metadataConfig = c

	return mm, nil
}

/*
Проверка наличия файла metadata.json и чтение данных в структуру Metadata
В случае если его нету, создастся новый файл metadata.json
*/
func (mm MetadataManager) getMetadataOrCreate() error {
	isentry := false

	files, _ := os.ReadDir(mm.metadataConfig.Basepath)

	if len(files) != 0 {
		for _, file := range files {
			if !file.IsDir() && file.Name() == mm.metadataConfig.MetadataFilename {
				isentry = true
				break
			}
		}
	}

	p := filepath.Join(mm.metadataConfig.Basepath, mm.metadataConfig.MetadataFilename)

	if !isentry {
		b, _ := json.Marshal(mm.m)
		os.Create(p)
		f, _ := os.OpenFile(p, os.O_RDWR, 0666)
		n, err := f.Write(b)
		if n == 0 {

		}
		if err != nil {

		}
	} else {
		b, _ := os.ReadFile(p)
		_ = json.Unmarshal(b, &mm.m)

	}

	return nil
}

func (mm *MetadataManager) AddTheme(title string) error {
	metadata := Metadata{}

	p := filepath.Join(mm.BasePath(), mm.MetadataFileName())
	b, _ := os.ReadFile(p)
	json.Unmarshal(b, &metadata)

	for i := 0; i < len(metadata.Themes); i++ {
		if metadata.Themes[i].Title == title {
			return fmt.Errorf("failed append theme, an themes with a similar title already exists, title: %s", title)

		}
	}

	theme := Theme{
		Id:    len(metadata.Themes),
		Title: title,
	}

	// append to file slice
	metadata.Themes = append(metadata.Themes, theme)
	// append to virtual slice
	mm.m.Themes = append(mm.m.Themes, theme)

	b, _ = json.Marshal(metadata)
	os.WriteFile(p, b, 0666)

	return nil
}

// append tags to file and virtual
func (mm *MetadataManager) AddTag(title string, colorId int) error {
	metadata := Metadata{}

	p := filepath.Join(mm.BasePath(), mm.MetadataFileName())
	b, _ := os.ReadFile(p)
	json.Unmarshal(b, &metadata)

	for i := 0; i < len(metadata.Tags); i++ {
		if metadata.Tags[i].Title == title {
			return fmt.Errorf("failed append tags, an tag with a similar title already exists, title: %s", title)
		}
	}

	tag := Tag{
		Id:      len(metadata.Tags),
		Title:   title,
		ColorId: colorId,
	}

	// append to file slice
	metadata.Tags = append(metadata.Tags, tag)
	// append to virtual slice
	mm.m.Tags = append(mm.m.Tags, tag)

	b, _ = json.Marshal(metadata)
	os.WriteFile(p, b, 0666)

	return nil
}

func (mm *MetadataManager) AddTagColor(name, variable string) error {
	metadata := Metadata{}

	p := filepath.Join(mm.BasePath(), mm.MetadataFileName())
	b, _ := os.ReadFile(p)
	json.Unmarshal(b, &metadata)

	for i := 0; i < len(metadata.TagColors); i++ {
		if strings.EqualFold(metadata.TagColors[i].Name, name) {
			return fmt.Errorf("failed append tag color, an tag colors with a similar name already exists, name: %s", name)
		} else if strings.EqualFold(metadata.TagColors[i].Variable, variable) {
			return fmt.Errorf("failed append tag color, an tag colors with a similar variable already exists, variable: %s", variable)
		}
	}

	tagColor := Color{
		Id:       len(metadata.TagColors),
		Name:     name,
		Variable: variable,
	}

	// append to file slice
	metadata.TagColors = append(metadata.TagColors, tagColor)
	// append to virtual slice
	mm.m.TagColors = append(mm.m.TagColors, tagColor)

	b, _ = json.Marshal(metadata)
	os.WriteFile(p, b, 0666)

	return nil
}

// need synchronyzed CurrentId Increment, Decrement, Value
func (mm *MetadataManager) GetNoteId() int {
	// var mu sync.Mutex
	// mu.Lock()
	// defer mu.Unlock()

	p := filepath.Join(mm.BasePath(), mm.MetadataFileName())
	metadate := Metadata{}
	b, _ := os.ReadFile(p)
	json.Unmarshal(b, &metadate)

	id := metadate.CurrentId
	metadate.CurrentId += 1
	mm.m.CurrentId += 1

	b, _ = json.Marshal(metadate)
	os.WriteFile(p, b, 0666)

	return id
}

func (mm *MetadataManager) GetTags() ([]Tag, error) {
	return mm.m.Tags, nil
}
func (mm *MetadataManager) GetTagIds() ([]int, error) {
	tagIds := make([]int, 0, len(mm.m.Tags))

	for _, v := range mm.m.Tags {
		tagIds = append(tagIds, v.Id)
	}

	return tagIds, nil
}
func (mm *MetadataManager) GetThemes() ([]Theme, error) {
	return mm.m.Themes, nil
}
func (mm *MetadataManager) GetThemeIds() ([]int, error) {
	themeIds := make([]int, 0, len(mm.m.Themes))

	for _, v := range mm.m.Themes {
		themeIds = append(themeIds, v.Id)
	}

	return themeIds, nil
}

// append theme to file and virtual

func (mm *MetadataManager) BasePath() string          { return mm.metadataConfig.Basepath }
func (mm *MetadataManager) IndexPath() string         { return mm.metadataConfig.Indexpath }
func (mm *MetadataManager) NotePath() string          { return mm.metadataConfig.Notepath }
func (mm *MetadataManager) MetadataFileName() string  { return mm.metadataConfig.MetadataFilename }
func (mm *MetadataManager) NoteIndexFileName() string { return mm.metadataConfig.NoteIndexFilename }
func (mm *MetadataManager) NoteFileName() string      { return mm.metadataConfig.NoteFilename }
