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
