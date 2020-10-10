package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppMarshaling(t *testing.T) {
	app := &App{
		Meta: Meta{
			Lang:       "go",
			Main:       "main.go",
			Name:       "test",
			Port:       8080,
			Type:       AppTypeCLI,
			UsesClient: true,
		},
		Bindings: []*Component{
			&Component{
				Name: "c1",
				Type: "t1",
			},
			&Component{
				Name: "c2",
				Type: "t2",
			},
		},
		Components: []*Component{
			&Component{
				Name: "c1",
				Type: "t1",
			},
			&Component{
				Name: "c2",
				Type: "t2",
			},
		},
		PubSubs: []*PubSub{
			&PubSub{
				Name:  "p1",
				Type:  "t1",
				Topic: "t1",
			},
			&PubSub{
				Name:  "p2",
				Type:  "t2",
				Topic: "t2",
			},
		},
		Services: []*Service{
			&Service{
				Name: "s1",
			},
			&Service{
				Name: "s2",
			},
		},
	}

	b2, err := app.Marshal()
	if err != nil {
		t.FailNow()
	}

	app2, err := Unmarshal(b2)
	if err != nil {
		t.FailNow()
	}

	assert.Equal(t, app.Meta, app2.Meta)
}
