package awx

import "github.com/deckarep/golang-set/v2"

type TemplateKind int

const (
	JobTemplateKind TemplateKind = iota
	WorkflowTemplateKind
)

type Template struct {
	id          int
	kind        TemplateKind
	name        string
	description string
	labels      mapset.Set[string]
	variables   []Variable
}

func (t *Template) Id() int {
	return t.id
}
func (t *Template) Kind() TemplateKind {
	return t.kind
}
func (t *Template) Name() string {
	return t.name
}
func (t *Template) Description() string {
	return t.description
}
func (t *Template) Labels() mapset.Set[string] {
	return t.labels
}
func (t *Template) Variables() []Variable {
	return t.variables
}
