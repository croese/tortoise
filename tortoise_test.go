package tortoise

import "testing"

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