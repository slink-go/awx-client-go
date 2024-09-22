package api

import (
	"fmt"
	"github.com/slink-go/awx-client-go/awx/api/internal/data"
)

type WorkflowJobTemplatesResource struct {
	Resource
}

func NewWorkflowJobTemplatesResource(connection *Awx, path string) *WorkflowJobTemplatesResource {
	resource := new(WorkflowJobTemplatesResource)
	resource.connection = connection
	resource.path = path
	return resource
}

func (r *WorkflowJobTemplatesResource) Page(num int, size int) *WorkflowJobTemplatesGetRequest {
	request := new(WorkflowJobTemplatesGetRequest)
	request.resource = &r.Resource
	request.query = make(map[string][]string)
	request.query.Add("page", fmt.Sprintf("%v", num))
	request.query.Add("page_size", fmt.Sprintf("%v", size))
	return request
}

func (r *WorkflowJobTemplatesResource) Get() *WorkflowJobTemplatesGetRequest {
	request := new(WorkflowJobTemplatesGetRequest)
	request.resource = &r.Resource
	return request
}

func (r *WorkflowJobTemplatesResource) Id(id int) *WorkflowJobTemplateResource {
	return NewWorkflowJobTemplateResource(r.connection, fmt.Sprintf("%s/%d", r.path, id))
}

type WorkflowJobTemplatesGetRequest struct {
	Request
}

func (r *WorkflowJobTemplatesGetRequest) Filter(name string, value interface{}) *WorkflowJobTemplatesGetRequest {
	r.addFilter(name, value)
	return r
}

func (r *WorkflowJobTemplatesGetRequest) Send() (response *WorkflowJobTemplatesGetResponse, err error) {
	output := new(data.WorkflowJobTemplatesGetResponse)
	err = r.get(output)
	if err != nil {
		return
	}
	response = new(WorkflowJobTemplatesGetResponse)
	response.count = output.Count
	response.previous = output.Previous
	response.next = output.Next
	response.results = make([]*WorkflowJobTemplate, len(output.Results))
	for i := 0; i < len(output.Results); i++ {
		response.results[i] = new(WorkflowJobTemplate)
		response.results[i].id = output.Results[i].Id
		response.results[i].name = output.Results[i].Name
		if output.Results[i].Summary != nil && output.Results[i].Summary.Labels != nil {
			for _, l := range output.Results[i].Summary.Labels.Results {
				ll := Label{
					id:   l.Id,
					name: l.Name,
				}
				response.results[i].labels = append(response.results[i].labels, ll.name)
			}
		}
		//response.results[i].askLimitOnLaunch = output.Results[i].AskLimitOnLaunch
		//response.results[i].askVarsOnLaunch = output.Results[i].AskVarsOnLaunch
	}
	return
}

type WorkflowJobTemplatesGetResponse struct {
	ListGetResponse
	results []*WorkflowJobTemplate
}

func (r *WorkflowJobTemplatesGetResponse) Results() []*WorkflowJobTemplate {
	return r.results
}
func (r *WorkflowJobTemplatesGetResponse) HasNext() bool {
	return r.next != ""
}
