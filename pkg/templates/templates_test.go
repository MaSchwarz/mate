package templates

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateTemplateFile(t testing.TB, content string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), "unit-test.md")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Error(err)
	}

	return path
}

func generateBashFile(t testing.TB, script string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), "unit-test.sh")
	if err := os.WriteFile(path, []byte(script), 0644); err != nil {
		t.Error(err)
	}

	return path
}

func Test_Load(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	t.Run("returns open template with correnct content", func(t *testing.T) {
		content := "HELLO ${{TEXT:ASDF}}"
		path := generateTemplateFile(t, content)

		template, err := Load(path, os.Stdin)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, []byte(content), template.Content())
	})

	t.Run("returns open template with correnct amount of variables", func(t *testing.T) {
		content := "HELLO ${{TEXT:ASDF}}"
		path := generateTemplateFile(t, content)

		template, err := Load(path, os.Stdin)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, len(template.Variables()), 1)
	})

	t.Run("returns open template with correnct variable tag", func(t *testing.T) {
		content := "HELLO ${{TEXT:ASDF}}"
		path := generateTemplateFile(t, content)

		template, err := Load(path, os.Stdin)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, template.Variables()[0].Tag(), "TEXT")
	})

	t.Run("returns open template with correnct variable name", func(t *testing.T) {
		content := "HELLO ${{TEXT:ASDF}}"
		path := generateTemplateFile(t, content)

		template, err := Load(path, os.Stdin)
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, template.Variables()[0].Name(), "ASDF")
	})

	t.Run("returns error if path does not exist", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "unit-test.md")

		_, err := Load(path, os.Stdin)
		assert.NotNil(t, err)
	})
}

func Test_OpenTemplate_Seal(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	t.Run("returns sealed template with correnct content", func(t *testing.T) {
		scriptPath := generateBashFile(t, "echo WORLD")
		content := fmt.Sprintf("HELLO ${{BASH:%s}}", scriptPath)
		templatePath := generateTemplateFile(t, content)

		open, err := Load(templatePath, os.Stdin)
		if err != nil {
			t.Error(err)
		}

		sealed, err := open.Seal()
		if err != nil {
			t.Error(err)
		}

		expected := []byte("HELLO WORLD")
		assert.Equal(t, expected, sealed.Content())
	})

	t.Run("returns sealed template with correnct number of variables", func(t *testing.T) {
		scriptPath := generateBashFile(t, "echo WORLD")
		content := fmt.Sprintf("HELLO ${{BASH:%s}}", scriptPath)
		templatePath := generateTemplateFile(t, content)

		open, err := Load(templatePath, os.Stdin)
		if err != nil {
			t.Error(err)
		}

		sealed, err := open.Seal()
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, 1, len(sealed.Variables()))
	})

	t.Run("returns error if variable fails", func(t *testing.T) {
		scriptPath := generateBashFile(t, "exit 1")
		content := fmt.Sprintf("HELLO ${{BASH:%s}}", scriptPath)
		templatePath := generateTemplateFile(t, content)

		open, err := Load(templatePath, os.Stdin)
		if err != nil {
			t.Error(err)
		}

		_, err = open.Seal()
		assert.NotNil(t, err)
	})
}

func Test_SealedTemplateWrite(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	t.Run("writes to the corrent path", func(t *testing.T) {
		scriptPath := generateBashFile(t, "echo WORLD")
		content := fmt.Sprintf("HELLO ${{BASH:%s}}", scriptPath)
		templatePath := generateTemplateFile(t, content)

		open, err := Load(templatePath, os.Stdin)
		if err != nil {
			t.Error(err)
		}

		sealed, err := open.Seal()
		if err != nil {
			t.Error(err)
		}

		outputPath := filepath.Join(t.TempDir(), "unit-test.out")
		if err = sealed.Write(outputPath); err != nil {
			t.Error(err)
		}

		// Check if file exists
		if _, err = os.Stat(outputPath); err != nil {
			t.Error(err)
		}
	})

	t.Run("writes the correct content", func(t *testing.T) {
		scriptPath := generateBashFile(t, "echo WORLD")
		content := fmt.Sprintf("HELLO ${{BASH:%s}}", scriptPath)
		templatePath := generateTemplateFile(t, content)

		open, err := Load(templatePath, os.Stdin)
		if err != nil {
			t.Error(err)
		}

		sealed, err := open.Seal()
		if err != nil {
			t.Error(err)
		}

		outputPath := filepath.Join(t.TempDir(), "unit-test.out")
		if err = sealed.Write(outputPath); err != nil {
			t.Error(err)
		}

		outputContent, err := os.ReadFile(outputPath)
		if err = sealed.Write(outputPath); err != nil {
			t.Error(err)
		}

		expected := []byte("HELLO WORLD")
		assert.Equal(t, expected, outputContent)
	})

	t.Run("returns error if path is invalid", func(t *testing.T) {
		scriptPath := generateBashFile(t, "echo WORLD")
		content := fmt.Sprintf("HELLO ${{BASH:%s}}", scriptPath)
		templatePath := generateTemplateFile(t, content)

		open, err := Load(templatePath, os.Stdin)
		if err != nil {
			t.Error(err)
		}

		sealed, err := open.Seal()
		if err != nil {
			t.Error(err)
		}

		outputPath := "invalid/path/that/does/not/exist"
		err = sealed.Write(outputPath)

		assert.NotNil(t, err)
	})
}
