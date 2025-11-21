package gettags

import (
	"github.com/PlopyBlopy/notebot/pkg/note"
)

type IMetadataService interface {
	GetTags() ([]note.Tag, error)
}

func NewUsecase(metadataService IMetadataService) func() (output, error) {
	return func() (output, error) {
		tags, err := metadataService.GetTags()
		if err != nil {
			return output{}, err
		}

		if len(tags) == 0 {
			return output{}, nil
		}

		output := output{
			Tags: make([]tag, 0, len(tags)),
		}
		for _, t := range tags {
			output.Tags = append(output.Tags, tag{
				Id:      t.Id,
				Title:   t.Title,
				ColorId: t.ColorId,
			})
		}

		return output, nil
	}
}
