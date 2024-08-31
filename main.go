package govectorize

import (
	"fmt"
	"github.com/dmitriitimoshenko/govectorize/resources"
	vct "github.com/dmitriitimoshenko/govectorize/vector"
	"math"
	"strings"
)

type Vectors struct {
	vectors []vct.Vector
}

func (v *Vectors) Map() []map[string]float64 {
	var m []map[string]float64
	for _, vector := range v.vectors {
		m = append(m, vector.GetMap())
	}
	return m
}

func GetVectors(descriptions []string) Vectors {
	var (
		vectors     []vct.Vector
		vectorUnits []vct.Unit
	)

	res := make([]resources.Resources, len(descriptions))

	for i, desc := range descriptions {
		res[i].SetList(strings.Fields(desc))
		for _, field := range strings.Fields(desc) {
			fmt.Println(field)
		}
		var elemsToRemove []int
		for ii, r := range res[i].GetList() {
			r = strings.ToLower(r)
			r = strings.Trim(r, "():,.-_=+|\\/&*^;%$#@!?\"")
			if r == "" {
				elemsToRemove = append(elemsToRemove, ii)
				continue
			}
			res[i].GetList()[ii] = r
			if !isVectorPresent(r, vectorUnits) {
				newVectorUnit := vct.NewVectorUnit(r)
				vectorUnits = append(vectorUnits, newVectorUnit)
			}
		}
		for _, elemToRemove := range elemsToRemove {
			res[i].Remove(elemToRemove)
		}
	}

	for _, resourcesPerServiceOrTicket := range res {
		vector := vct.Vector{}
		vector.Init()
		for _, vectorUnit := range vectorUnits {
			var val float64
			if isVectorUnitInResources(vectorUnit, resourcesPerServiceOrTicket) {
				val = getTF(vectorUnit.String(), resourcesPerServiceOrTicket.GetList()) * getIDF(vectorUnit.String(), res)
			}
			vector.Add(vectorUnit, val)
		}
		vectors = append(vectors, vector)
	}

	return Vectors{vectors: vectors}
}

func getTF(term string, list []string) float64 {
	i := float64(0)
	for _, e := range list {
		if term == e {
			i++
		}
	}
	return i / float64(len(list))
}

func getIDF(term string, docs []resources.Resources) float64 {
	i := float64(0)
	for _, list := range docs {
		for _, e := range list.GetList() {
			if term == e {
				i++
				break
			}
		}
	}

	if i == 0 {
		panic("i = 0")
	}
	return math.Log(3.0 / i)
}

func isVectorPresent(a string, list []vct.Unit) bool {
	for _, b := range list {
		if b.String() == a {
			return true
		}
	}
	return false
}

func isVectorUnitInResources(vu vct.Unit, resources resources.Resources) bool {
	for _, r := range resources.GetList() {
		if r == vu.String() {
			return true
		}
	}
	return false
}
