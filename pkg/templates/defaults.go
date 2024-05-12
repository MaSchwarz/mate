package templates

import "mate/pkg/variables"

type defaultOpenTemplate struct {
	content   []byte
	variables []variables.OpenVariable
}

func (t defaultOpenTemplate) Content() []byte {
	return t.content
}

func (t defaultOpenTemplate) Variables() []variables.OpenVariable {
	return t.variables
}

type defaultSealedTemplate struct {
	content   []byte
	variables []variables.SealedVariable
}

func (t defaultSealedTemplate) Content() []byte {
	return t.content
}

func (t defaultSealedTemplate) Variables() []variables.SealedVariable {
	return t.variables
}
