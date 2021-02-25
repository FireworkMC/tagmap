package tagmap

//fnv32a constants for hashing
const (
	prime32 uint32 = 16777619
	offset  uint32 = 2166136261
)

//hash32 a fnv32a implimentation that can be inlined
func hash32(c uint32) (v uint32) {
	v = offset
	for i := 0; i < 32; i += 4 {
		v ^= (c >> i) & 0xff
		v *= prime32
	}
	return
}
