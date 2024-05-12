package variables

import (
	"fmt"
	"io"
	"log/slog"
	"regexp"
)

type OpenVariable interface {
	Tag() string
	Name() string
	Seal() (SealedVariable, error)
}

type SealedVariable interface {
	Tag() string
	Name() string
	Value() string
}

func FindVariables(content []byte, in io.Reader) []OpenVariable {
	slog.Debug("start finding variables", "content", len(content))

	formula := `\$\{\{(.+?):(.+?)\}\}`
	regex := regexp.MustCompile(formula)

	// Using map, because I don't want to write logic to prevent duplicates
	variables := make(map[string]OpenVariable)
	for _, group := range regex.FindAllSubmatch(content, -1) {
		if len(group) >= 3 {
			tag := string(group[1])
			name := string(group[2])

			switch tag {
			case TEXT_TAG:
				slog.Debug("found variable", "tag", TEXT_TAG, "name", name, "content", len(content))
				variables[tag+name] = newTextVariable(name, in)
			case BASH_TAG:
				slog.Debug("found variable", "tag", BASH_TAG, "name", name, "content", len(content))
				variables[tag+name] = newBashVariable(name)
			}
		}
	}

	slog.Debug("done finding variables", "found", len(variables))
	slog.Debug("start converting variables map to slice", "size", len(variables))

	result := make([]OpenVariable, 0)
	for _, variable := range variables {
		result = append(result, variable)
	}

	slog.Debug("done converting variables map to slice", "size", len(result))

	return result
}

func InjectVariables(content []byte, variables []SealedVariable) []byte {
	slog.Debug("start injecting variables", "content", len(content))

	result := content

	for _, variable := range variables {
		formula := fmt.Sprintf(`\$\{\{%s:%s\}\}`, variable.Tag(), variable.Name())
		regex := regexp.MustCompile(formula)
		result = regex.ReplaceAll(result, []byte(variable.Value()))

		slog.Debug("injected variable", "tag", variable.Tag(), "name", variable.Name(), "content", len(result))
	}

	slog.Debug("done injecting variables", "content", len(result))

	return result
}
