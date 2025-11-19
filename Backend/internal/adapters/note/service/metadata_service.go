package service

import "github.com/PlopyBlopy/notebot/internal/adapters/note"

type IMetadataManager interface{}

type MetadataService struct {
	metadataManager IMetadataManager
}

func NewMetadataService(mm IMetadataManager) (*MetadataService, error) {
	return &MetadataService{metadataManager: mm}, nil
}

func GetTags() ([]note.Tag, error) {
	return nil, nil
}
func GetThemes() ([]note.Theme, error) {
	return nil, nil
}
func GetTagsColor() ([]note.Color, error) {
	return nil, nil
}
func GetThemesColor() ([]note.Color, error) {
	return nil, nil
}
