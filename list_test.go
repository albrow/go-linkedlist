package linkedlist

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	l := New()
	{
		found := l.Add(0.0)
		compareBool(t, "Add(0.0)", false, found)
	}
	{
		found := l.Add(2.0)
		compareBool(t, "Add(2.0)", false, found)
	}
	{
		found := l.Add(1.0)
		compareBool(t, "Add(1.0)", false, found)
	}
	expected := []float64{0.0, 1.0, 2.0}
	compareFloatSlice(t, "List", expected, l.GetAll())
	compareInt(t, "List length", 3, l.Length)
	// Add 1.0 again. Since it is already in the set, it should
	// return true and not change the length of or the elements
	// the set.
	{
		found := l.Add(1.0)
		compareBool(t, "Add(1.0)", true, found)
	}
	compareFloatSlice(t, "List", expected, l.GetAll())
	compareInt(t, "List length", 3, l.Length)
}

func TestGetIndex(t *testing.T) {
	l := New()
	l.Add(0.0)
	l.Add(2.0)
	l.Add(1.0)
	{
		got := l.GetIndex(0.0)
		compareInt(t, "GetIndex(0.0)", 0, got)
	}
	{
		got := l.GetIndex(1.0)
		compareInt(t, "GetIndex(1.0)", 1, got)
	}
	{
		got := l.GetIndex(2.0)
		compareInt(t, "GetIndex(2.0)", 2, got)
	}
	{
		got := l.GetIndex(3.0)
		compareInt(t, "GetIndex(3.0)", -1, got)
	}
}

func TestGetAtIndex(t *testing.T) {
	l := New()
	l.Add(0.0)
	l.Add(2.0)
	l.Add(1.0)
	{
		got, found := l.GetAtIndex(0)
		compareFloat(t, "First return value of GetAtIndex(0)", 0.0, got)
		compareBool(t, "Second return value of GetAtIndex(0)", true, found)
	}
	{
		got, found := l.GetAtIndex(1)
		compareFloat(t, "First return value of GetAtIndex(1)", 1.0, got)
		compareBool(t, "Second return value of  GetAtIndex(1)", true, found)
	}
	{
		got, found := l.GetAtIndex(2)
		compareFloat(t, "First return value of GetAtIndex(2)", 2.0, got)
		compareBool(t, "Second return value of GetAtIndex(2)", true, found)
	}
	{
		_, found := l.GetAtIndex(3)
		compareBool(t, "Second return value of GetAtIndex(3)", false, found)
	}
}

func TestDel(t *testing.T) {
	l := New()
	l.Add(0.0)
	l.Add(2.0)
	l.Add(1.0)
	l.Add(3.0)
	ok := l.Del(2.0)
	compareBool(t, "Result of l.Del(2.0)", true, ok)
	expected := []float64{0.0, 1.0, 3.0}
	compareFloatSlice(t, "List", expected, l.GetAll())
	compareInt(t, "List length", 3, l.Length)
	ok = l.Del(5.0)
	compareBool(t, "Result of l.Del(5.0)", false, ok)
	compareInt(t, "List length", 3, l.Length)
}

func TestDelAtIndex(t *testing.T) {
	l := New()
	l.Add(0.0)
	l.Add(2.0)
	l.Add(1.0)
	l.Add(3.0)
	ok := l.DelAtIndex(2)
	compareBool(t, "Result of l.DelAtIndex(2)", true, ok)
	expected := []float64{0.0, 1.0, 3.0}
	compareFloatSlice(t, "List", expected, l.GetAll())
	compareInt(t, "List length", 3, l.Length)
	ok = l.DelAtIndex(4)
	compareBool(t, "Result of l.DelAtIndex(5)", false, ok)
	compareInt(t, "List length", 3, l.Length)
}

func TestGetAllRev(t *testing.T) {
	l := New()
	l.Add(0.0)
	l.Add(2.0)
	l.Add(1.0)
	expected := []float64{2.0, 1.0, 0.0}
	compareFloatSlice(t, "Reversed list", expected, l.GetAllRev())
}

func compareBool(t *testing.T, prefix string, expected bool, got bool) {
	if expected != got {
		t.Errorf("%s was wrong. Expected %t but got %t", prefix, expected, got)
	}
}

func compareInt(t *testing.T, prefix string, expected int, got int) {
	if expected != got {
		t.Errorf("%s was wrong. Expected %d but got %d.", prefix, expected, got)
	}
}

func compareFloat(t *testing.T, prefix string, expected float64, got float64) {
	if expected != got {
		t.Errorf("%s was wrong. Expected %0.1f but got %0.1f.", prefix, expected, got)
	}
}

func compareFloatSlice(t *testing.T, prefix string, expected []float64, got []float64) {
	if len(expected) != len(got) {
		t.Errorf("%s had wrong length. Expected %d but got %d", prefix, len(expected), len(got))
		t.Errorf("\nExpected: %v\nGot: %v\n", expected, got)
		return
	}
	for i, e := range expected {
		prefix += fmt.Sprintf(" value at index %d", i)
		compareFloat(t, prefix, e, got[i])
	}
}
