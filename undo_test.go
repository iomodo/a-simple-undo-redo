package undo

import (
	"testing"
)

func TestSimpleUndo(t *testing.T) {
	undoer := NewUndoer(10)
	state := undoer.State()
	if state != nil {
		t.Errorf("state not initialized, should be nil, was %v", state)
	}

	undoer.Save(1)
	undoer.Save(2)
	state = undoer.State().(int)
	if state != 2 {
		t.Errorf("state should be 2, was %v", state)
	}

	undoer.Undo()
	state = undoer.State().(int)
	if state != 1 {
		t.Errorf("state should be 1, was %v", state)
	}

	undoer.Undo()
	state = undoer.State()
	if state != 1 {
		t.Errorf("state should be 1, was %v", state)
	}

	undoer.Redo()
	state = undoer.State()
	if state != 2 {
		t.Errorf("state should be 2, was %v", state)
	}

	undoer.Redo()
	state = undoer.State()
	if state != 2 {
		t.Errorf("state should be 2, was %v", state)
	}
}

func TestUnlimitedUndo(t *testing.T) {
	undoer := NewUndoer(0)
	num := 100
	for i := 0; i < num; i++ {
		undoer.Save(i)
		state := undoer.State()
		if state != i {
			t.Errorf("state should be %v, was %v", i, state)
		}
		for j := 0; j < i; j++ {
			undoer.Undo()
			state := undoer.State()
			if state != i-j-1 {
				t.Errorf("state should be %v, was %v", i-j-1, state)
			}
		}
		for j := 0; j < i; j++ {
			undoer.Redo()
			state := undoer.State()
			if state != j+1 {
				t.Errorf("state should be %v, was %v", i-j-1, state)
			}
		}
	}
}

func TestLimitedUndo(t *testing.T) {
	undoer := NewUndoer(2)
	undoer.Save(1)
	undoer.Save(2)
	undoer.Save(3)
	undoer.Save(4)

	undoer.Undo()
	undoer.Undo()

	state := undoer.State()
	if state != 2 {
		t.Errorf("state should be 2, was %v", state)
	}
	undoer.Undo()
	state = undoer.State()
	if state != 2 {
		t.Errorf("state should be 2, was %v", state)
	}

	undoer.Redo()
	state = undoer.State()
	if state != 3 {
		t.Errorf("state should be 3, was %v", state)
	}

	undoer.Save(5)
	undoer.Redo()
	state = undoer.State()
	if state != 5 {
		t.Errorf("state should be 5, was %v", state)
	}
}

func TestClearUndo(t *testing.T) {
	undoer := NewUndoer(2)
	undoer.Save(1)
	undoer.Save(2)
	undoer.Save(3)
	undoer.Save(4)
	undoer.Clear()

	state := undoer.State()
	if state != 4 {
		t.Errorf("state should be 4, was %v", state)
	}

	// Does nothing
	undoer.Undo()
	state = undoer.State()
	if state != 4 {
		t.Errorf("state should be 4, was %v", state)
	}

	//does nothing
	undoer.Redo()
	state = undoer.State()
	if state != 4 {
		t.Errorf("state should be 4, was %v", state)
	}
}
