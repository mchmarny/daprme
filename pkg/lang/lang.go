package lang

import (
	"context"

	"github.com/dapr-templates/daprme/pkg/lang/golang"
	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

const (
	// LangGo represents go lang
	LangGo = "go"
)

// Configurable provides lang specific project info functionality
type Configurable interface {
	InitializeProject(ctx context.Context, usr, app string) error
	GetProjectConfig() *model.Project
}

// GetLangs lists supported languages
func GetLangs() []string {
	return []string{
		LangGo,
	}
}

// MakeConfigurable returns Configurable for specific lang
func MakeConfigurable(lang string) (Configurable, error) {
	switch lang {
	case LangGo:
		return &golang.Golang{}, nil
	default:
		return nil, errors.Errorf("invalid lang: %s", lang)
	}
}
