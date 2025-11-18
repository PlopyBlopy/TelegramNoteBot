package service

type IMetadataManager interface{}

type MetadataService struct {
	metadataManager IMetadataManager
}

func NewMetadataService(mm IMetadataManager) (*MetadataService, error) {
	return &MetadataService{metadataManager: mm}, nil
}
