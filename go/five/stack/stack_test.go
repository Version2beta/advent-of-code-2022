package stack

import (
	"testing"
)

func TestPush(t *testing.T) {
	got := &Stack{}
	pushData := []rune{'c', 'b', 'a'}
	expect := &Stack{Items: []rune{'a', 'b', 'c'}}
	got.Push(pushData)
	for i := range expect.Items {
		if got.Items[i] != expect.Items[i] {
			t.Errorf("Expected %#v, got %#v in item %d", expect, got, i)
		}
	}
}

func TestPop(t *testing.T) {
	expect := &Stack{Items: []rune{'b', 'c'}}
	expPopped := []rune{'a'}
	got := &Stack{Items: []rune{'a', 'b', 'c'}}
	gotPopped := got.Pop(1)
	if gotPopped[0] != expPopped[0] {
		t.Errorf("Expected %v, got %v", expPopped, gotPopped)
	}
	for i := range expect.Items {
		if got.Items[i] != expect.Items[i] {
			t.Errorf("Expected %#v, got %#v in item %d", expect, got, i)
		}
	}
}

func TestPopPush(t *testing.T) {
	expectA, expectB := &Stack{Items: []rune{'b', 'c'}}, &Stack{Items: []rune{'a'}}
	gotA, gotB := &Stack{Items: []rune{'a', 'b', 'c'}}, &Stack{}
	gotA, gotB = PopPush(1, gotA, gotB)
	for i := range expectA.Items {
		if gotA.Items[i] != expectA.Items[i] {
			t.Errorf("Expected %#v, got %#v in item %d", expectA, gotA, i)
		}
	}
	for i := range expectB.Items {
		if gotB.Items[i] != expectB.Items[i] {
			t.Errorf("Expected %#v, got %#v in item %d", expectB, gotB, i)
		}
	}
}
