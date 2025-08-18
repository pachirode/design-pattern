package registry

type Registry struct {
	registry map[string]interface{}
}

func (r *Registry) Register(name string, obj interface{}) {
	r.registry[name] = obj
}

func (r *Registry) Get(name string) interface{} {
	return r.registry[name]
}
