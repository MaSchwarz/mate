package variables

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DefaultOpenVariable(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	t.Run("returns corrent tag", func(t *testing.T) {
		tag := "ASDF"
		variable := defaultOpenVariable{tag: tag, name: "IDK"}

		assert.Equal(t, variable.Tag(), tag)
	})

	t.Run("returns corrent name", func(t *testing.T) {
		name := "ASDF"
		variable := defaultOpenVariable{tag: "IDK", name: name}

		assert.Equal(t, variable.Name(), name)
	})
}

func Test_DefaultSealedVariable(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	t.Run("returns corrent tag", func(t *testing.T) {
		tag := "ASDF"
		variable := defaultSealedVariable{tag: tag, name: "IDK", value: "IDC"}

		assert.Equal(t, variable.Tag(), tag)
	})

	t.Run("returns corrent name", func(t *testing.T) {
		name := "ASDF"
		variable := defaultSealedVariable{tag: "IDK", name: name, value: "IDC"}

		assert.Equal(t, variable.Name(), name)
	})

	t.Run("returns corrent value", func(t *testing.T) {
		value := "ASDF"
		variable := defaultSealedVariable{tag: "IDK", name: "IDC", value: value}

		assert.Equal(t, variable.Value(), value)
	})
}
