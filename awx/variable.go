package awx

type VariableKind int

const (
	TextVariable VariableKind = iota
	PasswordVariable
	IntVariable
	FloatVariable
	MultipleChoiceVariable
	MultipleSelectVariable
)

func (vk VariableKind) String() string {
	switch vk {
	case TextVariable:
		return "text"
	case PasswordVariable:
		return "password"
	case IntVariable:
		return "integer"
	case FloatVariable:
		return "float"
	case MultipleChoiceVariable:
		return "multiplechoice"
	case MultipleSelectVariable:
		return "multiselect"
	default:
		return "unknown"
	}
}
func parseVariableKind(input string) VariableKind {
	switch input {
	case "password":
		return PasswordVariable
	case "integer":
		return IntVariable
	case "float":
		return FloatVariable
	case "multiplechoice":
		return MultipleChoiceVariable
	case "multiselect":
		return MultipleSelectVariable
	default:
		return TextVariable
	}
}

type Variable struct {
	kind                VariableKind
	questionName        string
	questionDescription string
	variable            string
	choices             interface{}
	min                 interface{}
	max                 interface{}
	required            bool
	defaultValue        string
}

func (t *Variable) Kind() VariableKind {
	return t.kind
}
func (t *Variable) Question() string {
	return t.questionName
}
func (t *Variable) Description() string {
	return t.questionDescription
}
func (t *Variable) Name() string {
	return t.variable
}
func (t *Variable) Choices() any {
	return t.choices
}
func (t *Variable) Min() any {
	return t.min
}
func (t *Variable) Max() any {
	return t.min
}
func (t *Variable) Required() bool {
	return t.required
}
func (t *Variable) Default() string {
	return t.defaultValue
}
