package templates

import (
	"bytes"
	"log/slog"
	"mate/pkg/variables"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultOpenTemplate(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	t.Run("returns correct content", func(t *testing.T) {
		content := []byte("HELLO ${{BASH:A}}")
		template := defaultOpenTemplate{
			content:   content,
			variables: make([]variables.OpenVariable, 0),
		}

		assert.Equal(t, content, template.Content())
	})

	t.Run("returns correct variables", func(t *testing.T) {
		content := []byte("HELLO ${{BASH:A}}")
		vars := variables.FindVariables(content, os.Stdout)

		template := defaultOpenTemplate{
			content:   content,
			variables: vars,
		}

		assert.Equal(t, vars, template.Variables())
	})
}

func Test_DefaultSealedTemplate(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	t.Run("returns correct content", func(t *testing.T) {
		content := []byte("HELLO WORLD")
		template := defaultSealedTemplate{
			content:   content,
			variables: make([]variables.SealedVariable, 0),
		}

		assert.Equal(t, content, template.Content())
	})

	t.Run("returns correct variables", func(t *testing.T) {
		// Fill Stdin with TEXT variable prompt answer
		in := bytes.NewBuffer([]byte{})
		_, err := in.Write([]byte("WORLD\n"))
		if err != nil {
			t.Fatal(err)
		}

		content := []byte("HELLO ${{TEXT:A}}")
		vars := variables.FindVariables(content, in)

		sealedVars := make([]variables.SealedVariable, 0)
		for _, variable := range vars {
			sealed, err := variable.Seal()
			if err != nil {
				t.Fatal(err)
			}

			sealedVars = append(sealedVars, sealed)
		}

		sealedContent := variables.InjectVariables(content, sealedVars)

		template := defaultSealedTemplate{
			content:   sealedContent,
			variables: sealedVars,
		}

		assert.Equal(t, sealedVars, template.Variables())
	})
}
