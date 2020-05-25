package tortoise

import (
	"math"
	"testing"
)

const epsilon = 0.00001

func equalFloats(a, b float64) bool {
	return math.Abs(a-b) <= epsilon
}

func TestNewWithDefaultPosition(t *testing.T) {
	turtle := New()

	p := turtle.Position()
	a := turtle.Degrees()

	if p.X != 0 {
		t.Fatalf("p.X: expected=0. got=%v",
			p.X)
	}

	if p.Y != 0 {
		t.Fatalf("p.Y: expected=0. got=%v",
			p.X)
	}

	if a != 0 {
		t.Fatalf("p.Degrees(): expected=0. got=%v",
			a)
	}
}

func TestNewWithSpecificPosition(t *testing.T) {
	turtle := NewWithPosition(12.5, 7.0)

	p := turtle.Position()
	a := turtle.Degrees()

	if p.X != 12.5 {
		t.Fatalf("p.X: expected=12.5. got=%v",
			p.X)
	}

	if p.Y != 7.0 {
		t.Fatalf("p.Y: expected=7.0. got=%v",
			p.X)
	}

	if a != 0 {
		t.Fatalf("p.Degrees(): expected=0. got=%v",
			a)
	}
}

func TestTurtle_Left(t *testing.T) {
	turtle := New()

	change := 90.0
	turtle.Left(change)
	actual := turtle.Degrees()

	if actual != change {
		t.Fatalf("turtle.Left: expected=%v. got=%v",
			change, actual)
	}

	turtle.Left(-45.0)
	actual = turtle.Degrees()

	if actual != 45.0 {
		t.Fatalf("turtle.Left: expected=%v. got=%v",
			45.0, actual)
	}
}

func TestTurtle_Right(t *testing.T) {
	turtle := New()

	change := 90.0
	turtle.Right(change)
	actual := turtle.Degrees()

	if actual != -change {
		t.Fatalf("turtle.Right: expected=%v. got=%v",
			change, actual)
	}

	turtle.Right(-45.0)
	actual = turtle.Degrees()

	if actual != -45.0 {
		t.Fatalf("turtle.Right: expected=%v. got=%v",
			45.0, actual)
	}
}

func TestTurtle_Forward(t *testing.T) {
	turtle := New()

	turtle.Forward(5.5)
	turtle.Left(90)
	turtle.Forward(10)
	p := turtle.Position()

	if !equalFloats(p.X, 5.5) {
		t.Fatalf("p.X: expected=5.5. got=%v",
			p.X)
	}

	if !equalFloats(p.Y, 10.0) {
		t.Fatalf("p.Y: expected=10.0. got=%v",
			p.X)
	}

	turtle.Forward(-3)
	p = turtle.Position()

	if !equalFloats(p.X, 5.5) {
		t.Fatalf("p.X: expected=5.5. got=%v",
			p.X)
	}

	if !equalFloats(p.Y, 7.0) {
		t.Fatalf("p.Y: expected=7.0. got=%v",
			p.X)
	}
}

func TestTurtle_Back(t *testing.T) {
	turtle := New()

	turtle.Forward(5.5)
	turtle.Back(2.0)
	p := turtle.Position()

	if !equalFloats(p.X, 3.5) {
		t.Fatalf("p.X: expected=3.5. got=%v",
			p.X)
	}

	if !equalFloats(p.Y, 0.0) {
		t.Fatalf("p.Y: expected=0.0. got=%v",
			p.X)
	}
}
