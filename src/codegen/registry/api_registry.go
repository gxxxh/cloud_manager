package registry

type APIRegistry map[string]interface{}

func (r *APIRegistry) GetAPI(key string) interface{} {
	return (*r)[key]
}

func (r *APIRegistry) SetAPI(key string, value interface{}) {
	(*r)[key] = value
}

func (r *APIRegistry) DeleteAPI(key string) {
	delete(*r, key)
}

func NewRegistry() *APIRegistry {
	return &APIRegistry{}
}
