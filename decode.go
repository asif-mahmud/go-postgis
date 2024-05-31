//go:build linux || windows || darwin
// +build linux windows darwin

package gopostgis

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
)

// Decoder for hex encoded EWKB data.
// reference - https://github.com/postgis/postgis/blob/master/doc/bnf-wkb.txt
// It does not implement any specific postgis datatype, instead it is
// intended to be used as a reader for specific postgis datatype implementation.
// Instance of this struct should be initialized by [NewHexEWKBDecoder] function.
// Creating an instance of this struct directly will cause incorrect behavior
type HexEWKBDecoder interface {
	// Decoded datatype that represents a postgis datatype.
	// Example - 0x00000001 means it is Point type.
	// See the reference doc for full list of datatypes.
	Type() uint32

	// Reads a single byte
	ReadByte() (byte, error)

	// Reads a 32-bit unsigned integer
	ReadUint32() (uint32, error)

	// Reads a 64-bit float
	ReadDouble() (float64, error)

	// Reads into any fixed-size value
	ReadAny(interface{}) error
}

// Creates an instance of [HexEWKBDecoder].
func NewHexEWKBDecoder(data []byte) (HexEWKBDecoder, error) {
	hexData, err := hex.DecodeString(string(data))
	if err != nil {
		return nil, err
	}

	d := hexEWKBDecoder{
		buf: bytes.NewBuffer(hexData),
	}

	var order byte
	var dType uint32

	if err := binary.Read(d.buf, binary.LittleEndian, &order); err != nil {
		return nil, err
	}

	switch order {
	case 0x01:
		d.byteOrder = binary.LittleEndian

	case 0x00:
		d.byteOrder = binary.BigEndian

	default:
		return nil, fmt.Errorf("unknown byteorder. got: %v", order)
	}

	if err := binary.Read(d.buf, d.byteOrder, &dType); err != nil {
		return nil, err
	}

	d.dType = dType

	return &d, nil
}

type hexEWKBDecoder struct {
	buf       io.Reader
	byteOrder binary.ByteOrder

	dType uint32
}

// ReadAny implements HexEWKBDecoder.
func (h *hexEWKBDecoder) ReadAny(v interface{}) error {
	if err := binary.Read(h.buf, h.byteOrder, v); err != nil {
		return err
	}

	return nil
}

// ReadByte implements HexEWKBDecoder.
func (h *hexEWKBDecoder) ReadByte() (byte, error) {
	var v uint8
	if err := binary.Read(h.buf, h.byteOrder, &v); err != nil {
		return 0, err
	}

	return v, nil
}

// ReadDouble implements HexEWKBDecoder.
func (h *hexEWKBDecoder) ReadDouble() (float64, error) {
	var v float64
	if err := binary.Read(h.buf, h.byteOrder, &v); err != nil {
		return 0, err
	}

	return v, nil
}

// ReadUint32 implements HexEWKBDecoder.
func (h *hexEWKBDecoder) ReadUint32() (uint32, error) {
	var v uint32
	if err := binary.Read(h.buf, h.byteOrder, &v); err != nil {
		return 0, err
	}

	return v, nil
}

// Type implements HexEWKBDecoder.
func (h *hexEWKBDecoder) Type() uint32 {
	return h.dType
}
