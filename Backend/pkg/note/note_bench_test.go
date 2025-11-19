package note

// import (
// 	"testing"

// 	"github.com/PlopyBlopy/notebot/config"
// 	"github.com/rs/zerolog/log"
// )

// func BenchmarkAddNote(b *testing.B) {
// 	b.Run("GetUncompletedNotes", func(b *testing.B) {
// 		c, err := config.InitConfig()
// 		if err != nil {
// 			log.Fatal().Err(err).Msg(err.Error())
// 		}
// 		mc := c.Metadata.MetadataConfig
// 		ms := NewMetadataService(mc)
// 		ns := NewNoteService(ms)

// 		b.ResetTimer()
// 		for i := 0; i < b.N; i++ {
// 			ns.AddNote("First title", "First Description", "First Theme", "someOtherTag")
// 		}
// 	})
// }
