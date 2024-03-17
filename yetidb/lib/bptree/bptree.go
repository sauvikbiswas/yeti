package bptree

type BpTree struct {
	// pointer (a nonzero page number)
	root uint64
	// callbacks for managing on-disk pages
	get func(uint64) Page // dereference a pointer
	new func(Page) uint64 // allocate a new page
	del func(uint64)      // deallocate a page
}

// func (t BpTree) insert(p Page, id uint16, key []byte, value []byte) (Page, error) {
// 	p = p.copy()
// 	err := p.insertKeyValue(id, key, value)
// 	return p, err
// }
