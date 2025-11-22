package addtheme

type IMetadataService interface {
	AddTheme(string) error
}

func NewUsecase(metadataService IMetadataService) func(input) error {
	return func(input input) error {
		err := metadataService.AddTheme(input.Title)
		if err != nil {
			return err
		}

		return nil
	}
}
