package gettagcolor

import "github.com/PlopyBlopy/notebot/pkg/note"

type IMetadataService interface {
	GetTagColors() ([]note.Color, error)
}

func NewUsecase(metadataService IMetadataService) func() (output, error) {
	return func() (output, error) {
		colors, err := metadataService.GetTagColors()
		if err != nil {
			return output{}, err
		}

		if len(colors) == 0 {
			return output{}, nil
		}

		output := output{
			Colors: make([]color, 0, len(colors)),
		}
		for _, c := range colors {
			output.Colors = append(output.Colors, color{
				Id:       c.Id,
				Name:     c.Name,
				Variable: c.Variable,
			})
		}

		return output, nil
	}
}
