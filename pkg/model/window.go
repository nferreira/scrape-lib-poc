package model

import (
	"fmt"
	"github.com/robertkrimen/otto"
)

type Window struct {
}

func NewWindow(vm *otto.Otto, document *otto.Object) (object *otto.Object) {
	var err error
	if object, err = vm.Object("Document = {}"); err != nil {
		panic(err)
	}
	window := &Window{}
	object.Set("__window__", window)
	object.Set("document", document)
	object.Set("alert", func(call otto.FunctionCall) otto.Value {
		fmt.Printf("alert: %+v\n", call.ArgumentList)
		return otto.UndefinedValue()
	})
	return object
}
