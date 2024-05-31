package gopostgis_test

import (
	"testing"

	gopostgis "github.com/asif-mahmud/go-postgis"
)

func TestNil(t *testing.T) {
	d, e := gopostgis.NewHexEWKBDecoder(nil)
	if d != nil {
		t.Error("should not construct a decoder for nil")
	}

	if e == nil {
		t.Error("should not construct a decoder for nil")
	}
}

const pointType = 0x00000001

func testHex(s string, t *testing.T) {
	d, e := gopostgis.NewHexEWKBDecoder([]byte(s))
	if e != nil {
		t.Error(e)
	}

	if d.Type() != pointType {
		t.Error("datatype did not match. expected:", pointType, "found:", d.Type())
	}

	xy := make([]float64, 2)

	if e := d.ReadAny(xy); e != nil {
		t.Error(e)
	}

	if xy[0] != 10 || xy[1] != 20 {
		t.Error("data did not match. expected:", 10, 20, "found:", xy[0], xy[1])
	}
}

func TestLittleEndian(t *testing.T) {
	// POINT(10 20)
	s := "010100000000000000000024400000000000003440"
	testHex(s, t)
}

func TestBigEndian(t *testing.T) {
	// POINT(10 20)
	s := "000000000140240000000000004034000000000000"
	testHex(s, t)
}
