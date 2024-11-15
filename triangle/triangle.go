package triangle

type Kind int

const (
    NaT = iota // not a triangle
    Equ // equilateral
    Iso // isosceles
    Sca // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	// validate if all sides are greater than 0
	if a <= 0 || b <= 0 || c <= 0 {
		return NaT
	}

	// validate if the sum of two sides is greater than the third side
	if a + b < c || b + c < a || a + c < b {
		return NaT
	}

	if a == b && b == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	} else {
		return Sca
	}
}
