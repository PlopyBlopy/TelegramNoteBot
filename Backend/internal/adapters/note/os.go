package note

import (
	"bytes"
	"io"
	"os"
)

func WriteAt(b []byte, f *os.File) (int64, int) {
	size := Lenf(f)
	off := GetInsert(f)

	buf := bytes.NewBuffer(make([]byte, 0, len(b)+2))
	if size > 10 {
		buf.WriteByte(',')
	}
	buf.Write(b)
	buf.WriteByte(']')

	n, _ := f.WriteAt(buf.Bytes(), off)

	return off + 1, n - 2
}

func GetInsert(f *os.File) (off int64) {
	size := Lenf(f)

	pos, _ := f.Seek(-1, io.SeekEnd)

	for pos >= 0 {
		var b [1]byte
		_, err := f.Read(b[:])
		if err != nil {
			return -1
		}

		if b[0] == ']' && size < 10 {
			return pos
		} else if b[0] == '}' {
			return pos + 1
		}

		// -1 от текущей позиции, -1 от только что прочитанного
		newPos, err := f.Seek(-2, io.SeekCurrent)
		if err != nil {
			return -1
		}
		pos = newPos
	}

	return -1
}

func Lenf(f *os.File) int64 {
	stat, _ := f.Stat()
	return stat.Size()
}
