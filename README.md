# a-simple-undo-redo
A simple undo/redo functionality in Go
## Interface
~~~Go
type State interface{}

type Undoer interface {
    State() State // returns the current state
    Save(newState State) // saves the new state onto the history
    Undo() // undos the state
    Redo() // redos the state
    Clear() // clears the history
}
~~~
## Example
~~~Go
// create undoer with unlimited capacity
undoer := NewUndoer(0)
undoer.Save(1)
undoer.Save(2)
undoer.Save(3)
undoer.Undo()
undoer.Undo()
undoer.State() // returns 1
undoer.Redo()
undoer.State() // returns 2
~~~
