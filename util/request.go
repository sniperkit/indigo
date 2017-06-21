//  Copyright (c) 2017 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

type Document struct {
	Id     string                 `json:"id,omitempty"`
	Fields map[string]interface{} `json:"fields,omitempty"`
}

type UpdateRequest struct {
	Method   string   `json:"method,omitempty"`
	Document Document `json:"document,omitempty"`
}

type BulkResource struct {
	BatchSize int32           `json:"batch_size,omitempty"`
	Requests  []UpdateRequest `json:"requests,omitempty"`
}

type PutDocumentResource struct {
	Id     string                 `json:"id,omitempty"`
	Fields map[string]interface{} `json:"fields,omitempty"`
}
