package noteservice

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

/*

✅CreateAndWrite (name string , note Note) error
- формирует файл и добавляет первую запись

✅Write (name string , note Note) error
- добавляет в конец файла note

✅Read (name string, id uuid.UUID) (*Note, error)
- name - имя файла без json
- возвращает note по указателю
- в нутри обращается к Store для поиска файла с: name + json
❗- если файл изменился в runtime, например при форматировании - все offsets сдвинулись

❌ReadByFilter (limit int, sort string) (*[]Note, off int, error)
- чтения опредленного кол-ва файлов из ходя из переданнго limit и по какому либо переданному sort полю
- возвращать срез Note и off - что = последней ноте что была прочитана
❗- изменить NoteInf и добавить sort поля, например: isDeleted, createdAt

❌Update (name string, note Note) error
- удаляет старый note вызывая Remove
- добавляет в конец новый note
- обновляет FileOffset его offset и size

❌Remove (name string, id uuid.UUID) error
- удаляет старый note - ставится флаг isDeleted=true

*/

// Contains file names and their offset notes in a file
type Store struct {
	// key = fileName.json, val = FileOffset
	Files map[string]FileOff
}

// Offset notes in a file
type FileOff struct {
	// key = note.id: val = NoteInfo
	Offsets map[uuid.UUID]NoteInf
}

// Offset - skipping from the beginning of the file
// Size - the number of bytes written
type NoteInf struct {
	Offset int64
	Size   int
}

// Container for detecting a deleted Note
type NoteBox struct {
	Isdeleted bool `json:"is_deleted"`
	Note
}

// Note structure
type Note struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Text      string    `json:"text"`
}

const rootPath = "./src/note/" // The basic path to the note files

// Creates a new Store
// Returns a pointer to the Store
func NewStore() (*Store, error) {
	s := scanDir(rootPath)

	for k := range s.Files {
		fo, err := s.scanFile(k)
		if err != nil {
			return nil, fmt.Errorf("cannot scan file: %w", err)
		}

		s.Files[k] = fo
	}

	return s, nil
}

// Creating a file with a theme for Note and initial initialization
func (s Store) CreateFileAndWrite(name string, n Note) error {
	filePath := filepath.Join(rootPath, name+".json")

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("failed create and write file: %w", err)
	}
	defer file.Close()

	noteBox := NoteBox{
		Isdeleted: false,
		Note:      n,
	}

	data, err := json.Marshal(noteBox)
	if err != nil {
		return fmt.Errorf("failed marshal Note: %w", err)
	}

	// an opening and closing square bracket has been added, necessary for syntax .json file
	buf := bytes.NewBuffer(make([]byte, 0, len(data)+2))
	buf.WriteByte('[')
	buf.Write(data)
	buf.WriteByte(']')

	wb, err := file.Write(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed write to file: %w", err)
	}

	fileOff := FileOff{
		Offsets: map[uuid.UUID]NoteInf{n.Id: {
			Offset: 1,      //Offset = 1, this is the removed opening square bracket
			Size:   wb - 2, //Size = wb - 2, this is the opening and closing square bracket removed from the size of the written bytes
		}},
	}
	s.addFile(name, fileOff)
	return nil
}

// Adds a NotePad (Note) to the end of the json file
func (s Store) Write(name string, note Note) error {
	filePath := filepath.Join(rootPath, name+".json")

	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if err := s.CreateFileAndWrite(name, note); err != nil {
				return fmt.Errorf("failed write and create: %w", err)
			}
			return nil
		}
		return fmt.Errorf("failed write: %w", err)
	}
	defer file.Close()

	noteBox := NoteBox{
		Isdeleted: false,
		Note:      note,
	}

	data, err := json.Marshal(noteBox)
	if err != nil {
		return fmt.Errorf("failed marshal Note: %w", err)
	}

	insert := getEndInsert(file)

	buf := bytes.NewBuffer(make([]byte, 0, len(data)+2))
	buf.WriteByte(',')
	buf.Write(data)
	buf.WriteByte(']')

	n, err := file.WriteAt(buf.Bytes(), insert)
	if err != nil {
		return fmt.Errorf("failed write to file: %w", err)
	}

	noteInf := NoteInf{
		Offset: insert + 1,
		Size:   n - 2,
	}

	s.addNoteInf(name, noteInf, note.Id)

	return nil
}

// To read the NoteBox -> Note from a file
func (s Store) Read(name string, id uuid.UUID) (*Note, error) {
	noteInf := s.Files[name+".json"].Offsets[id]
	filePath := filepath.Join(rootPath, name+".json")

	file, _ := os.OpenFile(filePath, os.O_RDONLY, 0666)
	defer file.Close()

	data := make([]byte, noteInf.Size)
	if _, err := file.ReadAt(data, noteInf.Offset); err != nil {
		return nil, fmt.Errorf("failed read file: %w", err)
	}

	if len(data) > 0 && data[len(data)-1] == ',' {
		data = data[:len(data)-1]
	}
	var nc NoteBox
	json.Unmarshal(data, &nc)

	return &nc.Note, nil
}

// Scans the folder using the base path where all the note files are located.
// Adds their names to the map, creates an empty FileOff, not nil
// Returns a pointer to the Store
func scanDir(path string) *Store {
	dir, _ := os.ReadDir(path)

	s := Store{
		Files: make(map[string]FileOff, len(dir)),
	}

	for _, d := range dir {
		s.Files[d.Name()] = FileOff{
			Offsets: make(map[uuid.UUID]NoteInf),
		}
	}

	return &s
}

// Scans the file to get a FileOff and a map of its Offset data to search for all the Notes that are available at the time of launch.
// The entire file is loaded into memory.
func (s Store) scanFile(name string) (FileOff, error) {
	fileOff := FileOff{
		Offsets: make(map[uuid.UUID]NoteInf),
	}

	filePath := filepath.Join(rootPath, name)

	stat, err := os.Lstat(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fileOff, err
		}
		return fileOff, fmt.Errorf("failed describing the named file: %w", err)
	}
	file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
	if err != nil {
		return fileOff, fmt.Errorf("failed open file. %w", err)
	}
	defer file.Close()

	bytes := make([]byte, stat.Size())
	file.Read(bytes) // load into memory

	var off int64 = 0 // offset

	for i := 0; i < len(bytes); i++ {
		// The first byte of the NoteBox.
		if bytes[i] == '{' {
			off = int64(i)
		}

		// The last byte of the NoteBox
		if bytes[i] == '}' {
			size := i + 1 - int(off)
			temp := make([]byte, size)
			file.ReadAt(temp, off)

			var noteBox NoteBox
			json.Unmarshal(temp, &noteBox)

			if !noteBox.Isdeleted {
				fileOff.Offsets[noteBox.Id] = NoteInf{
					Offset: off,
					Size:   size,
				}
			}
		}
	}

	return fileOff, nil
}

// Adds a new key=file name and value=FileOff
func (s Store) addFile(name string, fo FileOff) error {
	fileName := name + ".json"

	s.Files[fileName] = fo

	return nil
}

// Adds NoteInf to the FileOff from the Store by file name
func (s Store) addNoteInf(name string, ni NoteInf, id uuid.UUID) {
	fileName := name + ".json"

	s.Files[fileName].Offsets[id] = ni
}

// Returns the place to insert the +1 character while saving the character.
func getEndInsert(file *os.File) (off int64) {
	pos, _ := file.Seek(-1, io.SeekEnd)

	for pos >= 0 {
		var b [1]byte
		_, err := file.Read(b[:])
		if err != nil {
			return -1
		}

		if b[0] == '}' {
			return pos + 1
		}

		// -1 от текущей позиции, -1 от только что прочитанного
		newPos, err := file.Seek(-2, io.SeekCurrent)
		if err != nil {
			return -1
		}
		pos = newPos
	}

	return -1
}
