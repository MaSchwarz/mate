package variables

import (
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindVariables(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	t.Run("finds single TEXT variable", func(t *testing.T) {
		content := []byte("HELLO ${{TEXT:A}}")
		result := FindVariables(content, os.Stdin)

		expected := []OpenVariable{newTextVariable("A", os.Stdin)}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("finds single TEXT variable if multiple have same name", func(t *testing.T) {
		content := []byte("HELLO ${{TEXT:A}} ${{TEXT:A}}")
		result := FindVariables(content, os.Stdin)

		expected := []OpenVariable{newTextVariable("A", os.Stdin)}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("finds single BASH variable", func(t *testing.T) {
		content := []byte("HELLO ${{BASH:A}}")
		result := FindVariables(content, os.Stdin)

		expected := []OpenVariable{newBashVariable("A")}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("finds single BASH variable if multiple have same name", func(t *testing.T) {
		content := []byte("HELLO ${{BASH:A}} ${{BASH:A}}")
		result := FindVariables(content, os.Stdin)

		expected := []OpenVariable{newBashVariable("A")}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("finds multiple variables", func(t *testing.T) {
		content := []byte("HELLO ${{TEXT:A}} ${{BASH:B}}")
		result := FindVariables(content, os.Stdin)

		expected := []OpenVariable{
			newTextVariable("A", os.Stdin),
			newBashVariable("B"),
		}

		assert.ElementsMatch(t, expected, result)
	})

	t.Run("finds multiple variables if multiple have same name but different tags", func(t *testing.T) {
		content := []byte("HELLO ${{TEXT:A}} ${{BASH:A}}")
		result := FindVariables(content, os.Stdin)

		expected := []OpenVariable{
			newTextVariable("A", os.Stdin),
			newBashVariable("A"),
		}

		assert.ElementsMatch(t, expected, result)
	})
}

func Test_InjectVariables(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})))

	t.Run("injects TEXT variable", func(t *testing.T) {
		content := []byte("HELLO ${{TEXT:A}}")
		variables := []SealedVariable{defaultSealedVariable{tag: "TEXT", name: "A", value: "WORLD"}}
		result := InjectVariables(content, variables)

		expected := []byte("HELLO WORLD")

		assert.Equal(t, expected, result)
	})

	t.Run("injects BASH variable", func(t *testing.T) {
		content := []byte("HELLO ${{BASH:A}}")
		variables := []SealedVariable{defaultSealedVariable{tag: "BASH", name: "A", value: "WORLD"}}
		result := InjectVariables(content, variables)

		expected := []byte("HELLO WORLD")

		assert.Equal(t, expected, result)
	})

	t.Run("injects multiple variables", func(t *testing.T) {
		content := []byte("HELLO ${{TEXT:A}} ${{BASH:B}}")
		variables := []SealedVariable{
			defaultSealedVariable{tag: "TEXT", name: "A", value: "WORLD"},
			defaultSealedVariable{tag: "BASH", name: "B", value: "AGAIN"},
		}

		result := InjectVariables(content, variables)
		expected := []byte("HELLO WORLD AGAIN")

		assert.Equal(t, expected, result)
	})

	t.Run("injects nothing if variable tag is not in content", func(t *testing.T) {
		content := []byte("HELLO ${{TEXT:A}}")
		variables := []SealedVariable{defaultSealedVariable{tag: "BASH", name: "B", value: "WORLD"}}
		result := InjectVariables(content, variables)

		expected := []byte("HELLO ${{TEXT:A}}")
		assert.Equal(t, expected, result)
	})

	t.Run("injects nothing if variable name is not in content", func(t *testing.T) {
		content := []byte("HELLO ${{TEXT:B}}")
		variables := []SealedVariable{defaultSealedVariable{tag: "BASH", name: "B", value: "WORLD"}}
		result := InjectVariables(content, variables)

		expected := []byte("HELLO ${{TEXT:B}}")
		assert.Equal(t, expected, result)
	})
}
