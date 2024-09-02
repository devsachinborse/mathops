package mathops

import "testing"

func TestAdd(t *testing.T) {
    result := Add(3, 4)
    expected := 7
    if result != expected {
        t.Errorf("Add(3, 4) = %d; want %d", result, expected)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(10, 4)
    expected := 6
    if result != expected {
        t.Errorf("Subtract(10, 4) = %d; want %d", result, expected)
    }
}
