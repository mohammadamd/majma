package handler

import (
	"fmt"
	"majma/handler/resource"
)

type Logger struct{}

func (logger Logger) Log(log string) {
	fmt.Println(log)
}

var resources []resource.Resource

func Initialize(r []resource.Resource) {
	resource.Initialize(r, nil)
	resources = r
}

func Handle(request interface{}) {
	response := map[string]interface{}{}

	for _, res := range resources {
		req := res.GetTranslator().Translate(request)

		r, err := res.GetData(nil, req)
		if err != nil {
			fmt.Println(err.Error())
		}

		response[res.GetKey()] = r
	}
}
