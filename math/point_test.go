package math

import "testing"

func TestPoint3f_Max(t *testing.T) {
	p1 := NewPoint3f(1.0, 0.5, 3.0)
	p2 := NewPoint3f(-10.0, 4.0, 3.0)
	p3 := p1.Max(p2)
	gt := NewPoint3f(1.0, 4.0, 3.0)

	if p3 != gt {
		t.Errorf("Expect %v, got %v", gt, p3)
	}
}

func TestPoint3f_Min(t *testing.T) {
	p1 := NewPoint3f(1.0, 0.5, 3.0)
	p2 := NewPoint3f(-10.0, 4.0, 3.0)
	p3 := p1.Min(p2)
	gt := NewPoint3f(-10.0, 0.5, 3.0)

	if p3 != gt {
		t.Errorf("Expect %v, got %v", gt, p3)
	}
}
