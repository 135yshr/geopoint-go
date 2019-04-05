package geopoint

import (
	"database/sql/driver"
	"fmt"
)

// GeoPoint simple types GIS point
type GeoPoint [2]float64

func (g *GeoPoint) String() string {
	return fmt.Sprintf("POINT(%f, %f)", g[0], g[1])
}

func (p GeoPoint) Value() (driver.Value, error) {
	return p.String(), nil
}
