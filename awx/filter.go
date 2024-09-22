package awx

import "github.com/deckarep/golang-set/v2"

type TemplateFilter interface {
	Apply(template Template) bool
}

// region - ANY label filter

type anyLabelFilter struct {
	interestingLabels []string
}

func (lf *anyLabelFilter) Apply(template Template) bool {
	return template.Labels().ContainsAny(lf.interestingLabels...)
}
func NewAnyLabelFilter(values ...string) TemplateFilter {
	return &anyLabelFilter{
		interestingLabels: values,
	}
}

// endregion
// region - ALL label filter

type allLabelFilter struct {
	interestingLabels []string
}

func (lf *allLabelFilter) Apply(template Template) bool {
	return template.Labels().Contains(lf.interestingLabels...)
}
func NewAllLabelFilter(values ...string) TemplateFilter {
	return &allLabelFilter{
		interestingLabels: values,
	}
}

// endregion
// region - negated template id

type negatedTemplateIdFilter struct {
	filteredIds mapset.Set[int]
}

func (lf *negatedTemplateIdFilter) Apply(template Template) bool {
	return !lf.filteredIds.Contains(template.Id())
}
func NewNegatedTemplateIdFilter(values ...int) TemplateFilter {
	return &negatedTemplateIdFilter{
		filteredIds: mapset.NewSet[int](values...),
	}
}

// endregion
