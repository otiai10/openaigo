package functioncall

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Invocation interface {
	Name() string
	Args() map[string]any
}

func (funcs Funcs) Call(invocation Invocation) string {
	b, err := json.Marshal(funcs.Invoke(invocation))
	if err != nil {
		return err.Error()
	}
	return string(b)
}

func (funcs Funcs) Invoke(invocation Invocation) any {
	f, ok := funcs[invocation.Name()]
	if !ok {
		return fmt.Sprintf("function not found: %s", invocation.Name())
	}
	v := reflect.ValueOf(f.Value)
	if !v.IsValid() || v.IsZero() {
		return fmt.Sprintf("function is invalid: %s", invocation.Name())
	}
	if v.Kind() != reflect.Func {
		return fmt.Sprintf("function is not a function: %s", invocation.Name())
	}
	if v.Type().NumIn() != len(invocation.Args()) {
		return fmt.Sprintf("function argument length mismatch: %s", invocation.Name())
	}
	// Call the function with given arguments by using `reflect` package
	args := invocation.Args()
	params := []reflect.Value{}
	for i, p := range f.Parameters {
		if arg, ok := args[p.Name]; ok {
			params = append(params, reflect.ValueOf(arg))
		} else {
			params = append(params, reflect.Zero(v.Type().In(i)))
		}
	}
	rets := []any{}
	for _, r := range v.Call(params) {
		rets = append(rets, r.Interface())
	}
	return rets
}
