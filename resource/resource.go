package resource

type Resource interface {
	initialize() error
	GetData(request interface{}) (interface{}, error)
	GetTranslator() Translator
	GetKey() string
}

type logger interface {
	Log(string string)
}

type Translator interface {
	Translate(interface{}) interface{}
}


func Initialize(resources []Resource, logger logger) {
	for _, resource := range resources {
		err := resource.initialize()
		if err != nil {
			logger.Log("Failed to initialize resource " + err.Error())
			panic(err)
		}
	}
}
