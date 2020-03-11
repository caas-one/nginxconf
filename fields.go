package nginxconf

import (
	"errors"
	"reflect"
	"strings"
)

var (
	egTag  = "eg"
	keyTag = "key"
)

// Object object field message
type Object struct {
	Field   string      `json:"field,omitempty"`
	Key     string      `json:"key,omitempty"`
	Value   interface{} `json:"value,omitempty"`
	Type    string      `json:"type,omitempty"`
	JSONKey string      `json:"jsonKey,omitempty"`
	Example string      `json:"example,omitempty"`
}

// GetFields  get fields type eg
func GetFields(object interface{}) (interface{}, error) {
	t := reflect.TypeOf(object)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	v := reflect.ValueOf(object)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	objMsg := make([]*Object, 0)
	for index := 0; index < v.NumField(); index++ {
		jsonKey := t.Field(index).Tag.Get("json")
		if jsonKey != "" {
			if strings.Contains(jsonKey, ",") {
				jsonKey = strings.Split(jsonKey, ",")[0]
			}
		}
		key := t.Field(index).Tag.Get(keyTag)
		eg := t.Field(index).Tag.Get(egTag)
		eg = strings.TrimPrefix(strings.TrimPrefix(eg, key+" "), " ")
		objMsg = append(objMsg, &Object{
			Field:   t.Field(index).Name,
			Type:    t.Field(index).Type.String(),
			Example: eg,
			Key:     key,
			JSONKey: jsonKey,
		})
	}
	if len(objMsg) <= 0 {
		return nil, errors.New("can't get fields")
	}
	return objMsg, nil
}

// GetFieldsAndValue get object field value type and so on
func GetFieldsAndValue(object interface{}) ([]*Object, error) {
	t := reflect.TypeOf(object)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	v := reflect.ValueOf(object)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	objMsg := make([]*Object, 0)
	for index := 0; index < v.NumField(); index++ {
		jsonKey := t.Field(index).Tag.Get("json")
		if jsonKey != "" {
			if strings.Contains(jsonKey, ",") {
				jsonKey = strings.Split(jsonKey, ",")[0]
			}
		}
		tt := t.Field(index).Type.String()
		value := v.Field(index).Interface()
		key := t.Field(index).Tag.Get(keyTag)
		objMsg = append(objMsg, &Object{
			Field:   t.Field(index).Name,
			Type:    tt,
			Example: t.Field(index).Tag.Get(egTag),
			Key:     key,
			Value:   value,
			JSONKey: jsonKey,
		})
	}
	if len(objMsg) <= 0 {
		return nil, errors.New("can't get fields")
	}
	return objMsg, nil
}

func effective(object interface{}, tt string) bool {
	if tt == "[]string" {
		if valueStr, ok := object.([]string); ok {
			if valueStr == nil {
				return false
			}
		}
	}
	if tt == "string" {
		if valueStr, ok := object.(string); ok {
			return !(valueStr == "")
		}
	}
	return objMsg, nil
}
