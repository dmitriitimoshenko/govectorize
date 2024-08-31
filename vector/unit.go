package vector

type Unit string

func (vu Unit) String() string {
	return string(vu)
}

func NewVectorUnit(s string) Unit {
	return Unit(s)
}
