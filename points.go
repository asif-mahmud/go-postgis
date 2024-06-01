package gopostgis

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

var (
	jsonNullString = "null"
	jsonNullValue  = []byte(jsonNullString)
)

// Point (X, Y) datatype
// Supports NULL value.
type Point struct {
	X     float64
	Y     float64
	Valid bool `json:"-"`
}

// Scan implements sql.Scanner
func (p *Point) Scan(src any) error {
	if src == nil {
		p.X = 0
		p.Y = 0
		p.Valid = false
		return nil
	}

	switch v := src.(type) {
	case []byte:
		d, e := NewHexEWKBDecoder(v)
		if e != nil {
			return e
		}
		values := make([]float64, 2)
		if e := d.ReadAny(values); e != nil {
			return e
		}
		p.X = values[0]
		p.Y = values[1]
		p.Valid = true
		return nil

	default:
		return fmt.Errorf("driver datatype not supported")
	}
}

// Value implements driver.Valuer
func (p Point) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return fmt.Sprintf("POINT(%g %g)", p.X, p.Y), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (p *Point) UnmarshalJSON(d []byte) error {
	if string(d) == jsonNullString {
		p.X = 0
		p.Y = 0
		p.Valid = false
		return nil
	}

	type point Point
	var tp point

	if err := json.Unmarshal(d, &tp); err != nil {
		return err
	}

	p.X = tp.X
	p.Y = tp.Y
	p.Valid = true

	return nil
}

// MarshalJSON implements json.Marshaler
func (p Point) MarshalJSON() ([]byte, error) {
	if !p.Valid {
		return jsonNullValue, nil
	}

	type point Point
	tp := point(p)

	return json.Marshal(tp)
}

// PointS (SRID X, Y) datatype
// Supports NULL value.
type PointS struct {
	SRID  uint32
	X     float64
	Y     float64
	Valid bool `json:"-"`
}

// Scan implements sql.Scanner
func (p *PointS) Scan(src any) error {
	if src == nil {
		p.SRID = 0
		p.X = 0
		p.Y = 0
		p.Valid = false
		return nil
	}

	switch v := src.(type) {
	case []byte:
		d, e := NewHexEWKBDecoder(v)
		if e != nil {
			return e
		}
		srid, e := d.ReadUint32()
		if e != nil {
			return e
		}
		values := make([]float64, 2)
		if e := d.ReadAny(values); e != nil {
			return e
		}
		p.SRID = srid
		p.X = values[0]
		p.Y = values[1]
		p.Valid = true
		return nil

	default:
		return fmt.Errorf("driver datatype not supported")
	}
}

// Value implements driver.Valuer
func (p PointS) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return fmt.Sprintf("SRID=%d;POINT(%g %g)", p.SRID, p.X, p.Y), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (p *PointS) UnmarshalJSON(d []byte) error {
	if string(d) == jsonNullString {
		p.SRID = 0
		p.X = 0
		p.Y = 0
		p.Valid = false
		return nil
	}

	type point PointS
	var tp point

	if err := json.Unmarshal(d, &tp); err != nil {
		return err
	}

	p.SRID = tp.SRID
	p.X = tp.X
	p.Y = tp.Y
	p.Valid = true

	return nil
}

// MarshalJSON implements json.Marshaler
func (p PointS) MarshalJSON() ([]byte, error) {
	if !p.Valid {
		return jsonNullValue, nil
	}

	type point PointS
	tp := point(p)

	return json.Marshal(tp)
}
