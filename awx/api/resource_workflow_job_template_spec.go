package api

import (
	"fmt"
	"github.com/slink-go/awx-client-go/awx/api/internal/data"
)

type WorkflowJobTemplateSurveySpecResource struct {
	Resource
}

func NewWorkflowJobTemplateSurveySpecResource(connection *Awx, path string) *WorkflowJobTemplateSurveySpecResource {
	resource := new(WorkflowJobTemplateSurveySpecResource)
	resource.connection = connection
	resource.templatePath = path
	return resource
}

func (r *WorkflowJobTemplateSurveySpecResource) Get(id int) *WorkflowJobTemplateSurveySpecGetRequest {
	request := new(WorkflowJobTemplateSurveySpecGetRequest)
	request.resource = &r.Resource
	r.path = fmt.Sprintf(r.templatePath, id)
	return request
}

type WorkflowJobTemplateSurveySpecGetRequest struct {
	Request
}

func (r *WorkflowJobTemplateSurveySpecGetRequest) Send() (response *TemplateSurveySpecGetResponse, err error) {
	output := new(data.TemplateSurveySpecGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(TemplateSurveySpecGetResponse)
	response.result = new(TemplateSurveySpec)
	response.result.name = output.Name
	response.result.description = output.Description
	for _, sv := range output.Spec {
		tv := TemplateVariable{
			kind:                sv.Type,
			questionName:        sv.QuestionName,
			questionDescription: sv.QuestionDescription,
			variable:            sv.Variable,
			choices:             sv.Choices,
			min:                 sv.Min,
			max:                 sv.Max,
			required:            sv.Required,
			defaultValue:        sv.Default,
		}
		response.result.variables = append(response.result.variables, tv)
	}
	return
}
