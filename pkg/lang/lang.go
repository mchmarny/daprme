package lang

import (
	"context"

	"github.com/dapr-templates/daprme/pkg/lang/golang"
	"github.com/dapr-templates/daprme/pkg/lang/nodejs"
	"github.com/dapr-templates/daprme/pkg/model"
	"github.com/pkg/errors"
)

const (
	// supported languages
	langGo   = "go"
	langNode = "node"
)

// Configurable provides lang specific project info functionality
type Configurable interface {
	InitializeProject(ctx context.Context, dir, usr, app string) error
	GetProjectConfig() *model.Project
}

// GetLangs lists supported languages
func GetLangs() []string {
	return []string{
		langGo,
		langNode,
	}
}

// MakeConfigurable returns Configurable for specific lang
func MakeConfigurable(lang string) (Configurable, error) {
	switch lang {
	case langGo:
		return &golang.Golang{}, nil
	case langNode:
		return &nodejs.NodeJs{}, nil
	default:
		return nil, errors.Errorf("invalid lang: %s", lang)
	}
}
