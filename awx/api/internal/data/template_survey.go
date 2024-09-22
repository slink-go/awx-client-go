package data

type TemplateSurveySpec struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Spec        []struct {
		Type                string      `json:"type"`
		QuestionName        string      `json:"question_name"`
		QuestionDescription string      `json:"question_description"`
		Variable            string      `json:"variable"`
		Choices             interface{} `json:"choices"`
		Min                 interface{} `json:"min"`
		Max                 interface{} `json:"max"`
		Required            bool        `json:"required"`
		Default             string      `json:"default"`
	} `json:"spec"`
}

type TemplateSurveySpecGetResponse struct {
	TemplateSurveySpec
}
