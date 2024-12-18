package util

// gcd calculates the greatest common divisor of a and b
func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

// extendedGCD returns the gcd of a and b and the coefficients x, y for the equation a*x + b*y = gcd(a, b)
func ExtendedGCD(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}
	g, x1, y1 := ExtendedGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1
	return g, x, y
}
