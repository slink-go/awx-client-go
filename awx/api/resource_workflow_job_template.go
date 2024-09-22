package api

import (
	"github.com/slink-go/awx-client-go/awx/api/internal/data"
)

type WorkflowJobTemplateResource struct {
	Resource
}

func NewWorkflowJobTemplateResource(connection *Awx, path string) *WorkflowJobTemplateResource {
	resource := new(WorkflowJobTemplateResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *WorkflowJobTemplateResource) Get() *WorkflowJobTemplateGetRequest {
	request := new(WorkflowJobTemplateGetRequest)
	request.resource = &r.Resource
	return request
}

//func (r *WorkflowJobTemplateResource) Launch() *WorkflowJobTemplateLaunchResource {
//	return NewJobTemplateLaunchResource(r.connection, r.path+"/launch")
//}

type WorkflowJobTemplateGetRequest struct {
	Request
}

func (r *WorkflowJobTemplateGetRequest) Send() (response *WorkflowJobTemplateGetResponse, err error) {
	output := new(data.WorkflowJobTemplateGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(WorkflowJobTemplateGetResponse)
	response.result = new(WorkflowJobTemplate)
	response.result.id = output.Id
	response.result.name = output.Name
	if output.Summary != nil && output.Summary.Labels != nil {
		for _, l := range output.Summary.Labels.Results {
			ll := Label{
				id:   l.Id,
				name: l.Name,
			}
			response.result.labels = append(response.result.labels, ll.name)
		}
	}

	//response.result.askLimitOnLaunch = output.AskLimitOnLaunch
	//response.result.askVarsOnLaunch = output.AskVarsOnLaunch

	return
}

type WorkflowJobTemplateGetResponse struct {
	result *WorkflowJobTemplate
}

func (r *WorkflowJobTemplateGetResponse) Result() *WorkflowJobTemplate {
	return r.result
}
