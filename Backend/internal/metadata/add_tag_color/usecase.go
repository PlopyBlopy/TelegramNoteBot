package addtagcolor

type IMetadataService interface {
	AddTagColor(name, variable string) error
}

func NewUsecase(metadataService IMetadataService) func(input) error {
	return func(input input) error {
		err := metadataService.AddTagColor(input.Name, input.Variable)
		if err != nil {
			return err
		}

		return nil
	}
}
