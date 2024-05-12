package templates

import (
	"io"
	"log/slog"
	"mate/pkg/variables"
	"os"
)

type OpenTemplate interface {
	Content() []byte
	Variables() []variables.OpenVariable
	Seal() (SealedTemplate, error)
}

type SealedTemplate interface {
	Content() []byte
	Variables() []variables.SealedVariable
	Write(path string) error
}

type openTemplate struct {
	defaultOpenTemplate
}

type sealedTemplate struct {
	defaultSealedTemplate
}

func Load(path string, in io.Reader) (OpenTemplate, error) {
	slog.Debug("start loading template file", "path", path)

	content, err := os.ReadFile(path)
	if err != nil {
		return openTemplate{}, err
	}

	slog.Debug("done loading template file", "path", path)
	slog.Debug("start searching variables in template", "path", path)

	vars := variables.FindVariables(content, in)

	slog.Debug("done searching variables in template", "path", path, "size", len(vars))

	return openTemplate{
		defaultOpenTemplate{
			content:   content,
			variables: vars,
		},
	}, nil
}

func (t openTemplate) Seal() (SealedTemplate, error) {
	slog.Debug("start sealing variables", "content", len(t.content), "variables", len(t.variables))

	vars := make([]variables.SealedVariable, 0)
	for _, v := range t.variables {
		v, err := v.Seal()
		if err != nil {
			return sealedTemplate{}, err
		}

		vars = append(vars, v)
	}

	slog.Debug("done sealing variables", "content", len(t.content), "variables", len(t.variables))
	slog.Debug("start injecting variables in template content", "content", len(t.content), "variables", len(t.variables))

	content := variables.InjectVariables(t.content, vars)

	slog.Debug("done injecting variables in template content", "content", len(t.content), "variables", len(t.variables))

	return sealedTemplate{
		defaultSealedTemplate{
			content:   content,
			variables: vars,
		},
	}, nil
}

func (t sealedTemplate) Write(path string) error {
	slog.Debug("start writing template to fs", "destination", path)

	if err := os.WriteFile(path, t.content, 0644); err != nil {
		return err
	}

	slog.Debug("done writing template to fs", "destination", path)

	return nil
}
