package tortoise

import (
	"math"
	"testing"
)

func TestPlus(t *testing.T) {
	v := Vec2D{X: 1.5, Y: 3.0}
	w := Vec2D{X: 2.0, Y: 2.0}

	z := v.Plus(w)

	expected := Vec2D{X: 3.5, Y: 5.0}
	if z != expected {
		t.Fatalf("expected %+v, got %+v", expected, z)
	}
}

func TestTimes(t *testing.T) {
	v := Vec2D{X: 1.5, Y: 3.0}
	w := Vec2D{X: 2.0, Y: 2.0}

	product := v.Times(w)

	if product != 9.0 {
		t.Fatalf("expected 9.0, got %v", product)
	}
}

func TestScale(t *testing.T) {
	v := Vec2D{X: 1.5, Y: 3.0}

	w := v.Scale(10)

	expected := Vec2D{X: 15.0, Y: 30.0}
	if w != expected {
		t.Fatalf("expected %+v, got %+v", expected, w)
	}
}

func TestMinus(t *testing.T) {
	v := Vec2D{X: 1.5, Y: 3.0}
	w := Vec2D{X: 2.0, Y: 2.0}

	z := v.Minus(w)

	expected := Vec2D{X: -0.5, Y: 1.0}
	if expected != z {
		t.Fatalf("expected %+v, got %+v", expected, z)
	}
}

func TestNeg(t *testing.T) {
	v := Vec2D{X: 1.5, Y: 3.0}

	w := v.Neg()

	expected := Vec2D{X: -1.5, Y: -3.0}
	if expected != w {
		t.Fatalf("expected %+v, got %+v", expected, w)
	}
}

const epsilon = 0.00001

func equalVectors(a, b Vec2D) bool {
	return (math.Abs(a.X-b.X) <= epsilon) &&
		(math.Abs(a.Y-b.Y) <= epsilon)
}

func TestAbs(t *testing.T) {
	v := Vec2D{X: 3.0, Y: 4.0}

	abs := v.Abs()

	if abs != 5.0 {
		t.Fatalf("expected 5.0, got %v", abs)
	}
}

func TestRotate(t *testing.T) {
	v := Vec2D{X: 1.0, Y: 0.0}

	w := v.Rotate(90.0)

	expected := Vec2D{X: 0.0, Y: 1.0}
	if !equalVectors(w, expected) {
		t.Fatalf("expected %+v, got %+v", expected, w)
	}
}
