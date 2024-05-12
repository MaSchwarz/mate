package variables

import (
	"bytes"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTextVariable(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	t.Run("returns open variable with correct tag", func(t *testing.T) {
		name := "ASDF"
		result := newTextVariable(name, os.Stdin)

		assert.Equal(t, "TEXT", result.Tag())
	})

	t.Run("returns open variable with correct name", func(t *testing.T) {
		name := "ASDF"
		result := newTextVariable(name, os.Stdin)

		assert.Equal(t, name, result.Name())
	})
}

func TestOpenTextVariableSeal(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	t.Run("returns sealed variable with correct tag", func(t *testing.T) {
		in := bytes.NewBuffer([]byte{})
		_, err := in.Write([]byte("WORLD\n"))
		if err != nil {
			t.Fatal(err)
		}

		open := newTextVariable("HELLO", in)
		sealed, err := open.Seal()
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "TEXT", sealed.Tag())
	})

	t.Run("returns sealed variable with correct name", func(t *testing.T) {
		in := bytes.NewBuffer([]byte{})
		_, err := in.Write([]byte("WORLD\n"))
		if err != nil {
			t.Fatal(err)
		}

		open := newTextVariable("HELLO", in)
		sealed, err := open.Seal()
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "HELLO", sealed.Name())
	})

	t.Run("returns sealed variable with correct value", func(t *testing.T) {
		in := bytes.NewBuffer([]byte{})
		_, err := in.Write([]byte("WORLD\n"))
		if err != nil {
			t.Fatal(err)
		}

		open := newTextVariable("HELLO", in)
		sealed, err := open.Seal()
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, "WORLD", sealed.Value())
	})
}
