package tagmap

//fnv32a constants for hashing
const (
	prime32 uint32 = 16777619
	offset  uint32 = 2166136261
)

//hash32 a fnv32a implimentation that can be inlined
func hash32(c uint32) uint32 {
	return ((((offset^(c)&0xFF)*
		prime32^(c>>8)&0xFF)*
		prime32^(c>>16)&0xFF)*
		prime32 ^ (c >> 24)) *
		prime32
}
