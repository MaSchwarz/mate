package variables

type defaultOpenVariable struct {
	tag  string
	name string
}

func (v defaultOpenVariable) Tag() string {
	return v.tag
}

func (v defaultOpenVariable) Name() string {
	return v.name
}

type defaultSealedVariable struct {
	tag   string
	name  string
	value string
}

func (v defaultSealedVariable) Tag() string {
	return v.tag
}

func (v defaultSealedVariable) Name() string {
	return v.name
}

func (v defaultSealedVariable) Value() string {
	return v.value
}
