# majma*

Is an data aggregator from multiple services sync or async

Supported services:
* HTTP request

Example:
```go
import (
	"majma/handler"
	"majma/resource"
)

type http struct{}

func (http) Translate(json interface{}) interface{} {
	return resource.HttpRequest{}
}

func main() {
	tr := new(http)
	res := []resource.Resource{resource.NewHttpResource(tr, "GET", "https://yesno.wtf/api", "yesno")}

	handler.Initialize(res)

	response := handler.Handle(nil)
}

```

You can also use your service just need to implement resource interface 

---
Todo:
- [ ] Add tests
- [ ] Support GRPC as a service
- [ ] Support database as a service
- [ ] Support caching
---
*In persian majma is some place that people gathering there
