package tagmap

import (
	"math/bits"
)

//startSize the initial size for the TagMap
const startSize = 256

//TagMap a read only map with bool keys
type TagMap struct {
	len     int
	hasZero bool

	keys  []uint32
	mask  uint32
	chain uint32
}

//Has returns if v is in the tag map
func (t *TagMap) Has(v uint32) bool {
	if v == 0 {
		return t.hasZero
	}

	idx := hash32(v)

	var i uint32

	for ; i < t.chain; i++ {
		k := t.keys[(idx+i)&t.mask]
		if k == v || k == 0 {
			return k == v

		}
	}

	return false
}

//Len gets the length of the tag map
func (t *TagMap) Len() int { return t.len }

//Keys get all the keys
func (t *TagMap) Keys() []uint32 {
	var start = 0
	if t.hasZero {
		start = 1
	}
	v := make([]uint32, start, t.len)
	for _, k := range t.keys {
		if k != 0 {
			v = append(v, k)
		}
	}
	return v
}

//New creates a new TagMap
func New(m map[uint32]struct{}) *TagMap {
	length := uint32(1 << (bits.Len(uint(len(m))) + 1))
	t := TagMap{mask: length - 1, keys: make([]uint32, length), hasZero: true, len: len(m)}
	max := 0
	for k := range m {
		if k == 0 {
			continue
		}
		idx := hash32(k)
		for i := 0; i < len(m); i++ {
			h := t.keys[(idx+uint32(i))&t.mask]
			if h == 0 {
				t.keys[(idx+uint32(i))&t.mask] = k
				if i > max {
					max = i
				}
				break
			}
		}
	}
	t.chain = uint32(max)
	return &t
}
