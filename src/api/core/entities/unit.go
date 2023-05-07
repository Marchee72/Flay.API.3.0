package entities

import (
	"fmt"
	"strconv"
	"strings"
)

type Unit struct {
	Floor     int
	Apartment string
}

func (u Unit) String() string {
	return fmt.Sprintf("%d-%s", u.Floor, u.Apartment)
}

func NewUnit(unit string) (*Unit, error) {
	u := strings.Split(unit, "-")
	floor, err := strconv.Atoi(u[0])
	if len(u) != 2 || err != nil {
		return nil, fmt.Errorf("cannot parse string to Unit: %s", unit)
	}
	return &Unit{
		Floor:     floor,
		Apartment: strings.ToUpper(u[1]),
	}, nil
}
