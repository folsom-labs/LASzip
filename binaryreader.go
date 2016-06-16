package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"
)

// BinaryReader is a helper for reading binary data. It keeps track of Error
// so that the caller can reduce error checking boiler-plate from Read* functions
type BinaryReader struct {
	r             io.Reader
	BytesConsumed int
	Error         error
}

// NewBinaryReader creates a new binary reader
func NewBinaryReader(r io.Reader) *BinaryReader {
	return &BinaryReader{
		r: r,
	}
}

// ReadBytes reads nBytes bytes from the reader
func (r *BinaryReader) ReadBytes(nBytes int) []byte {
	if r.Error != nil {
		return nil
	}
	res := make([]byte, nBytes, nBytes)
	n, err := r.r.Read(res[:])
	if err == nil && n != nBytes {
		err = fmt.Errorf("ReadBytes: wanted to read %d bytes, only read %d", nBytes, n)
	}
	r.Error = err
	r.BytesConsumed += n
	return res
}

// ReadFixedString reads a fixed string of nChars characters
func (r *BinaryReader) ReadFixedString(nChars int) string {
	if r.Error != nil {
		return ""
	}
	data := make([]byte, nChars, nChars)
	n, err := r.r.Read(data[:])
	if err == nil && n != nChars {
		err = fmt.Errorf("ReadFixedString: wanted to read %d bytes, only read %d", nChars, n)
	}
	var res string
	if err == nil {
		res = string(data)
		res = strings.TrimRight(res, "\000")
	}
	r.Error = err
	r.BytesConsumed += n
	return res
}

// Skip skips n bytes
func (r *BinaryReader) Skip(nBytes int) {

	// TODO: why this breaks if enabled?
	if false {
		if seeker, ok := r.r.(io.Seeker); ok {
			_, err := seeker.Seek(int64(nBytes), 1)
			if err != nil {
				r.Error = err
			}
			return
		}
	}

	for nBytes > 0 {
		// don't read more than 16k at a time, to prevent using too much memory
		// at a time
		n := nBytes
		if n > 4096*4 {
			n = 4096 * 4
		}
		r.ReadFixedString(n)
		nBytes -= n
	}
}

// ReadUint8 reads unsigned byte
func (r *BinaryReader) ReadUint8() byte {
	var res uint8
	if r.Error != nil {
		return res
	}
	r.Error = binary.Read(r.r, binary.LittleEndian, &res)
	r.BytesConsumed++
	return res
}

// ReadInt8 reads signed byte
func (r *BinaryReader) ReadInt8() int8 {
	var res int8
	if r.Error != nil {
		return res
	}
	r.Error = binary.Read(r.r, binary.LittleEndian, &res)
	r.BytesConsumed++
	return res
}

// ReadUint16 reads uint16
func (r *BinaryReader) ReadUint16() uint16 {
	var res uint16
	if r.Error != nil {
		return res
	}
	r.Error = binary.Read(r.r, binary.LittleEndian, &res)
	r.BytesConsumed += 2
	return res
}

// ReadUint32 reads uint32
func (r *BinaryReader) ReadUint32() uint32 {
	var res uint32
	if r.Error != nil {
		return res
	}
	r.Error = binary.Read(r.r, binary.LittleEndian, &res)
	r.BytesConsumed += 4
	return res
}

// ReadInt32 reads int32
func (r *BinaryReader) ReadInt32() int32 {
	var res int32
	if r.Error != nil {
		return res
	}
	r.Error = binary.Read(r.r, binary.LittleEndian, &res)
	r.BytesConsumed += 4
	return res
}

// ReadFloat64 reads float64
func (r *BinaryReader) ReadFloat64() float64 {
	var res float64
	if r.Error != nil {
		return res
	}
	r.Error = binary.Read(r.r, binary.LittleEndian, &res)
	r.BytesConsumed += 8
	return res
}
