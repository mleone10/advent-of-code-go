package geo

type Location struct {
	A, B, C, D int
}

func LocSum(l, m Location) Location {
	return Location{A: l.A + m.A, B: l.B + m.B, C: l.C + m.C, D: l.D + m.D}
}
