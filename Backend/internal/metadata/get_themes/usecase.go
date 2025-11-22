package getthemes

import "github.com/PlopyBlopy/notebot/pkg/note"

type IMetadataService interface {
	GetThemes() ([]note.Theme, error)
}

func NewUsecase(metadataService IMetadataService) func() (output, error) {
	return func() (output, error) {

		tags, err := metadataService.GetThemes()
		if err != nil {
			return output{}, err
		}

		if len(tags) == 0 {
			return output{}, nil
		}

		output := output{
			Themes: make([]theme, 0, len(tags)),
		}
		for _, t := range tags {
			output.Themes = append(output.Themes, theme{
				Id:    t.Id,
				Title: t.Title,
			})
		}

		return output, nil
	}
}
