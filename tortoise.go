package tortoise

import "math"

type turtleState struct {
  position Vec2D
  angleInRadians float64
}

type Turtle struct {
  state *turtleState
}

func New() *Turtle {
  return NewWithPosition(0,0)
}

func NewWithPosition(x,y float64) *Turtle {
  return &Turtle{
    state: &turtleState{
      position: Vec2D{
        X: x,
        Y: y,
      },
      angleInRadians: 0,
    },
  }
}

func (t *Turtle) Position() Vec2D {
  return t.state.position
}

func radiansToDegrees(rad float64) float64 {
  return rad * math.Pi / 180.0
}

func (t *Turtle) Degrees() float64 {
  return radiansToDegrees(t.state.angleInRadians)
}