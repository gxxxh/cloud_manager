package registry

type Registry map[string]interface{}

func (r *Registry) GetAPI(apiName string) interface{} {
	return (*r)[apiName]
}

func (r *Registry) SetAPI(apiName string, api interface{}) {
	(*r)[apiName] = api
}
