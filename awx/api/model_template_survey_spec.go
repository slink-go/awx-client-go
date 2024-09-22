package api

type TemplateSurveySpec struct {
	name        string
	description string
	variables   []TemplateVariable
}

func (t *TemplateSurveySpec) Name() string {
	return t.name
}
func (t *TemplateSurveySpec) Description() string {
	return t.description
}
func (t *TemplateSurveySpec) Variables() []TemplateVariable {
	return t.variables
}
