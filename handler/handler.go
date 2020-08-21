package handler

import (
	"fmt"
	"majma/resource"
	"sync"
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

func Handle(request interface{}) map[string]interface{} {
	response := map[string]interface{}{}

	for _, res := range resources {
		req := res.GetTranslator().Translate(request)

		r, err := res.GetData(req)
		if err != nil {
			fmt.Println(err.Error())
		}

		response[res.GetKey()] = r
	}

	return response
}

func HandleAsync(response chan map[string]interface{}, req interface{}) {
	wg := sync.WaitGroup{}
	wg.Add(len(resources))
	for _, res := range resources {
		go func() {
			r, err := res.GetData(req)
			if err != nil {
				fmt.Println(err.Error())
			}
			response <- map[string]interface{}{res.GetKey(): r}
			wg.Done()
		}()
	}
	wg.Wait()
}
