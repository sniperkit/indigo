package proto

import (
	"encoding/json"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/mapping"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/mosuka/indigo/bulk"
	"reflect"
)

var (
	typeRegistry = make(map[string]reflect.Type)
)

func init() {
	typeRegistry["mapping.IndexMappingImpl"] = reflect.TypeOf(mapping.IndexMappingImpl{})
	typeRegistry["bleve.IndexStat"] = reflect.TypeOf(bleve.IndexStat{})
	typeRegistry["interface {}"] = reflect.TypeOf((map[string]interface{})(nil))
	typeRegistry["bulk.Request"] = reflect.TypeOf(bulk.Resource{})
	typeRegistry["bleve.SearchRequest"] = reflect.TypeOf(bleve.SearchRequest{})
	typeRegistry["bleve.SearchResult"] = reflect.TypeOf(bleve.SearchResult{})
}

func MarshalAny(instance interface{}) (any.Any, error) {
	var message any.Any

	value, err := json.Marshal(instance)
	if err != nil {
		return message, err
	}

	message.TypeUrl = reflect.TypeOf(instance).Elem().String()
	message.Value = value

	return message, nil
}

func UnmarshalAny(message *any.Any) (interface{}, error) {
	typeUrl := message.TypeUrl
	value := message.Value

	instance := reflect.New(typeRegistry[typeUrl]).Interface()

	err := json.Unmarshal(value, instance)
	if err != nil {
		return nil, err
	}

	return instance, nil
}
