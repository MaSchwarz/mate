package variables

import (
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateBashScript(t testing.TB, script string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), "unit-test.sh")
	if err := os.WriteFile(path, []byte(script), 0644); err != nil {
		t.Error(err)
	}

	return path
}

func Test_NewBashVariable(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	t.Run("returns open variable with corrent tag", func(t *testing.T) {
		name := "ASDF"
		result := newBashVariable(name)

		assert.Equal(t, "BASH", result.Tag())
	})

	t.Run("returns open variable with corrent name", func(t *testing.T) {
		name := "ASDF"
		result := newBashVariable(name)

		assert.Equal(t, name, result.Name())
	})
}

func Test_OpenBashVariableSeal(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	t.Run("returns sealed variable with corrent tag", func(t *testing.T) {
		script := generateBashScript(t, "echo HELLO")
		variable := newBashVariable(script)

		result, err := variable.Seal()
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "BASH", result.Tag())
	})

	t.Run("returns sealed variable with corrent name", func(t *testing.T) {
		script := generateBashScript(t, "echo HELLO")
		variable := newBashVariable(script)

		result, err := variable.Seal()
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, script, result.Name())
	})

	t.Run("returns sealed variable with corrent value", func(t *testing.T) {
		script := generateBashScript(t, "echo HELLO")
		variable := newBashVariable(script)

		result, err := variable.Seal()
		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, "HELLO", result.Value())
	})

	t.Run("returns error if script fails", func(t *testing.T) {
		script := generateBashScript(t, "exit 1")
		variable := newBashVariable(script)

		_, err := variable.Seal()
		assert.NotNil(t, err)
	})

	t.Run("returns error if script does not exist", func(t *testing.T) {
		script := "/this/path/does/not/exist.sh"
		variable := newBashVariable(script)

		_, err := variable.Seal()
		assert.NotNil(t, err)
	})
}
