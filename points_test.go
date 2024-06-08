package gopostgis_test

import (
	"database/sql/driver"
	"encoding/json"
	"testing"

	gopostgis "github.com/asif-mahmud/go-postgis"
)

func TestPoint(t *testing.T) {
	t.Run("scan", func(t *testing.T) {
		// POINT(10 20)
		s := "010100000000000000000024400000000000003440"
		var p gopostgis.Point
		expected := gopostgis.Point{
			X:     10,
			Y:     20,
			Valid: true,
		}
		e := p.Scan([]byte(s))
		if e != nil {
			t.Error(e)
		}
		if !p.Valid || p.X != expected.X || p.Y != expected.Y {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("scan null", func(t *testing.T) {
		var p gopostgis.Point
		expected := gopostgis.Point{
			X:     0,
			Y:     0,
			Valid: false,
		}
		e := p.Scan(nil)
		if e != nil {
			t.Error(e)
		}
		if p.Valid || p.X != expected.X || p.Y != expected.Y {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("value", func(t *testing.T) {
		p := gopostgis.Point{
			X:     10,
			Y:     20,
			Valid: true,
		}
		found, e := p.Value()
		expected := "POINT(10 20)"
		if e != nil {
			t.Error(e)
		}
		if expected != found.(string) {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("value null", func(t *testing.T) {
		p := gopostgis.Point{
			X:     0,
			Y:     0,
			Valid: false,
		}
		found, e := p.Value()
		var expected driver.Value = nil
		if e != nil {
			t.Error(e)
		}
		if expected != found {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("marshal", func(t *testing.T) {
		p := gopostgis.Point{
			X:     10,
			Y:     20,
			Valid: true,
		}
		expected := `{"X":10,"Y":20}`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("marshal null", func(t *testing.T) {
		p := gopostgis.Point{
			X:     0,
			Y:     0,
			Valid: false,
		}
		expected := `null`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("unmarshal", func(t *testing.T) {
		var p gopostgis.Point
		s := `{"X":10,"Y":20}`
		expected := gopostgis.Point{
			X:     10,
			Y:     20,
			Valid: true,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if !p.Valid || p.X != expected.X || p.Y != expected.Y {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("unmarshal null", func(t *testing.T) {
		var p gopostgis.Point
		s := `null`
		expected := gopostgis.Point{
			X:     0,
			Y:     0,
			Valid: false,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if p.Valid || p.X != expected.X || p.Y != expected.Y {
			t.Error("expected:", expected, "found:", p)
		}
	})
}

func TestPointS(t *testing.T) {
	t.Run("scan", func(t *testing.T) {
		// SRID=4326;POINT(10 20)
		s := "0101000020E610000000000000000024400000000000003440"
		var p gopostgis.PointS
		expected := gopostgis.PointS{
			SRID:  4326,
			X:     10,
			Y:     20,
			Valid: true,
		}
		e := p.Scan([]byte(s))
		if e != nil {
			t.Error(e)
		}
		if !p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("scan null", func(t *testing.T) {
		var p gopostgis.PointS
		expected := gopostgis.PointS{
			SRID:  0,
			X:     0,
			Y:     0,
			Valid: false,
		}
		e := p.Scan(nil)
		if e != nil {
			t.Error(e)
		}
		if p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("value", func(t *testing.T) {
		p := gopostgis.PointS{
			SRID:  4326,
			X:     10,
			Y:     20,
			Valid: true,
		}
		found, e := p.Value()
		expected := "SRID=4326;POINT(10 20)"
		if e != nil {
			t.Error(e)
		}
		if expected != found.(string) {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("value null", func(t *testing.T) {
		p := gopostgis.PointS{
			SRID:  0,
			X:     0,
			Y:     0,
			Valid: false,
		}
		found, e := p.Value()
		var expected driver.Value = nil
		if e != nil {
			t.Error(e)
		}
		if expected != found {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("marshal", func(t *testing.T) {
		p := gopostgis.PointS{
			SRID:  4326,
			X:     10,
			Y:     20,
			Valid: true,
		}
		expected := `{"SRID":4326,"X":10,"Y":20}`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("marshal null", func(t *testing.T) {
		p := gopostgis.PointS{
			SRID:  0,
			X:     0,
			Y:     0,
			Valid: false,
		}
		expected := `null`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("unmarshal", func(t *testing.T) {
		var p gopostgis.PointS
		s := `{"SRID":4326,"X":10,"Y":20}`
		expected := gopostgis.PointS{
			SRID:  4326,
			X:     10,
			Y:     20,
			Valid: true,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if !p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("unmarshal null", func(t *testing.T) {
		var p gopostgis.PointS
		s := `null`
		expected := gopostgis.PointS{
			SRID:  0,
			X:     0,
			Y:     0,
			Valid: false,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y {
			t.Error("expected:", expected, "found:", p)
		}
	})
}

func TestPointZ(t *testing.T) {
	t.Run("scan", func(t *testing.T) {
		// POINT(10 20 30)
		s := "0101000080000000000000244000000000000034400000000000003E40"
		var p gopostgis.PointZ
		expected := gopostgis.PointZ{
			X:     10,
			Y:     20,
			Z:     30,
			Valid: true,
		}
		e := p.Scan([]byte(s))
		if e != nil {
			t.Error(e)
		}
		if !p.Valid || p.X != expected.X || p.Y != expected.Y || p.Z != expected.Z {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("scan null", func(t *testing.T) {
		var p gopostgis.PointZ
		expected := gopostgis.PointZ{
			X:     0,
			Y:     0,
			Z:     0,
			Valid: false,
		}
		e := p.Scan(nil)
		if e != nil {
			t.Error(e)
		}
		if p.Valid || p.X != expected.X || p.Y != expected.Y || p.Z != expected.Z {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("value", func(t *testing.T) {
		p := gopostgis.PointZ{
			X:     10,
			Y:     20,
			Z:     30,
			Valid: true,
		}
		found, e := p.Value()
		expected := "POINT(10 20 30)"
		if e != nil {
			t.Error(e)
		}
		if expected != found.(string) {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("value null", func(t *testing.T) {
		p := gopostgis.PointZ{
			X:     0,
			Y:     0,
			Z:     0,
			Valid: false,
		}
		found, e := p.Value()
		var expected driver.Value = nil
		if e != nil {
			t.Error(e)
		}
		if expected != found {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("marshal", func(t *testing.T) {
		p := gopostgis.PointZ{
			X:     10,
			Y:     20,
			Z:     30,
			Valid: true,
		}
		expected := `{"X":10,"Y":20,"Z":30}`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("marshal null", func(t *testing.T) {
		p := gopostgis.PointZ{
			X:     0,
			Y:     0,
			Z:     0,
			Valid: false,
		}
		expected := `null`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("unmarshal", func(t *testing.T) {
		var p gopostgis.PointZ
		s := `{"X":10,"Y":20,"Z":30}`
		expected := gopostgis.PointZ{
			X:     10,
			Y:     20,
			Z:     30,
			Valid: true,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if !p.Valid || p.X != expected.X || p.Y != expected.Y || p.Z != expected.Z {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("unmarshal null", func(t *testing.T) {
		var p gopostgis.PointZ
		s := `null`
		expected := gopostgis.PointZ{
			X:     0,
			Y:     0,
			Z:     0,
			Valid: false,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if p.Valid || p.X != expected.X || p.Y != expected.Y || p.Z != expected.Z {
			t.Error("expected:", expected, "found:", p)
		}
	})
}

func TestPointZS(t *testing.T) {
	t.Run("scan", func(t *testing.T) {
		// SRID=4326;POINT(10 20 30)
		s := "01010000A0E6100000000000000000244000000000000034400000000000003E40"
		var p gopostgis.PointZS
		expected := gopostgis.PointZS{
			SRID:  4326,
			X:     10,
			Y:     20,
			Z:     30,
			Valid: true,
		}
		e := p.Scan([]byte(s))
		if e != nil {
			t.Error(e)
		}
		if !p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y ||
			p.Z != expected.Z {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("scan null", func(t *testing.T) {
		var p gopostgis.PointZS
		expected := gopostgis.PointZS{
			SRID:  0,
			X:     0,
			Y:     0,
			Z:     0,
			Valid: false,
		}
		e := p.Scan(nil)
		if e != nil {
			t.Error(e)
		}
		if p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y ||
			p.Z != expected.Z {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("value", func(t *testing.T) {
		p := gopostgis.PointZS{
			SRID:  4326,
			X:     10,
			Y:     20,
			Z:     30,
			Valid: true,
		}
		found, e := p.Value()
		expected := "SRID=4326;POINT(10 20 30)"
		if e != nil {
			t.Error(e)
		}
		if expected != found.(string) {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("value null", func(t *testing.T) {
		p := gopostgis.PointZS{
			SRID:  0,
			X:     0,
			Y:     0,
			Z:     0,
			Valid: false,
		}
		found, e := p.Value()
		var expected driver.Value = nil
		if e != nil {
			t.Error(e)
		}
		if expected != found {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("marshal", func(t *testing.T) {
		p := gopostgis.PointZS{
			SRID:  4326,
			X:     10,
			Y:     20,
			Z:     30,
			Valid: true,
		}
		expected := `{"SRID":4326,"X":10,"Y":20,"Z":30}`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("marshal null", func(t *testing.T) {
		p := gopostgis.PointZS{
			SRID:  0,
			X:     0,
			Y:     0,
			Z:     0,
			Valid: false,
		}
		expected := `null`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("unmarshal", func(t *testing.T) {
		var p gopostgis.PointZS
		s := `{"SRID":4326,"X":10,"Y":20,"Z":30}`
		expected := gopostgis.PointZS{
			SRID:  4326,
			X:     10,
			Y:     20,
			Z:     30,
			Valid: true,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if !p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y ||
			p.Z != expected.Z {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("unmarshal null", func(t *testing.T) {
		var p gopostgis.PointZS
		s := `null`
		expected := gopostgis.PointZS{
			SRID:  0,
			X:     0,
			Y:     0,
			Z:     0,
			Valid: false,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y ||
			p.Z != expected.Z {
			t.Error("expected:", expected, "found:", p)
		}
	})
}

func TestPointM(t *testing.T) {
	t.Run("scan", func(t *testing.T) {
		// POINT(10 20 30)
		s := "0101000080000000000000244000000000000034400000000000003E40"
		var p gopostgis.PointM
		expected := gopostgis.PointM{
			X:     10,
			Y:     20,
			M:     30,
			Valid: true,
		}
		e := p.Scan([]byte(s))
		if e != nil {
			t.Error(e)
		}
		if !p.Valid || p.X != expected.X || p.Y != expected.Y || p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("scan null", func(t *testing.T) {
		var p gopostgis.PointM
		expected := gopostgis.PointM{
			X:     0,
			Y:     0,
			M:     0,
			Valid: false,
		}
		e := p.Scan(nil)
		if e != nil {
			t.Error(e)
		}
		if p.Valid || p.X != expected.X || p.Y != expected.Y || p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("value", func(t *testing.T) {
		p := gopostgis.PointM{
			X:     10,
			Y:     20,
			M:     30,
			Valid: true,
		}
		found, e := p.Value()
		expected := "POINT(10 20 30)"
		if e != nil {
			t.Error(e)
		}
		if expected != found.(string) {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("value null", func(t *testing.T) {
		p := gopostgis.PointM{
			X:     0,
			Y:     0,
			M:     0,
			Valid: false,
		}
		found, e := p.Value()
		var expected driver.Value = nil
		if e != nil {
			t.Error(e)
		}
		if expected != found {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("marshal", func(t *testing.T) {
		p := gopostgis.PointM{
			X:     10,
			Y:     20,
			M:     30,
			Valid: true,
		}
		expected := `{"X":10,"Y":20,"M":30}`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("marshal null", func(t *testing.T) {
		p := gopostgis.PointM{
			X:     0,
			Y:     0,
			M:     0,
			Valid: false,
		}
		expected := `null`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("unmarshal", func(t *testing.T) {
		var p gopostgis.PointM
		s := `{"X":10,"Y":20,"M":30}`
		expected := gopostgis.PointM{
			X:     10,
			Y:     20,
			M:     30,
			Valid: true,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if !p.Valid || p.X != expected.X || p.Y != expected.Y || p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("unmarshal null", func(t *testing.T) {
		var p gopostgis.PointM
		s := `null`
		expected := gopostgis.PointM{
			X:     0,
			Y:     0,
			M:     0,
			Valid: false,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if p.Valid || p.X != expected.X || p.Y != expected.Y || p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})
}

func TestPointMS(t *testing.T) {
	t.Run("scan", func(t *testing.T) {
		// SRID=4326;POINT(10 20 30)
		s := "01010000A0E6100000000000000000244000000000000034400000000000003E40"
		var p gopostgis.PointMS
		expected := gopostgis.PointMS{
			SRID:  4326,
			X:     10,
			Y:     20,
			M:     30,
			Valid: true,
		}
		e := p.Scan([]byte(s))
		if e != nil {
			t.Error(e)
		}
		if !p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y ||
			p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("scan null", func(t *testing.T) {
		var p gopostgis.PointMS
		expected := gopostgis.PointMS{
			SRID:  0,
			X:     0,
			Y:     0,
			M:     0,
			Valid: false,
		}
		e := p.Scan(nil)
		if e != nil {
			t.Error(e)
		}
		if p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y ||
			p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("value", func(t *testing.T) {
		p := gopostgis.PointMS{
			SRID:  4326,
			X:     10,
			Y:     20,
			M:     30,
			Valid: true,
		}
		found, e := p.Value()
		expected := "SRID=4326;POINT(10 20 30)"
		if e != nil {
			t.Error(e)
		}
		if expected != found.(string) {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("value null", func(t *testing.T) {
		p := gopostgis.PointMS{
			SRID:  0,
			X:     0,
			Y:     0,
			M:     0,
			Valid: false,
		}
		found, e := p.Value()
		var expected driver.Value = nil
		if e != nil {
			t.Error(e)
		}
		if expected != found {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("marshal", func(t *testing.T) {
		p := gopostgis.PointMS{
			SRID:  4326,
			X:     10,
			Y:     20,
			M:     30,
			Valid: true,
		}
		expected := `{"SRID":4326,"X":10,"Y":20,"M":30}`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("marshal null", func(t *testing.T) {
		p := gopostgis.PointMS{
			SRID:  0,
			X:     0,
			Y:     0,
			M:     0,
			Valid: false,
		}
		expected := `null`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("unmarshal", func(t *testing.T) {
		var p gopostgis.PointMS
		s := `{"SRID":4326,"X":10,"Y":20,"M":30}`
		expected := gopostgis.PointMS{
			SRID:  4326,
			X:     10,
			Y:     20,
			M:     30,
			Valid: true,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if !p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y ||
			p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("unmarshal null", func(t *testing.T) {
		var p gopostgis.PointMS
		s := `null`
		expected := gopostgis.PointMS{
			SRID:  0,
			X:     0,
			Y:     0,
			M:     0,
			Valid: false,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y ||
			p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})
}

func TestPointZM(t *testing.T) {
	t.Run("scan", func(t *testing.T) {
		// POINT(10 20 30 40)
		s := "01010000C0000000000000244000000000000034400000000000003E400000000000004440"
		var p gopostgis.PointZM
		expected := gopostgis.PointZM{
			X:     10,
			Y:     20,
			Z:     30,
			M:     40,
			Valid: true,
		}
		e := p.Scan([]byte(s))
		if e != nil {
			t.Error(e)
		}
		if !p.Valid || p.X != expected.X || p.Y != expected.Y || p.Z != expected.Z ||
			p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("scan null", func(t *testing.T) {
		var p gopostgis.PointZM
		expected := gopostgis.PointZM{
			X:     0,
			Y:     0,
			Z:     0,
			M:     0,
			Valid: false,
		}
		e := p.Scan(nil)
		if e != nil {
			t.Error(e)
		}
		if p.Valid || p.X != expected.X || p.Y != expected.Y || p.Z != expected.Z ||
			p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("value", func(t *testing.T) {
		p := gopostgis.PointZM{
			X:     10,
			Y:     20,
			Z:     30,
			M:     40,
			Valid: true,
		}
		found, e := p.Value()
		expected := "POINT(10 20 30 40)"
		if e != nil {
			t.Error(e)
		}
		if expected != found.(string) {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("value null", func(t *testing.T) {
		p := gopostgis.PointZM{
			X:     0,
			Y:     0,
			Z:     0,
			M:     0,
			Valid: false,
		}
		found, e := p.Value()
		var expected driver.Value = nil
		if e != nil {
			t.Error(e)
		}
		if expected != found {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("marshal", func(t *testing.T) {
		p := gopostgis.PointZM{
			X:     10,
			Y:     20,
			Z:     30,
			M:     40,
			Valid: true,
		}
		expected := `{"X":10,"Y":20,"Z":30,"M":40}`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("marshal null", func(t *testing.T) {
		p := gopostgis.PointZM{
			X:     0,
			Y:     0,
			Z:     0,
			M:     0,
			Valid: false,
		}
		expected := `null`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("unmarshal", func(t *testing.T) {
		var p gopostgis.PointZM
		s := `{"X":10,"Y":20,"Z":30,"M":40}`
		expected := gopostgis.PointZM{
			X:     10,
			Y:     20,
			Z:     30,
			M:     40,
			Valid: true,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if !p.Valid || p.X != expected.X || p.Y != expected.Y || p.Z != expected.Z ||
			p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("unmarshal null", func(t *testing.T) {
		var p gopostgis.PointZM
		s := `null`
		expected := gopostgis.PointZM{
			X:     0,
			Y:     0,
			Z:     0,
			M:     0,
			Valid: false,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if p.Valid || p.X != expected.X || p.Y != expected.Y || p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})
}

func TestPointZMS(t *testing.T) {
	t.Run("scan", func(t *testing.T) {
		// SRID=4326;POINT(10 20 30 40)
		s := "01010000E0E6100000000000000000244000000000000034400000000000003E400000000000004440"
		var p gopostgis.PointZMS
		expected := gopostgis.PointZMS{
			SRID:  4326,
			X:     10,
			Y:     20,
			Z:     30,
			M:     40,
			Valid: true,
		}
		e := p.Scan([]byte(s))
		if e != nil {
			t.Error(e)
		}
		if !p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y ||
			p.Z != expected.Z || p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("scan null", func(t *testing.T) {
		var p gopostgis.PointZMS
		expected := gopostgis.PointZMS{
			SRID:  0,
			X:     0,
			Y:     0,
			Z:     0,
			M:     0,
			Valid: false,
		}
		e := p.Scan(nil)
		if e != nil {
			t.Error(e)
		}
		if p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y ||
			p.Z != expected.Z || p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("value", func(t *testing.T) {
		p := gopostgis.PointZMS{
			SRID:  4326,
			X:     10,
			Y:     20,
			Z:     30,
			M:     40,
			Valid: true,
		}
		found, e := p.Value()
		expected := "SRID=4326;POINT(10 20 30 40)"
		if e != nil {
			t.Error(e)
		}
		if expected != found.(string) {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("value null", func(t *testing.T) {
		p := gopostgis.PointZMS{
			SRID:  0,
			X:     0,
			Y:     0,
			Z:     0,
			M:     0,
			Valid: false,
		}
		found, e := p.Value()
		var expected driver.Value = nil
		if e != nil {
			t.Error(e)
		}
		if expected != found {
			t.Error("expected:", expected, "found:", found)
		}
	})

	t.Run("marshal", func(t *testing.T) {
		p := gopostgis.PointZMS{
			SRID:  4326,
			X:     10,
			Y:     20,
			Z:     30,
			M:     40,
			Valid: true,
		}
		expected := `{"SRID":4326,"X":10,"Y":20,"Z":30,"M":40}`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("marshal null", func(t *testing.T) {
		p := gopostgis.PointZMS{
			SRID:  0,
			X:     0,
			Y:     0,
			Z:     0,
			M:     0,
			Valid: false,
		}
		expected := `null`
		found, e := p.MarshalJSON()
		if e != nil {
			t.Error(e)
		}
		if expected != string(found) {
			t.Error("expected:", expected, "found:", string(found))
		}
	})

	t.Run("unmarshal", func(t *testing.T) {
		var p gopostgis.PointZMS
		s := `{"SRID":4326,"X":10,"Y":20,"Z":30,"M":40}`
		expected := gopostgis.PointZMS{
			SRID:  4326,
			X:     10,
			Y:     20,
			Z:     30,
			M:     40,
			Valid: true,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if !p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y ||
			p.Z != expected.Z || p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})

	t.Run("unmarshal null", func(t *testing.T) {
		var p gopostgis.PointZMS
		s := `null`
		expected := gopostgis.PointZMS{
			SRID:  0,
			X:     0,
			Y:     0,
			Z:     0,
			M:     0,
			Valid: false,
		}
		if e := json.Unmarshal([]byte(s), &p); e != nil {
			t.Error(e)
		}
		if p.Valid || p.SRID != expected.SRID || p.X != expected.X || p.Y != expected.Y ||
			p.Z != expected.Z || p.M != expected.M {
			t.Error("expected:", expected, "found:", p)
		}
	})
}
