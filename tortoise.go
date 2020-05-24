package tortoise

import "math"

type Point2D struct {
	X, Y float64
}

type turtleState struct {
	x, y, angleInRadians float64
}

type Turtle struct {
	state *turtleState
}

func New() *Turtle {
	return NewWithPosition(0, 0)
}

func NewWithPosition(x, y float64) *Turtle {
	return &Turtle{
		state: &turtleState{
			x:              x,
			y:              y,
			angleInRadians: 0,
		},
	}
}

func (t *Turtle) Position() Point2D {
	return Point2D{
		X: t.state.x,
		Y: t.state.y,
	}
}

func radiansToDegrees(rad float64) float64 {
	return rad * math.Pi / 180.0
}

func (t *Turtle) Degrees() float64 {
	return radiansToDegrees(t.state.angleInRadians)
}
