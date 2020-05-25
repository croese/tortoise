package tortoise

import (
	"math"
)

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
	return rad * 180.0 / math.Pi
}

func (t *Turtle) Degrees() float64 {
	return radiansToDegrees(t.state.angleInRadians)
}

func (t *Turtle) Left(degrees float64) *Turtle {
	targetAngle := degrees * math.Pi / 180

	t.state.angleInRadians += targetAngle
	return t
}

func (t *Turtle) Right(degrees float64) *Turtle {
	return t.Left(-degrees)
}

func (t *Turtle) Forward(distance float64) *Turtle {
	dx := math.Cos(t.state.angleInRadians)
	dy := math.Sin(t.state.angleInRadians)

	t.state.x += distance * dx
	t.state.y += distance * dy
	return t
}

func (t *Turtle) Back(distance float64) *Turtle {
	return t.Forward(-distance)
}
