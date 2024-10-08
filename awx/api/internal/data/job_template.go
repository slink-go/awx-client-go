/*
Copyright (c) 2018 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file contains the data structures used for sending and receiving job templates.

package data

type JobTemplate struct {
	Id          int      `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Summary     *Summary `json:"summary_fields,omitempty"`
	//AskLimitOnLaunch bool   `json:"ask_limit_on_launch,omitempty"`
	//AskVarsOnLaunch  bool   `json:"ask_variables_on_launch,omitempty"`
}

type JobTemplateGetResponse struct {
	JobTemplate
}
