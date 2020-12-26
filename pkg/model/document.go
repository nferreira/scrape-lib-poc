package model

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/robertkrimen/otto"
	"strings"
)

type Document struct {
	doc *goquery.Document
}

func NewDocument(vm *otto.Otto, doc *goquery.Document) (object *otto.Object) {
	var err error
	if object, err = vm.Object("Document = {}"); err != nil {
		return nil
	}
	document := &Document{
		doc: doc,
	}
	object.Set("__document__", document)
	object.Set("getElementById", func(call otto.FunctionCall) otto.Value {
		fmt.Printf("getElementById: %+v\n", call.ArgumentList)
		id := call.Argument(0).String()
		if !strings.HasPrefix(id, "#") {
			id = fmt.Sprintf("#%s", id)
		}
		selection := doc.Find(id)
		if len(selection.Nodes) > 0 {
			result, err := vm.ToValue(selection.Nodes)
			if err != nil {
				return vm.MakeCustomError("Error", err.Error())
			}
			return result
		} else {
			return otto.NullValue()
		}

	})
	return object
}