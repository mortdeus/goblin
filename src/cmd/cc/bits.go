package cc

func bor(a, b Bits) Bits {
	var c Bits
	var i int

	c.b[i] = a.b[i] | b.b[i]
	return c
}

func band(a, b Bits) Bits {
	var c Bits
	var i int

	for i = 0; i < BITS; i++ {
		c.b[i] = a.b[i] & b.b[i]
	}
	return c
}

/*
func bnot(Bits a) Bits{
	Bits c;
	int i;

	for i=0; i<BITS; i++{
			c.b[i] = ~a.b[i];
			}
	return c;
}
*/

func bany(a *Bits) int {
	for i := 0; i < BITS; i++ {
		if a.b[i] {
			return 1
		}
	}
	return 0
}

func beq(a, b Bits) int {
	for i := 0; i < BITS; i++ {
		if a.b[i] != b.b[i] {
			return 0
		}
	}
	return 1
}

func bnum(a Bits) int {

	var b int32

	for i := 0; i < BITS; i++ {
		b = a.b[i]
		if b {
			return 32*i + bitno(b)
		}
	}
	diag(Z, "bad in bnum")
	return 0
}

func blsh(n uint) Bits {
	var c Bits

	c = zbits
	c.b[n/32] = 1 << (n % 32)
	return c
}

func bset(Bits a, uint n) int {
	if a.b[n/32] & (1 << (n % 32)) {
		return 1
	}
	return 0
}
