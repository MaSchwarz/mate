package variables

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

const BASH_TAG = "BASH"

type openBashVariable struct {
	defaultOpenVariable
}

func newBashVariable(name string) openBashVariable {
	return openBashVariable{
		defaultOpenVariable{
			tag:  BASH_TAG,
			name: name,
		},
	}
}

func (v openBashVariable) Seal() (SealedVariable, error) {
	if _, err := os.Stat(v.name); errors.Is(err, os.ErrNotExist) {
		return defaultSealedVariable{}, err
	}

	ouput, err := exec.Command("/bin/sh", v.name).Output()
	if err != nil {
		return defaultSealedVariable{}, err
	}

	value := strings.TrimSpace(string(ouput))

	return defaultSealedVariable{
		tag:   v.tag,
		name:  v.name,
		value: value,
	}, nil
}
