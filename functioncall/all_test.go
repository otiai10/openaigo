package functioncall

import (
	"encoding/json"
	"testing"

	. "github.com/otiai10/mint"
)

func TestFunctions(t *testing.T) {
	funcs := Funcs{}
	Expect(t, funcs).TypeOf("functioncall.Funcs")
}

func TestFunctions_MarshalJSON(t *testing.T) {
	repeat := func(word string, count int) (r string) {
		for i := 0; i < count; i++ {
			r += word
		}
		return r
	}
	funcs := Funcs{
		"repeat": Func{repeat, "Repeat given string N times", Params{
			{"word", "string", "String to be repeated", true},
			{"count", "number", "How many times to repeat", true},
		}},
	}
	b, err := funcs.MarshalJSON()
	Expect(t, err).ToBe(nil)

	v := []map[string]any{}
	err = json.Unmarshal(b, &v)
	Expect(t, err).ToBe(nil)

	Expect(t, v).Query("0.name").ToBe("repeat")
	Expect(t, v).Query("0.description").ToBe("Repeat given string N times")
	Expect(t, v).Query("0.parameters.type").ToBe("object")
	Expect(t, v).Query("0.parameters.properties.word.type").ToBe("string")
	Expect(t, v).Query("0.parameters.required.1").ToBe("count")
}

func TestAs(t *testing.T) {
	repeat := func(word string, count int) (r string) {
		for i := 0; i < count; i++ {
			r += word
		}
		return r
	}
	funcs := Funcs{
		"repeat": Func{repeat, "Repeat given string N times", Params{
			{"word", "string", "String to be repeated", true},
			{"count", "number", "How many times to repeat", true},
		}},
	}
	a := As[[]map[string]any](funcs)
	Expect(t, a).TypeOf("[]map[string]interface {}")
	Expect(t, a).Query("0.name").ToBe("repeat")
	Expect(t, a).Query("0.parameters.type").ToBe("object")
}
