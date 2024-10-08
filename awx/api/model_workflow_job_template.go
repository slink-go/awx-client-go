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

// This file contains the implementation of the job template type.

package api

type WorkflowJobTemplate struct {
	id          int
	name        string
	description string
	labels      []string
}

func (t *WorkflowJobTemplate) Id() int {
	return t.id
}

func (t *WorkflowJobTemplate) Name() string {
	return t.name
}
func (t *WorkflowJobTemplate) Description() string {
	return t.description
}
func (t *WorkflowJobTemplate) Labels() []string {
	return t.labels
}
