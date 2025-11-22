package service

import (
	"fmt"

	"github.com/PlopyBlopy/notebot/pkg/note"
)

type IMetadataManager interface {
	AddTag(title string, colorId int) error
	AddTheme(title string) error
	AddTagColor(name, variable string) error
	// AddCardColor(name, variable string) error

	GetTags() ([]note.Tag, error)
	GetThemes() ([]note.Theme, error)
	// GetTagColors() ([]note.Color, error)
	// GetThemeColors() ([]note.Color, error)
}

type MetadataService struct {
	metadataManager IMetadataManager
}

func NewMetadataService(mm IMetadataManager) (*MetadataService, error) {
	return &MetadataService{metadataManager: mm}, nil
}
func (ms MetadataService) AddTag(title string, colorId int) error {
	err := ms.metadataManager.AddTag(title, colorId)
	if err != nil {
		return fmt.Errorf("failed add tag: %w", err)
	}

	return nil
}

func (ms MetadataService) AddTheme(title string) error {
	err := ms.metadataManager.AddTheme(title)
	if err != nil {
		return fmt.Errorf("failed add theme: %w", err)
	}

	return nil
}

func (ms MetadataService) AddTagColor(name, variable string) error {
	err := ms.metadataManager.AddTagColor(name, variable)
	if err != nil {
		return fmt.Errorf("failed add theme: %w", err)
	}

	return nil
}

func (ms MetadataService) GetTags() ([]note.Tag, error) {
	tags, err := ms.metadataManager.GetTags()
	if err != nil {
		return nil, fmt.Errorf("failed get tags: %w", err)
	}
	return tags, nil
}
func (ms MetadataService) GetThemes() ([]note.Theme, error) {
	themes, err := ms.metadataManager.GetThemes()
	if err != nil {
		return nil, fmt.Errorf("failed get themes: %w", err)
	}
	return themes, nil
}

// func (ms MetadataService) GetTagColors() ([]note.Color, error) {
// 	tagColors, err := ms.metadataManager.GetTagColors()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed get theme colors: %w", err)
// 	}
// 	return tagColors, nil
// }
// func (ms MetadataService) GetThemeColors() ([]note.Color, error) {
// 	themeColors, err := ms.metadataManager.GetThemeColors()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed get tag colors: %w", err)
// 	}
// 	return themeColors, nil
// }
