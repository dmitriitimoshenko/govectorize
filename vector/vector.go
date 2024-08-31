package vector

type Vector struct {
	vec map[string]float64
}

func (v *Vector) Init() {
	v.vec = make(map[string]float64)
}

func (v *Vector) Add(vu Unit, val float64) {
	v.vec[vu.String()] = val
}

func (v *Vector) GetMap() map[string]float64 {
	return v.vec
}
