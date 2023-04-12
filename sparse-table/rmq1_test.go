package sparsetable

import "testing"

func TestRMQpm1_Min(t *testing.T) {
	t.Parallel()
	rmq := NewRMQpm1([]int{1, 2, 1, 0, -1, -2, -1, 0, 1, 2, 1, 0, -1, -2, -3, -2, -3})
	if value := rmq.Min(0, 15); value != -3 {
		t.Errorf("wrong answer %d", value)
	}
	if value := rmq.Min(0, 15); value != -3 {
		t.Errorf("wrong answer %d", value)
	}
	if value := rmq.Min(0, 0); value != 1 {
		t.Errorf("wrong answer %d", value)
	}
	if value := rmq.Min(2, 3); value != 0 {
		t.Errorf("wrong answer %d", value)
	}
	if value := rmq.Min(7, 13); value != -2 {
		t.Errorf("wrong answer %d", value)
	}
	if value := rmq.Min(10, 14); value != -3 {
		t.Errorf("wrong answer %d", value)
	}
	if value := rmq.Min(2, 5); value != -2 {
		t.Errorf("wrong answer %d", value)
	}
}

func TestRMQpm1_Min2(t *testing.T) {
	t.Parallel()
	rmq := NewRMQpm1([]int{1})
	if value := rmq.Min(0, 0); value != 1 {
		t.Errorf("wrong answer %d", value)
	}
}
