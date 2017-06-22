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

import (
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
)

type GetDocumentResponse struct {
	Id     string                  `json:"id,omitempty"`
	Fields *map[string]interface{} `json:"fields,omitempty"`
}

type GetIndexResponse struct {
	Path         string                    `json:"path,omitempty"`
	IndexMapping *mapping.IndexMappingImpl `json:"index_mapping,omitempty"`
	IndexType    string                    `json:"index_type,omitempty"`
	Kvstore      string                    `json:"kvstore,omitempty"`
	Kvconfig     interface{}               `json:"kvconfig,omitempty"`
}

type SearchResponse struct {
	SearchResult *bleve.SearchResult `json:"search_result,omitempty"`
}