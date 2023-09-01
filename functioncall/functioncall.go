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

type Param struct {
	Name        string `json:"-"`
	Type        string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"-"`
	// Enum        []any  `json:"enum,omitempty"`
	Items Params `json:",omitempty"`
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
	required := []string{}
	props := map[string]any{}
	for _, p := range params {
		if p.Required {
			required = append(required, p.Name)
		}
		if p.Type == "array" && p.Items != nil {
			schema := map[string]any{
				"type":     "array",
				"items":    p.Items,
				"required": required,
			}
			props[p.Name] = schema
		} else {
			props[p.Name] = p
		}
	}
	schema := map[string]any{
		"type":       "object",
		"properties": props,
		"required":   required,
	}
	return json.Marshal(schema)
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
