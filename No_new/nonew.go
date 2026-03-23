package nonew

type S struct {
	N1, N2, N3, N4 int
}

func F(pn *S) {
	x := *pn
	_ = x // Use x
}

func UseF() {
	var v S
	// F(&v)
	F(new(v))
}

func New[T any](v T) *T {
	return &v
}

func F2(pn *S) {
	pn.N1 = 200
}
