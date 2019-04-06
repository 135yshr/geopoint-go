package geopoint

import (
	"bytes"
	"database/sql/driver"
	"encoding/binary"
	"fmt"
)

// GeoPoint simple types GIS point
type GeoPoint [2]float64

func (g *GeoPoint) String() string {
	return fmt.Sprintf("POINT(%f %f)", g[0], g[1])
}

func (p GeoPoint) Value() (driver.Value, error) {
	return p.String(), nil
}

func (p *GeoPoint) Scan(val interface{}) error {
	b := val.([]byte)
	r := bytes.NewReader(b)

	var header [9]byte
	if err := binary.Read(r, binary.LittleEndian, &header); err != nil {
		return err
	}

	var lat float64
	if err := binary.Read(r, binary.LittleEndian, &lat); err != nil {
		return err
	}

	var lng float64
	if err := binary.Read(r, binary.LittleEndian, &lng); err != nil {
		return err
	}
	*p = GeoPoint{lat, lng}

	return nil
}
