package main

import (
	"math"
	"testing"
)

func TestLength(t *testing.T) {
	v := Vec3{X: 3.0, Y: 4.0, Z: 5.0}
	if v.Length() != math.Sqrt(50) {
		t.Fatalf("Length should be return sqrt(50) for the vector (3, 4, 5)")
	}
}

func TestSquaredLength(t *testing.T) {
	v := Vec3{X: 3.0, Y: 4.0, Z: 5.0}
	if v.SquaredLength() != 50 {
		t.Fatalf("SquaredLength should be return 50 for the vector (3, 4, 5)")
	}
}

func TestAdd(t *testing.T) {
	v := Vec3{X: 1.5, Y: -0.2, Z: -1}
	w := Add(v, UnitVector)
	if w.X != 2.5 || w.Y != 0.8 || w.Z != 0 {
		t.Fatalf("(1.5, -0.2, -1) + (1, 1, 1) != (2.5, 0.8, 0)")
	}
}

func TestSub(t *testing.T) {
	v := Vec3{X: 0.5, Y: 1, Z: 1.7}
	w := Sub(v, UnitVector)
	if w.X != -0.5 || w.Y != 0 || w.Z != 0.7 {
		t.Fatalf("(0.5, 1, 1.7) - (1, 1, 1) != (-0.5, 0, 0.7)")
	}
}

func TestScale(t *testing.T) {
	v := Vec3{X: 1, Y: -1, Z: 0}
	w := v.Scale(-1.5)
	if w.X != -1.5 || w.Y != 1.5 || w.Z != 0 {
		t.Fatalf("-1.5 * (1, -1, 0) != (-1.5, 1.5, 0): (%v, %v, %v)", w.X, w.Y, w.Z)
	}
}
