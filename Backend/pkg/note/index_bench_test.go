package note

// import (
// 	"testing"

// 	"github.com/PlopyBlopy/notebot/config"
// 	"github.com/PlopyBlopy/notebot/internal/adapters/note"
// 	"github.com/rs/zerolog/log"
// )

// func BenchmarkNewIndex(b *testing.B) {
// 	b.Run("default NewIndex", func(b *testing.B) {
// 		c, err := config.InitConfig()
// 		if err != nil {
// 			log.Fatal().Err(err).Msg(err.Error())
// 		}
// 		// mc := c.Metadata.MetadataConfig
// 		ms := NewMetadataService(note.MetadataConfig(c.Metadata))
// 		b.ResetTimer()
// 		for i := 0; i < b.N; i++ {
// 			NewIndex(ms)
// 		}
// 	})
// }

// func BenchmarkScanOffSize(b *testing.B) {
// 	b.Run("setting values outside the benchmark", func(b *testing.B) {
// 		c, err := config.InitConfig()
// 		if err != nil {
// 			log.Fatal().Err(err).Msg(err.Error())
// 		}
// 		mc := c.Metadata.MetadataConfig
// 		ms := NewMetadataService(mc)

// 		in := &Index{
// 			NoteIndexes:      []NoteIndex{},
// 			CompletedNotes:   make([]Note, 0),
// 			UnCompletedNotes: make([]Note, 0),
// 			DeletedNotes:     make([]Note, 0),
// 			OffSize:          make([]OffSize, 0),
// 			ms:               ms,
// 		}

// 		in.scanNoteIndex()

// 		b.ResetTimer()

// 		for i := 0; i < b.N; i++ {

// 			in.OffSize = in.OffSize[:0]

// 			in.scanOffSize()
// 		}
// 	})

// 	b.Run("setting values in the benchmark", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			c, err := config.InitConfig()
// 			if err != nil {
// 				log.Fatal().Err(err).Msg(err.Error())
// 			}
// 			mc := c.Metadata.MetadataConfig
// 			ms := NewMetadataService(mc)

// 			in := &Index{
// 				NoteIndexes:      []NoteIndex{},
// 				CompletedNotes:   make([]Note, 0),
// 				UnCompletedNotes: make([]Note, 0),
// 				DeletedNotes:     make([]Note, 0),
// 				OffSize:          make([]OffSize, 0),
// 				ms:               ms,
// 			}

// 			in.scanNoteIndex()

// 			in.OffSize = in.OffSize[:0]

// 			in.scanOffSize()
// 		}
// 	})
// }

// func BenchmarkGenNote(b *testing.B) {
// 	b.Run("GetUncompletedNotes", func(b *testing.B) {
// 		c, err := config.InitConfig()
// 		if err != nil {
// 			log.Fatal().Err(err).Msg(err.Error())
// 		}
// 		mc := c.Metadata.MetadataConfig
// 		ms := NewMetadataService(mc)
// 		in := NewIndex(ms)

// 		b.ResetTimer()
// 		for i := 0; i < b.N; i++ {
// 			in.GetUncompletedNotes(0, 3)
// 		}
// 	})
// }
//
