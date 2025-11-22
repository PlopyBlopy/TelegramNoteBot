package addtag

type IMetadataService interface {
	AddTag(string, int) error
}

func NewUsecase(metadataService IMetadataService) func(input) error {
	return func(input input) error {
		err := metadataService.AddTag(input.Title, input.ColorId)
		if err != nil {
			return err
		}

		return nil
	}
}
