package template

import (
	_ "embed"
	"fmt"

	"github.com/goccy/go-json"
)

type Template struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

var (
	//go:embed template.json
	content []byte
	//go:embed body.html
	body string
)

func Load() (Template, error) {
	var template Template

	if err := json.Unmarshal(content, &template); err != nil {
		return Template{}, err
	}

	if template.Subject == "" {
		return Template{}, fmt.Errorf("missing subject")
	}

	template.Body = body

	return template, nil
}
