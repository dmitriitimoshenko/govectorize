package resources

type Resources struct {
	resources []string
}

func (r *Resources) GetList() []string {
	return r.resources
}

func (r *Resources) Remove(elementID int) {
	var resources []string
	for i, resource := range r.resources {
		if i == elementID {
			continue
		}
		resources = append(resources, resource)
	}
	r.resources = resources
}

func (r *Resources) SetList(resources []string) {
	r.resources = resources
}
