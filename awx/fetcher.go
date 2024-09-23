package awx

import (
	"github.com/deckarep/golang-set/v2"
	"github.com/slink-go/awx-client-go/awx/api"
	"time"
)

const RequestPageSize = 25

type TemplateFetcher struct {
	awxClient       *api.Awx
	requestDelay    time.Duration
	templateFilters []TemplateFilter
}

// region - public API

//
// Basic Auth, Token Auth, Cookie Auth						<- AuthenticationSource / AuthenticationMethod
// AuthenticationN 	A13N: 	local, delegated, ... 			<- AuthenticationProvider
// AuthorizatioN	A11N:	local (?), SpiceDB, permify 	<- AuthorizationProvider
//

func NewTemplateFetcher(awxClient *api.Awx, requestDelay time.Duration, templateFilters ...TemplateFilter) *TemplateFetcher {
	return &TemplateFetcher{
		awxClient:       awxClient,
		requestDelay:    requestDelay,
		templateFilters: templateFilters,
	}
}
func (f *TemplateFetcher) Fetch() ([]Template, error) {

	var templates []Template

	jt, err := f.listJobTemplates()
	if err != nil {
		return nil, err
	}
	for _, t := range jt {
		// TODO: worker pool to get template variables
		v, err := f.getJobTemplateSurveySpec(t.id)
		if err != nil {
			return nil, err
		}
		t.variables = v
		templates = append(templates, *t)
		time.Sleep(f.requestDelay)
	}

	wt, err := f.listWorkflowTemplates()
	if err != nil {
		return nil, err
	}
	for _, t := range wt {
		// TODO: worker pool to get template variables
		v, err := f.getWorkflowTemplateSurveySpec(t.id)
		if err != nil {
			return nil, err
		}
		t.variables = v
		templates = append(templates, *t)
		time.Sleep(f.requestDelay)
	}

	return templates, nil

}

// region - private methods

func (f *TemplateFetcher) listJobTemplates() ([]*Template, error) {
	var result []*Template
	var page = 1
	var resource = f.awxClient.JobTemplates()
	for {
		response, err := resource.Page(page, RequestPageSize).Send()
		if err != nil {
			return nil, err
		}
		for _, p := range response.Results() {
			t := Template{
				id:          p.Id(),
				kind:        JobTemplateKind,
				name:        p.Name(),
				description: p.Description(),
				labels:      mapset.NewSet[string](p.Labels()...),
			}
			// apply all filters (AND-logic)
			apply := true
			for _, f := range f.templateFilters {
				apply = apply && f.Apply(t)
			}
			if apply {
				result = append(result, &t)
			}
		}
		if !response.HasNext() {
			break
		}
		page++
		time.Sleep(f.requestDelay)
	}
	return result, nil
}
func (f *TemplateFetcher) listWorkflowTemplates() ([]*Template, error) {
	var result []*Template
	var page = 1
	var resource = f.awxClient.WorkflowJobTemplates()
	for {
		response, err := resource.Page(page, RequestPageSize).Send()
		if err != nil {
			return nil, err
		}
		for _, p := range response.Results() {
			t := Template{
				id:          p.Id(),
				kind:        WorkflowTemplateKind,
				name:        p.Name(),
				description: p.Description(),
				labels:      mapset.NewSet[string](p.Labels()...),
			}
			// apply all filters (AND-logic)
			apply := true
			for _, f := range f.templateFilters {
				apply = apply && f.Apply(t)
			}
			if apply {
				result = append(result, &t)
			}
		}
		if !response.HasNext() {
			break
		}
		page++
		time.Sleep(f.requestDelay)
	}
	return result, nil
}
func (f *TemplateFetcher) getJobTemplateSurveySpec(templateId int) ([]Variable, error) {
	var resource = f.awxClient.JobTemplateSurveySpec().Get(templateId)
	res, err := resource.Send()
	if err != nil {
		return nil, err
	}
	var result []Variable
	for _, v := range res.Result().Variables() {
		vv := Variable{
			kind:                parseVariableKind(v.Kind()),
			questionName:        v.Question(),
			questionDescription: v.Description(),
			variable:            v.Name(),
			choices:             v.Choices(),
			min:                 v.Min(),
			max:                 v.Max(),
			required:            v.Required(),
			defaultValue:        v.Default(),
		}
		result = append(result, vv)
	}
	return result, nil
}
func (f *TemplateFetcher) getWorkflowTemplateSurveySpec(templateId int) ([]Variable, error) {
	var resource = f.awxClient.WorkflowJobTemplateSurveySpec().Get(templateId)
	res, err := resource.Send()
	if err != nil {
		return nil, err
	}
	var result []Variable
	for _, v := range res.Result().Variables() {
		vv := Variable{
			kind:                parseVariableKind(v.Kind()),
			questionName:        v.Question(),
			questionDescription: v.Description(),
			variable:            v.Name(),
			choices:             v.Choices(),
			min:                 v.Min(),
			max:                 v.Max(),
			required:            v.Required(),
			defaultValue:        v.Default(),
		}
		result = append(result, vv)
	}
	return result, nil
}

// endregion
