package data

type WorkflowJobTemplate struct {
	Id          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Summary     *Summary `json:"summary_fields,omitempty"`
	//AskLimitOnLaunch bool   `json:"ask_limit_on_launch,omitempty"`
	//AskVarsOnLaunch  bool   `json:"ask_variables_on_launch,omitempty"`
}

type WorkflowJobTemplateGetResponse struct {
	WorkflowJobTemplate
}
