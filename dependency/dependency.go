package dependency

var dependencies map[string]interface{} = make(map[string]interface{})

func GetDependencies() *map[string]interface{} {
	return &dependencies
}

func Add(signature string, dependency interface{}) {
	dependencies[signature] = dependency
}

func Get(signature string) interface{} {
	return dependencies[signature]
}
