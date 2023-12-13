package lists

type node[T any] struct {
	next  *node[T]
	value T
}
