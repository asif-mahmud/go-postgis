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

// Point (X, Y) datatype.
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

// PointS (SRID X, Y) datatype.
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

// PointZ (X, Y, Z) datatype.
// Supports NULL value.
type PointZ struct {
	X     float64
	Y     float64
	Z     float64
	Valid bool `json:"-"`
}

// Scan implements sql.Scanner
func (p *PointZ) Scan(src any) error {
	if src == nil {
		p.X = 0
		p.Y = 0
		p.Z = 0
		p.Valid = false
		return nil
	}

	switch v := src.(type) {
	case []byte:
		d, e := NewHexEWKBDecoder(v)
		if e != nil {
			return e
		}
		values := make([]float64, 3)
		if e := d.ReadAny(values); e != nil {
			return e
		}
		p.X = values[0]
		p.Y = values[1]
		p.Z = values[2]
		p.Valid = true
		return nil

	default:
		return fmt.Errorf("driver datatype not supported")
	}
}

// Value implements driver.Valuer
func (p PointZ) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return fmt.Sprintf("POINT(%g %g %g)", p.X, p.Y, p.Z), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (p *PointZ) UnmarshalJSON(d []byte) error {
	if string(d) == jsonNullString {
		p.X = 0
		p.Y = 0
		p.Z = 0
		p.Valid = false
		return nil
	}

	type point PointZ
	var tp point

	if err := json.Unmarshal(d, &tp); err != nil {
		return err
	}

	p.X = tp.X
	p.Y = tp.Y
	p.Z = tp.Z
	p.Valid = true

	return nil
}

// MarshalJSON implements json.Marshaler
func (p PointZ) MarshalJSON() ([]byte, error) {
	if !p.Valid {
		return jsonNullValue, nil
	}

	type point PointZ
	tp := point(p)

	return json.Marshal(tp)
}

// PointZS (SRID X, Y, Z) datatype.
// Supports NULL value.
type PointZS struct {
	SRID  uint32
	X     float64
	Y     float64
	Z     float64
	Valid bool `json:"-"`
}

// Scan implements sql.Scanner
func (p *PointZS) Scan(src any) error {
	if src == nil {
		p.SRID = 0
		p.X = 0
		p.Y = 0
		p.Z = 0
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
		values := make([]float64, 3)
		if e := d.ReadAny(values); e != nil {
			return e
		}
		p.SRID = srid
		p.X = values[0]
		p.Y = values[1]
		p.Z = values[2]
		p.Valid = true
		return nil

	default:
		return fmt.Errorf("driver datatype not supported")
	}
}

// Value implements driver.Valuer
func (p PointZS) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return fmt.Sprintf("SRID=%d;POINT(%g %g %g)", p.SRID, p.X, p.Y, p.Z), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (p *PointZS) UnmarshalJSON(d []byte) error {
	if string(d) == jsonNullString {
		p.SRID = 0
		p.X = 0
		p.Y = 0
		p.Z = 0
		p.Valid = false
		return nil
	}

	type point PointZS
	var tp point

	if err := json.Unmarshal(d, &tp); err != nil {
		return err
	}

	p.SRID = tp.SRID
	p.X = tp.X
	p.Y = tp.Y
	p.Z = tp.Z
	p.Valid = true

	return nil
}

// MarshalJSON implements json.Marshaler
func (p PointZS) MarshalJSON() ([]byte, error) {
	if !p.Valid {
		return jsonNullValue, nil
	}

	type point PointZS
	tp := point(p)

	return json.Marshal(tp)
}

// PointM (X, Y, M) datatype.
// Supports NULL value.
type PointM struct {
	X     float64
	Y     float64
	M     float64
	Valid bool `json:"-"`
}

// Scan implements sql.Scanner
func (p *PointM) Scan(src any) error {
	if src == nil {
		p.X = 0
		p.Y = 0
		p.M = 0
		p.Valid = false
		return nil
	}

	switch v := src.(type) {
	case []byte:
		d, e := NewHexEWKBDecoder(v)
		if e != nil {
			return e
		}
		values := make([]float64, 3)
		if e := d.ReadAny(values); e != nil {
			return e
		}
		p.X = values[0]
		p.Y = values[1]
		p.M = values[2]
		p.Valid = true
		return nil

	default:
		return fmt.Errorf("driver datatype not supported")
	}
}

// Value implements driver.Valuer
func (p PointM) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return fmt.Sprintf("POINT(%g %g %g)", p.X, p.Y, p.M), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (p *PointM) UnmarshalJSON(d []byte) error {
	if string(d) == jsonNullString {
		p.X = 0
		p.Y = 0
		p.M = 0
		p.Valid = false
		return nil
	}

	type point PointM
	var tp point

	if err := json.Unmarshal(d, &tp); err != nil {
		return err
	}

	p.X = tp.X
	p.Y = tp.Y
	p.M = tp.M
	p.Valid = true

	return nil
}

// MarshalJSON implements json.Marshaler
func (p PointM) MarshalJSON() ([]byte, error) {
	if !p.Valid {
		return jsonNullValue, nil
	}

	type point PointM
	tp := point(p)

	return json.Marshal(tp)
}

// PointMS (SRID X, Y, M) datatype.
// Supports NULL value.
type PointMS struct {
	SRID  uint32
	X     float64
	Y     float64
	M     float64
	Valid bool `json:"-"`
}

// Scan implements sql.Scanner
func (p *PointMS) Scan(src any) error {
	if src == nil {
		p.SRID = 0
		p.X = 0
		p.Y = 0
		p.M = 0
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
		values := make([]float64, 3)
		if e := d.ReadAny(values); e != nil {
			return e
		}
		p.SRID = srid
		p.X = values[0]
		p.Y = values[1]
		p.M = values[2]
		p.Valid = true
		return nil

	default:
		return fmt.Errorf("driver datatype not supported")
	}
}

// Value implements driver.Valuer
func (p PointMS) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return fmt.Sprintf("SRID=%d;POINT(%g %g %g)", p.SRID, p.X, p.Y, p.M), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (p *PointMS) UnmarshalJSON(d []byte) error {
	if string(d) == jsonNullString {
		p.SRID = 0
		p.X = 0
		p.Y = 0
		p.M = 0
		p.Valid = false
		return nil
	}

	type point PointMS
	var tp point

	if err := json.Unmarshal(d, &tp); err != nil {
		return err
	}

	p.SRID = tp.SRID
	p.X = tp.X
	p.Y = tp.Y
	p.M = tp.M
	p.Valid = true

	return nil
}

// MarshalJSON implements json.Marshaler
func (p PointMS) MarshalJSON() ([]byte, error) {
	if !p.Valid {
		return jsonNullValue, nil
	}

	type point PointMS
	tp := point(p)

	return json.Marshal(tp)
}

// PointZM (X, Y, Z, M) datatype.
// Supports NULL value.
type PointZM struct {
	X     float64
	Y     float64
	Z     float64
	M     float64
	Valid bool `json:"-"`
}

// Scan implements sql.Scanner
func (p *PointZM) Scan(src any) error {
	if src == nil {
		p.X = 0
		p.Y = 0
		p.Z = 0
		p.M = 0
		p.Valid = false
		return nil
	}

	switch v := src.(type) {
	case []byte:
		d, e := NewHexEWKBDecoder(v)
		if e != nil {
			return e
		}
		values := make([]float64, 4)
		if e := d.ReadAny(values); e != nil {
			return e
		}
		p.X = values[0]
		p.Y = values[1]
		p.Z = values[2]
		p.M = values[3]
		p.Valid = true
		return nil

	default:
		return fmt.Errorf("driver datatype not supported")
	}
}

// Value implements driver.Valuer
func (p PointZM) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return fmt.Sprintf("POINT(%g %g %g %g)", p.X, p.Y, p.Z, p.M), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (p *PointZM) UnmarshalJSON(d []byte) error {
	if string(d) == jsonNullString {
		p.X = 0
		p.Y = 0
		p.Z = 0
		p.M = 0
		p.Valid = false
		return nil
	}

	type point PointZM
	var tp point

	if err := json.Unmarshal(d, &tp); err != nil {
		return err
	}

	p.X = tp.X
	p.Y = tp.Y
	p.Z = tp.Z
	p.M = tp.M
	p.Valid = true

	return nil
}

// MarshalJSON implements json.Marshaler
func (p PointZM) MarshalJSON() ([]byte, error) {
	if !p.Valid {
		return jsonNullValue, nil
	}

	type point PointZM
	tp := point(p)

	return json.Marshal(tp)
}

// PointZMS (SRID X, Y, Z, M) datatype.
// Supports NULL value.
type PointZMS struct {
	SRID  uint32
	X     float64
	Y     float64
	Z     float64
	M     float64
	Valid bool `json:"-"`
}

// Scan implements sql.Scanner
func (p *PointZMS) Scan(src any) error {
	if src == nil {
		p.SRID = 0
		p.X = 0
		p.Y = 0
		p.Z = 0
		p.M = 0
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
		values := make([]float64, 4)
		if e := d.ReadAny(values); e != nil {
			return e
		}
		p.SRID = srid
		p.X = values[0]
		p.Y = values[1]
		p.Z = values[2]
		p.M = values[3]
		p.Valid = true
		return nil

	default:
		return fmt.Errorf("driver datatype not supported")
	}
}

// Value implements driver.Valuer
func (p PointZMS) Value() (driver.Value, error) {
	if !p.Valid {
		return nil, nil
	}

	return fmt.Sprintf("SRID=%d;POINT(%g %g %g %g)", p.SRID, p.X, p.Y, p.Z, p.M), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (p *PointZMS) UnmarshalJSON(d []byte) error {
	if string(d) == jsonNullString {
		p.SRID = 0
		p.X = 0
		p.Y = 0
		p.Z = 0
		p.M = 0
		p.Valid = false
		return nil
	}

	type point PointZMS
	var tp point

	if err := json.Unmarshal(d, &tp); err != nil {
		return err
	}

	p.SRID = tp.SRID
	p.X = tp.X
	p.Y = tp.Y
	p.Z = tp.Z
	p.M = tp.M
	p.Valid = true

	return nil
}

// MarshalJSON implements json.Marshaler
func (p PointZMS) MarshalJSON() ([]byte, error) {
	if !p.Valid {
		return jsonNullValue, nil
	}

	type point PointZMS
	tp := point(p)

	return json.Marshal(tp)
}
