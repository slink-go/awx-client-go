package data

type WorkflowJobTemplatesGetResponse struct {
	ListGetResponse

	Results []*WorkflowJobTemplate `json:"results,omitempty"`
}
