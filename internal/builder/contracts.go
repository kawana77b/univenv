package builder

type Builder[T any] interface {
	Build() (T, error)
}
