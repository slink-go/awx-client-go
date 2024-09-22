package api

type TemplateVariable struct {
	kind                string
	questionName        string
	questionDescription string
	variable            string
	choices             interface{}
	min                 interface{}
	max                 interface{}
	required            bool
	defaultValue        string
}

func (t *TemplateVariable) Kind() string {
	return t.kind
}
func (t *TemplateVariable) Question() string {
	return t.questionName
}
func (t *TemplateVariable) Description() string {
	return t.questionDescription
}
func (t *TemplateVariable) Name() string {
	return t.variable
}
func (t *TemplateVariable) Choices() any {
	return t.choices
}
func (t *TemplateVariable) Min() any {
	return t.min
}
func (t *TemplateVariable) Max() any {
	return t.min
}
func (t *TemplateVariable) Required() bool {
	return t.required
}
func (t *TemplateVariable) Default() string {
	return t.defaultValue
}
