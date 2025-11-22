package addcardcolor

type IMetadataService interface {
	AddCardColor(name, variable string) error
}

func NewUsecase(metadataService IMetadataService) func(input) error {
	return func(input input) error {
		err := metadataService.AddCardColor(input.Name, input.Variable)
		if err != nil {
			return err
		}

		return nil
	}
}
