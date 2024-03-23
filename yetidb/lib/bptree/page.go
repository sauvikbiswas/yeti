package bptree

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
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
// maxKeys: 2B
// numKeys: 2B
// pointers: maxKeys+1 * 8B
// offsets: maxKeys-1 * 2B
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
	Page_HEADER_SIZE  = 6
	Page_MAX_KEY_SIZE = 1000
	Page_MAX_VAL_SIZE = 3000
)

func NewPage() (Page, error) {
	// assumption: there is at least two pointers and one offset
	if (Page_HEADER_SIZE + 16 + 2 + Page_MAX_KEY_SIZE + Page_MAX_VAL_SIZE) > Page_SIZE {
		return nil, fmt.Errorf("cannot create page with given constants. page size: %d, header size: %d, max key size: %d, max val size: %d", Page_SIZE, Page_HEADER_SIZE, Page_MAX_KEY_SIZE, Page_MAX_VAL_SIZE)
	}
	return make([]byte, Page_SIZE), nil
}

func (p Page) copy() Page {
	pNew := make(Page, len(p))
	copy(pNew, p)
	return pNew
}

func (p Page) toString() string {
	return fmt.Sprintf("%v", p)
}

func (p Page) toBytes() []byte {
	return ([]byte)(p)
}

func (p Page) toHex() string {
	return hex.Dump(p)
}

// headers
func (p Page) getPageType() uint16 {
	return binary.LittleEndian.Uint16(p)
}

func (p Page) getMaxKeys() uint16 {
	return binary.LittleEndian.Uint16(p[2:4])
}

func (p Page) getNumKeys() uint16 {
	return binary.LittleEndian.Uint16(p[4:6])
}

func (p Page) setNumKeys(numKeys uint16) error {
	if numKeys > p.getMaxKeys() {
		return fmt.Errorf("number of keys cannot exceed maximum number of keys (%d) defined in page", p.getMaxKeys())
	}
	binary.LittleEndian.PutUint16(p[4:6], numKeys)
	return nil
}

func (p Page) setHeader(pageType uint16, maxKeys uint16) {
	binary.LittleEndian.PutUint16(p[0:2], pageType)
	binary.LittleEndian.PutUint16(p[2:4], maxKeys)
}

// pointers
func (p Page) getPointer(id uint16) (uint64, error) {
	if id >= p.getNumKeys() {
		return 0, fmt.Errorf("id must be less than or equal number of keys (%d) in page", p.getNumKeys())
	}
	pos := Page_HEADER_SIZE + 8*id
	return binary.LittleEndian.Uint64(p[pos:]), nil
}

func (p Page) setPointer(id uint16, val uint64) error {
	if id >= p.getNumKeys() {
		return fmt.Errorf("id must be less than or equal to number of keys (%d) in page", p.getNumKeys())
	}
	pos := Page_HEADER_SIZE + 8*id
	binary.LittleEndian.PutUint64(p[pos:], val)
	return nil
}

// offsets
func (p Page) offsetPosition(id uint16) (uint16, error) {
	if id <= 0 || id >= p.getNumKeys() {
		return 0, fmt.Errorf("offset id must be greater than 1 and less than or equal to number of keys (%d) in page", p.getNumKeys())
	}
	return Page_HEADER_SIZE + 8*(p.getMaxKeys()+1) + 2*(id-1), nil
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

func (p Page) setOffset(id uint16, value uint16) error {
	if id == 0 {
		return nil
	}
	offset, err := p.offsetPosition(id)
	if err != nil {
		return err
	}
	binary.LittleEndian.PutUint16(p[offset:], value)
	return nil
}

// key-value
func (p Page) kvPosition(id uint16) (uint16, error) {
	offset, err := p.getOffset(id)
	if err != nil {
		return 0, err
	}
	return Page_HEADER_SIZE + 8*(p.getMaxKeys()+1) + 2*(p.getMaxKeys()-1) + offset, nil
}

func (p Page) setKey(id uint16, key []byte) error {
	if id >= p.getNumKeys() {
		return fmt.Errorf("id must be less than number of keys (%d) in page", p.getNumKeys())
	}
	pos, err := p.kvPosition(id)
	if err != nil {
		return err
	}
	oldKey, err := p.getKey(id)
	if err != nil {
		return err
	}
	klen := len(key)
	oldKlen := len(oldKey)
	if ceilToUint16(klen) > ceilToUint16(oldKlen) {
		return fmt.Errorf("key size (%d) must be less than or equal to %d bytes for replacement", klen, ceilToUint16(oldKlen))
	}

	binary.LittleEndian.PutUint16(p[pos:], uint16(klen))
	for i := uint16(0); i < uint16(klen); i += 2 {
		if i == uint16(klen)-1 {
			b := make([]byte, 2)
			b[0] = byte(key[i])
			u16 := binary.LittleEndian.Uint16(b)
			binary.LittleEndian.PutUint16(p[pos+4+i:], u16)

		} else {
			u16 := binary.LittleEndian.Uint16(key[i:])
			binary.LittleEndian.PutUint16(p[pos+4+i:], u16)
		}
	}
	return nil
}

func (p Page) setValue(id uint16, value []byte) error {
	if id >= p.getNumKeys() {
		return fmt.Errorf("id must be less than number of keys (%d) in page", p.getNumKeys())
	}
	pos, err := p.kvPosition(id)
	if err != nil {
		return err
	}
	klen := binary.LittleEndian.Uint16(p[pos+0:])
	keyOffset := ceilToUint16(int(klen))

	oldValue, err := p.getValue(id)
	if err != nil {
		return err
	}

	vlen := len(value)
	oldVlen := len(oldValue)
	if ceilToUint16(vlen) > ceilToUint16(oldVlen) {
		return fmt.Errorf("value size (%d) must be less than or equal to %d bytes for replacement", vlen, ceilToUint16(oldVlen))
	}

	binary.LittleEndian.PutUint16(p[pos+2:], uint16(vlen))
	for i := uint16(0); i < uint16(vlen); i += 2 {
		if i == uint16(vlen)-1 {
			b := make([]byte, 2)
			b[0] = byte(value[i])
			u16 := binary.LittleEndian.Uint16(b)
			binary.LittleEndian.PutUint16(p[pos+4+keyOffset+i:], u16)
		} else {
			u16 := binary.LittleEndian.Uint16(value[i:])
			binary.LittleEndian.PutUint16(p[pos+4+keyOffset+i:], u16)
		}
	}
	return nil
}

func (p Page) insertKeyValue(id uint16, key []byte, value []byte) error {
	if id > p.getNumKeys() {
		return fmt.Errorf("id cannot be more than the current number of keys (%d) in page", p.getNumKeys())
	}

	offsetCorrection := kvSize(key, value)

	offsets := make([]uint16, p.getNumKeys()+1)
	keys := make([]string, p.getNumKeys()+1)
	values := make([]string, p.getNumKeys()+1)

	for i := uint16(0); i < p.getNumKeys(); i++ {
		baseOffset, err := p.getOffset(i)
		if err != nil {
			return err
		}
		keyi, err := p.getKey(i)
		if err != nil {
			return err
		}
		vali, err := p.getValue(i)
		if err != nil {
			return err
		}
		if i < id {
			offsets[i] = baseOffset
			keys[i] = string(keyi)
			values[i] = string(vali)
		}
		if i >= id {
			offsets[i+1] = baseOffset + offsetCorrection
			keys[i+1] = string(keyi)
			values[i+1] = string(vali)
		}
	}
	if id > 0 {
		offsets[id] = offsets[id-1] + offsetCorrection
	}
	keys[id] = string(key)
	values[id] = string(value)
	err := p.setNumKeys(uint16(len(offsets)))
	if err != nil {
		return err
	}
	for i := 0; i < len(offsets); i++ {
		p.setOffset(uint16(i), offsets[i])
		p.setKeyValue(uint16(i), []byte(keys[i]), []byte(values[i]))
	}
	return nil
}

func (p Page) setKeyValue(id uint16, key []byte, value []byte) error {
	if id >= p.getNumKeys() {
		return fmt.Errorf("id must be less than number of keys (%d) in page", p.getNumKeys())
	}
	pos, err := p.kvPosition(id)
	if err != nil {
		return err
	}
	klen := uint16(len(key))
	keyOffset := ceilToUint16(int(klen))
	binary.LittleEndian.PutUint16(p[pos:], uint16(klen))
	for i := uint16(0); i < klen; i += 2 {
		if i == klen-1 {
			b := make([]byte, 2)
			b[0] = byte(key[i])
			u16 := binary.LittleEndian.Uint16(b)
			binary.LittleEndian.PutUint16(p[pos+4+i:], u16)

		} else {
			u16 := binary.LittleEndian.Uint16(key[i:])
			binary.LittleEndian.PutUint16(p[pos+4+i:], u16)
		}
	}
	vlen := uint16(len(value))
	binary.LittleEndian.PutUint16(p[pos+2:], uint16(vlen))
	for i := uint16(0); i < vlen; i += 2 {
		if i == vlen-1 {
			b := make([]byte, 2)
			b[0] = byte(value[i])
			u16 := binary.LittleEndian.Uint16(b)
			binary.LittleEndian.PutUint16(p[pos+4+keyOffset+i:], u16)
		} else {
			u16 := binary.LittleEndian.Uint16(value[i:])
			binary.LittleEndian.PutUint16(p[pos+4+keyOffset+i:], u16)
		}
	}
	return nil
}

func (p Page) getKey(id uint16) ([]byte, error) {
	pos, err := p.kvPosition(id)
	if err != nil {
		return nil, err
	}
	klen := binary.LittleEndian.Uint16(p[pos:])
	return p[pos+4:][:klen], nil
}

func (p Page) getKeyPositionOrLess(key []byte) uint16 {
	nkeys := p.getNumKeys()
	var found uint16 = 0
	for i := uint16(0); i < nkeys; i++ {
		pageKey, _ := p.getKey(i)
		cmp := bytes.Compare(pageKey, key)
		if cmp <= 0 {
			found = i
		}
		if cmp >= 0 {
			break
		}
	}
	return found
}

func (p Page) getValue(id uint16) ([]byte, error) {
	pos, err := p.kvPosition(id)
	if err != nil {
		return nil, err
	}
	klen := binary.LittleEndian.Uint16(p[pos+0:])
	keyOffset := ceilToUint16(int(klen))
	vlen := binary.LittleEndian.Uint16(p[pos+2:])
	return p[pos+4+keyOffset:][:vlen], nil
}

// page size in bytes
func (p Page) size() uint16 {
	res, _ := p.kvPosition(p.getNumKeys())
	return res
}

func ceilToUint16(v int) uint16 {
	if v%2 == 0 {
		return uint16(v)
	}
	return uint16(v + 1)
}

func kvSize(key []byte, value []byte) uint16 {
	return uint16(2+2) + ceilToUint16(len(key)) + ceilToUint16(len(value))
}
