package bptree

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHeader(t *testing.T) {
	p, err := NewPage()
	require.NoError(t, err)

	p.setHeader(PageType_LEAF, 8)
	assert.Equal(t, PageType_LEAF, p.getPageType())
	assert.Equal(t, 8, int(p.getMaxKeys()))

	expected := []uint8{2, 0, 8}
	assert.Equal(t, uint8ToByte(expected), p.toBytes()[:len(expected)])
}

func TestPointerInPlaceReplacement(t *testing.T) {
	p, err := NewPage()
	require.NoError(t, err)

	p.setHeader(PageType_LEAF, 8)

	p.setNumKeys(8)

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
}

func TestKeyValueInPlaceReplacement(t *testing.T) {
	p, err := NewPage()
	require.NoError(t, err)

	p.setHeader(PageType_LEAF, 8)

	p.setNumKeys(8)

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
}

func TestKeyValueInsertion(t *testing.T) {
	p, err := NewPage()
	require.NoError(t, err)

	p.setHeader(PageType_LEAF, 8)

	err1 := p.insertKeyValue(0, []byte("test-key-0"), []byte("test-value-0"))
	assert.NoError(t, err1)

	for i := 7; i > 0; i-- {

		key := fmt.Sprintf("test-key-%d", i)
		val := fmt.Sprintf("test-value-%d", i)

		err2 := p.insertKeyValue(1, []byte(key), []byte(val))
		assert.NoError(t, err2)
	}

	err3 := p.insertKeyValue(8, []byte("test-key-8"), []byte("test-value-8"))
	assert.Error(t, err3)

	assert.Equal(t, uint16(8), p.getNumKeys())

	for i := 0; i < 8; i++ {

		expectedKey := fmt.Sprintf("test-key-%d", i)
		key, err4 := p.getKey(uint16(i))
		assert.NoError(t, err4)
		assert.Equal(t, []byte(expectedKey), key)

		expectedVal := fmt.Sprintf("test-value-%d", i)
		val, err5 := p.getValue(uint16(i))
		assert.NoError(t, err5)
		assert.Equal(t, []byte(expectedVal), val)
	}
}

func TestKeyValueReplacement(t *testing.T) {
	p, err := NewPage()
	require.NoError(t, err)

	p.setHeader(PageType_LEAF, 4)

	for i := 0; i < 4; i++ {

		key := fmt.Sprintf("test-key-%d", i)
		val := fmt.Sprintf("test-value-%d", i)

		err1 := p.insertKeyValue(uint16(i), []byte(key), []byte(val))
		assert.NoError(t, err1)
	}

	err2 := p.setKey(3, []byte("replaced-key"))
	assert.Error(t, err2)

	err3 := p.setValue(3, []byte("replaced-value"))
	assert.Error(t, err3)

}

// func TestKeyPositionOrLess(t *testing.T) {
// 	p, err := NewPage()
// 	require.NoError(t, err)

// 	p.setHeader(PageType_LEAF, 8)

// 	// err1 := p.insertKeyValue(0, []byte("test-key-0"), []byte("test-value-0"))
// 	// assert.NoError(t, err1)

// 	for i := 7; i >= 0; i-- {

// 		key := fmt.Sprintf("test-key-%d", i*2)
// 		val := fmt.Sprintf("test-value-%d", i*2)

// 		err2 := p.insertKeyValue(0, []byte(key), []byte(val))
// 		assert.NoError(t, err2)
// 	}

// 	for i := uint16(0); i < 8; i++ {
// 		key, _ := p.getKey(i)
// 		val, _ := p.getValue(i)

// 		fmt.Println(">" + string(key) + ":" + string(val) + "<")
// 	}
// }

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
