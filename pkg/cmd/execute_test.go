package cmd

import (
	"bytes"
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

func Test_NewExecute(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	t.Run("creates a new file in the correct place with correct content", func(t *testing.T) {
		in := bytes.NewBuffer([]byte{})
		path := generateTemplateFile(t, "HELLO ${{TEXT:WORLD}}")

		cmd := NewRoot("", "", "")
		cmd.SetArgs([]string{"exec", path})
		cmd.SetIn(in)

		// Fill TEXT variable prompt
		_, err := in.Write([]byte("WORLD\n"))
		if err != nil {
			t.Fatal(err)
		}

		err = cmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		outputPath := path + ".output"
		outputContent, err := os.ReadFile(outputPath)
		if err != nil {
			t.Fatal(err)
		}

		expected := []byte("HELLO WORLD")
		assert.Equal(t, expected, outputContent)
	})

	t.Run("respects output flag", func(t *testing.T) {
		in := bytes.NewBuffer([]byte{})
		path := generateTemplateFile(t, "HELLO ${{TEXT:WORLD}}")
		outputPath := filepath.Join(t.TempDir(), "output.md")

		cmd := NewRoot("", "", "")
		cmd.SetArgs([]string{"exec", path, fmt.Sprintf("--out=%s", outputPath)})
		cmd.SetIn(in)

		// Fill TEXT variable prompt
		_, err := in.Write([]byte("WORLD\n"))
		if err != nil {
			t.Fatal(err)
		}

		err = cmd.Execute()
		if err != nil {
			t.Fatal(err)
		}

		outputContent, err := os.ReadFile(outputPath)
		if err != nil {
			t.Fatal(err)
		}

		expected := []byte("HELLO WORLD")
		assert.Equal(t, expected, outputContent)
	})

	t.Run("fails if no arg is given", func(t *testing.T) {
		cmd := NewRoot("", "", "")
		cmd.SetArgs([]string{"exec"})

		err := cmd.Execute()
		assert.NotNil(t, err)
	})

	t.Run("fails if 2 or more args are given", func(t *testing.T) {
		path1 := generateTemplateFile(t, "HELLO ${{TEXT:WORLD}}")
		path2 := generateTemplateFile(t, "HELLO ${{TEXT:MOM}}")

		cmd := NewRoot("", "", "")
		cmd.SetArgs([]string{"exec", path1, path2})

		err := cmd.Execute()
		assert.NotNil(t, err)
	})

	t.Run("fails if template file is missing", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "unit-test.md")

		cmd := NewRoot("", "", "")
		cmd.SetArgs([]string{"exec", path})

		err := cmd.Execute()
		assert.NotNil(t, err)
	})
}
