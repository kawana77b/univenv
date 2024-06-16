package task

type Task[T any] struct {
	result T
	err    error
}

func NewTask[T any](fn func() (T, error)) *Task[T] {
	t := &Task[T]{}
	t.result, t.err = fn()
	return t
}

func (t *Task[T]) Then(fn func(T) (T, error)) *Task[T] {
	if t.err != nil {
		return t
	}
	t.result, t.err = fn(t.result)
	return t
}

func (t *Task[T]) Catch(fn func(error)) *Task[T] {
	if t.err != nil {
		fn(t.err)
	}
	return t
}

func (t *Task[T]) Result() (T, error) {
	return t.result, t.err
}
