package locations

import (
	//"fmt"
)

type LocationGroup interface {
	GetDisplay() []string
}

type LocationGroupImpl struct {
	locations []Location
}

func NewLocationGroup(locs []Location) LocationGroup {
	return &LocationGroupImpl {
		locations: locs,
	}
}

func (locgroup LocationGroupImpl) GetDisplay() []string {
	out := make([]string, 0)
	for _, v := range locgroup.locations {
		out = append(out, v.GetDisplay())
	}
	return out
}
