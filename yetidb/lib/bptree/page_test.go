package bptree

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeader(t *testing.T) {
	p, err := NewPage()
	require.NoError(t, err)

	p.setHeader(PageType_LEAF, 8)
	assert.Equal(t, PageType_LEAF, p.getPageType())
	assert.Equal(t, 8, int(p.getNumKeys()))

	expected := []uint8{2, 0, 8}
	assert.Equal(t, uint8ToByte(expected), p.toBytes()[:len(expected)])
}

func TestPointer(t *testing.T) {
	p, err := NewPage()
	require.NoError(t, err)

	p.setHeader(PageType_LEAF, 8)

	err1 := p.setPointer(1, 32)
	assert.NoError(t, err1)

	err2 := p.setPointer(8, 32)
	assert.Error(t, err2)

	v1, err3 := p.getPointer(1)
	assert.NoError(t, err3)
	assert.Equal(t, uint64(32), v1)

	_, err4 := p.getPointer(8)
	assert.Error(t, err4)

	err5 := p.setPointer(5, 32)
	assert.NoError(t, err5)

	expected := []uint8{2, 0, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 32}
	assert.Equal(t, uint8ToByte(expected), p.toBytes()[:len(expected)])
}

func TestKeyValue(t *testing.T) {
	p, err := NewPage()
	require.NoError(t, err)

	p.setHeader(PageType_LEAF, 8)

	err1 := p.setKeyValue(0, []byte("test-key"), []byte("test-value"))
	assert.NoError(t, err1)

	err2 := p.setKeyValue(8, []byte("test-key"), []byte("test-value"))
	assert.Error(t, err2)

	key, err3 := p.getKey(0)
	assert.NoError(t, err3)
	assert.Equal(t, []byte("test-key"), key)

	val, err4 := p.getValue(0)
	assert.NoError(t, err4)
	assert.Equal(t, []byte("test-value"), val)

	_, err5 := p.getKey(8)
	assert.Error(t, err5)

	_, err6 := p.getValue(8)
	assert.Error(t, err6)

	expected := []uint8{2, 0, 8, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		8, 0, 10, 0, 116, 101, 115, 116, 45, 107, 101, 121, 116, 101, 115, 116, 45, 118, 97, 108, 117, 101}
	assert.Equal(t, uint8ToByte(expected), p.toBytes()[:len(expected)])
}

func uint8ToByte(in []uint8) []byte {
	out := make([]byte, len(in))
	for i, v := range in {
		out[i] = byte(v)
	}
	return out
}

func byteToUint8(in []byte) []uint8 {
	out := make([]uint8, len(in))
	for i, v := range in {
		out[i] = uint8(v)
	}
	return out
}

// Tests
// s, g: pointer
// offset errors
//
// Page exceeds check
