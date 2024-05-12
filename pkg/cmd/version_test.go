package cmd

import (
	"bytes"
	"io"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ExecuteVersion(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	t.Run("returns correct version", func(t *testing.T) {
		b := bytes.NewBufferString("")

		date := "MY-DATE"
		commit := "MY-COMMIT"
		version := "MY-VERSION"
		cmd := NewRoot(version, commit, date)

		cmd.SetArgs([]string{"version"})
		cmd.SetOut(b)
		cmd.Execute()

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, string(out), version)
	})

	t.Run("returns correct commit", func(t *testing.T) {
		b := bytes.NewBufferString("")

		date := "MY-DATE"
		commit := "MY-COMMIT"
		version := "MY-VERSION"
		cmd := NewRoot(version, commit, date)

		cmd.SetArgs([]string{"version"})
		cmd.SetOut(b)
		cmd.Execute()

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, string(out), commit)
	})

	t.Run("returns correct date", func(t *testing.T) {
		b := bytes.NewBufferString("")

		date := "MY-DATE"
		commit := "MY-COMMIT"
		version := "MY-VERSION"
		cmd := NewRoot(version, commit, date)

		cmd.SetArgs([]string{"version"})
		cmd.SetOut(b)
		cmd.Execute()

		out, err := io.ReadAll(b)
		if err != nil {
			t.Fatal(err)
		}

		assert.Contains(t, string(out), date)
	})
}
