package keccak

var roundc = [24]uint64{
	0x0000000000000001,
	0x0000000000008082,
	0x800000000000808a,
	0x8000000080008000,
	0x000000000000808b,
	0x0000000080000001,
	0x8000000080008081,
	0x8000000000008009,
	0x000000000000008a,
	0x0000000000000088,
	0x0000000080008009,
	0x000000008000000a,
	0x000000008000808b,
	0x800000000000008b,
	0x8000000000008089,
	0x8000000000008003,
	0x8000000000008002,
	0x8000000000000080,
	0x000000000000800a,
	0x800000008000000a,
	0x8000000080008081,
	0x8000000000008080,
	0x0000000080000001,
	0x8000000080008008,
}
var rotc = [5][5]uint{
	{0x0, 0x1, 0x3e, 0x1c, 0x1b},
	{0x24, 0x2c, 0x6, 0x37, 0x14},
	{0x3, 0xa, 0x2b, 0x19, 0x27},
	{0x29, 0x2d, 0xf, 0x15, 0x8},
	{0x12, 0x2, 0x3d, 0x38, 0xe},
}