package variables

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"strings"
)

const TEXT_TAG = "TEXT"

type openTextVariable struct {
	defaultOpenVariable
	in io.Reader
}

func newTextVariable(name string, in io.Reader) OpenVariable {
	return openTextVariable{
		defaultOpenVariable: defaultOpenVariable{
			tag:  TEXT_TAG,
			name: name,
		},
		in: in,
	}
}

func (v openTextVariable) Seal() (SealedVariable, error) {
	slog.Debug("start sealing variable", "tag", v.Tag(), "name", v.Name())

	for {
		fmt.Printf("\n%s: ", v.Name())
		scanner := bufio.NewScanner(v.in)

		scanner.Scan()
		if err := scanner.Err(); err != nil {
			return defaultSealedVariable{}, err
		}

		input := strings.TrimSpace(scanner.Text())
		if len(input) > 0 {
			slog.Debug("done sealing variable", "tag", v.Tag(), "name", v.Name())

			return defaultSealedVariable{
				tag:   v.tag,
				name:  v.name,
				value: input,
			}, nil
		}
	}
}
