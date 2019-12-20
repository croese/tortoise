package tortoise

import "math"

// Vec2D represents a two-dimensional vector
type Vec2D struct {
	X, Y float64
}

// Plus returns a new Vec2D that is the sum of the receiver and other
func (v Vec2D) Plus(other Vec2D) Vec2D {
	return Vec2D{X: v.X + other.X, Y: v.Y + other.Y}
}

// Times returns the inner product of the receiver and other
func (v Vec2D) Times(other Vec2D) float64 {
	return v.X*other.X + v.Y*other.Y
}

// Scale returns new vector, scaled by scalar
func (v Vec2D) Scale(scalar float64) Vec2D {
	return Vec2D{X: v.X * scalar, Y: v.Y * scalar}
}

// Minus returns a new Vec2D that is the
// difference of the receiver and other
func (v Vec2D) Minus(other Vec2D) Vec2D {
	return Vec2D{X: v.X - other.X, Y: v.Y - other.Y}
}

// Neg returns a negated copy of the receiver
func (v Vec2D) Neg() Vec2D {
	return Vec2D{X: -v.X, Y: -v.Y}
}

// Abs returns the absolute value of the receiver
func (v Vec2D) Abs() float64 {
	return math.Sqrt(math.Pow(v.X, 2.0) + math.Pow(v.Y, 2.0))
}

// Rotate rotates the receiver counterclockwise by angle
func (v Vec2D) Rotate(angle float64) Vec2D {
	perp := Vec2D{X: -v.Y, Y: v.X}
	angle = angle * math.Pi / 180.0
	c := math.Cos(angle)
	s := math.Sin(angle)
	return Vec2D{
		X: v.X*c + perp.X*s,
		Y: v.Y*c + perp.Y*s,
	}
}
