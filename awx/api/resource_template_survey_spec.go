package api

type TemplateSurveySpecGetResponse struct {
	result *TemplateSurveySpec
}

func (r *TemplateSurveySpecGetResponse) Result() *TemplateSurveySpec {
	return r.result
}
