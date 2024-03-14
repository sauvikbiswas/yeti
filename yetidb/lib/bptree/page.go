package bptree

import (
	"encoding/binary"
	"fmt"
)

type Page []byte

const (
	PageType_NODE uint16 = 1
	PageType_LEAF uint16 = 2
)

// Layout of a page
// ----------------
// page-type: 2B
// numKeys: 2B
// pointers: numKeys * 8B
// offsets: numKeys * 2B
// key-values:
// 	key-0-length: 2B
// 	val-0-length: 2B
// 	key-0: ...
// 	val-0: ...
// 	...
// 	key-(n-1)-length: 2B
// 	val-(n-1)-length: 2B
// 	key-(n-1): ...
// 	val-(n-1): ...
//  <free-space>

const (
	Page_SIZE         = 4096
	Page_HEADER_SIZE  = 4
	Page_MAX_KEY_SIZE = 1000
	Page_MAX_VAL_SIZE = 3000
)

func NewPage() (Page, error) {
	// assumption: there is at least one pointer and one offset
	if (Page_HEADER_SIZE + 8 + 2 + 4 + Page_MAX_KEY_SIZE + Page_MAX_VAL_SIZE) > Page_SIZE {
		return nil, fmt.Errorf("cannot create page with given constants. page size: %d, header size: %d, max key size: %d, max val size: %d", Page_SIZE, Page_HEADER_SIZE, Page_MAX_KEY_SIZE, Page_MAX_VAL_SIZE)
	}
	return make([]byte, Page_SIZE), nil
}

func (p Page) toString() string {
	return fmt.Sprintf("%v", p)
}

func (p Page) toBytes() []byte {
	return ([]byte)(p)
}

// headers
func (p Page) getPageType() uint16 {
	return binary.LittleEndian.Uint16(p)
}

func (p Page) getNumKeys() uint16 {
	return binary.LittleEndian.Uint16(p[2:4])
}

func (p Page) setHeader(pageType uint16, numKeys uint16) {
	binary.LittleEndian.PutUint16(p[0:2], pageType)
	binary.LittleEndian.PutUint16(p[2:4], numKeys)
}

// pointers
func (p Page) getPointer(id uint16) (uint64, error) {
	if id >= p.getNumKeys() {
		return 0, fmt.Errorf("id must be less than number of keys (%d) in page", p.getNumKeys())
	}
	pos := Page_HEADER_SIZE + 8*id
	return binary.LittleEndian.Uint64(p[pos:]), nil
}

func (p Page) setPointer(id uint16, val uint64) error {
	if id >= p.getNumKeys() {
		return fmt.Errorf("id must be less than number of keys (%d) in page", p.getNumKeys())
	}
	pos := Page_HEADER_SIZE + 8*id
	binary.LittleEndian.PutUint64(p[pos:], val)
	return nil
}

// offsets
func (p Page) offsetPosition(id uint16) (uint16, error) {
	if id <= 0 || id >= p.getNumKeys() {
		return 0, fmt.Errorf("offset id must be greater than 1 and less than number of keys (%d) in page", p.getNumKeys())
	}
	return Page_HEADER_SIZE + 8*p.getNumKeys() + 2*(id-1), nil
}

func (p Page) getOffset(id uint16) (uint16, error) {
	if id == 0 {
		return 0, nil
	}
	offset, err := p.offsetPosition(id)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(p[offset:]), nil
}

// key-value
func (p Page) kvPosition(id uint16) (uint16, error) {
	offset, err := p.getOffset(id)
	if err != nil {
		return 0, err
	}
	return Page_HEADER_SIZE + 8*p.getNumKeys() + 2*p.getNumKeys() + offset, nil
}

func (p Page) setKey(id uint16, key []byte) error {
	pos, err := p.kvPosition(id)
	if err != nil {
		return err
	}
	klen := uint16(len(key))
	binary.LittleEndian.PutUint16(p[pos:], uint16(klen))
	var i uint16
	for i = 0; i < klen; i += 2 {
		u16 := binary.LittleEndian.Uint16(key[i:])
		binary.LittleEndian.PutUint16(p[pos+4+i:], u16)
	}
	return nil
}

func (p Page) setValue(id uint16, value []byte) error {
	pos, err := p.kvPosition(id)
	if err != nil {
		return err
	}
	klen := binary.LittleEndian.Uint16(p[pos+0:])
	vlen := uint16(len(value))
	binary.LittleEndian.PutUint16(p[pos+2:], uint16(vlen))
	var i uint16
	for i = 0; i < vlen; i += 2 {
		u16 := binary.LittleEndian.Uint16(value[i:])
		binary.LittleEndian.PutUint16(p[pos+4+klen+i:], u16)
	}
	return nil
}

func (p Page) setKeyValue(id uint16, key []byte, value []byte) error {
	err := p.setKey(id, key)
	if err != nil {
		return err
	}
	err = p.setValue(id, value)
	return err
}

func (p Page) getKey(id uint16) ([]byte, error) {
	pos, err := p.kvPosition(id)
	if err != nil {
		return nil, err
	}
	klen := binary.LittleEndian.Uint16(p[pos:])
	return p[pos+4:][:klen], nil
}

func (p Page) getValue(id uint16) ([]byte, error) {
	pos, err := p.kvPosition(id)
	if err != nil {
		return nil, err
	}
	klen := binary.LittleEndian.Uint16(p[pos+0:])
	vlen := binary.LittleEndian.Uint16(p[pos+2:])
	return p[pos+4+klen:][:vlen], nil
}

// page size in bytes
func (p Page) size() uint16 {
	res, _ := p.kvPosition(p.getNumKeys())
	return res
}
