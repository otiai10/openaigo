package functioncall

import (
	"encoding/json"
)

type Funcs map[string]Func

type Func struct {
	Value       any    `json:"-"`
	Description string `json:"description,omitempty"`
	Parameters  Params `json:"parameters,omitempty"`
}

type Params []Param

type NestedParams []Param

type Param struct {
	Name        string `json:"-"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"-"`
	// Enum        []any  `json:"enum,omitempty"`
	Items NestedParams `json:",omitempty"`
}

func (funcs Funcs) MarshalJSON() ([]byte, error) {
	// Convert map to slice
	sl := []map[string]any{}
	for key, fun := range funcs {
		f := map[string]any{
			"name":        key,
			"description": fun.Description,
			"parameters":  fun.Parameters,
		}
		sl = append(sl, f)
	}
	return json.Marshal(sl)
}

func (params Params) MarshalJSON() ([]byte, error) {
	return marshalObject(params)
}

func (params NestedParams) MarshalJSON() ([]byte, error) {
	if len(params) == 1 {
		return json.Marshal(params[0])
	}

	return marshalObject(params)
}

func marshalObject[T ~[]Param](params T) ([]byte, error) {
	required := []string{}
	props := map[string]Param{}
	for _, p := range params {
		if p.Required {
			required = append(required, p.Name)
		}
		props[p.Name] = p
	}

	schema := map[string]any{
		"type":       "object",
		"properties": props,
		"required":   required,
	}
	return json.Marshal(schema)
}

func (param Param) MarshalJSON() ([]byte, error) {
	switch param.Type {
	case "array":
		schema := map[string]any{
			"type":  "array",
			"items": param.Items,
		}
		if param.Description != "" {
			schema["description"] = param.Description
		}
		return json.Marshal(schema)
	case "object":
		return marshalObject(param.Items)
	default:
		type Alias Param
		return json.Marshal(Alias(param))
	}
}

func As[T any](funcs Funcs) (dest T) {
	b, err := funcs.MarshalJSON()
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &dest)
	if err != nil {
		panic(err)
	}
	return dest
}
