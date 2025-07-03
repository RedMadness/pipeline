package pipeline

// Handler — это функция, принимающая и возвращающая значение одного и того же типа.
// Используется как строительный блок для цепочки обработки.
type Handler[T any] func(T) T

// Pipeline реализует паттерн "конвейера" (pipeline), позволяет последовательно выполнять функции по цепочке.
type Pipeline[T any] struct {
	steps []Handler[T]
}

// New создает новый экземпляр Pipeline для указанного типа T.
func New[T any]() *Pipeline[T] {
	return &Pipeline[T]{}
}

// Through добавляет один или несколько обработчиков (steps) в конвейер.
func (p *Pipeline[T]) Through(steps ...Handler[T]) *Pipeline[T] {
	p.steps = append(p.steps, steps...)
	return p
}

// Finalize завершает сборку конвейера, принимая финальный обработчик (final).
// Все шаги из Through будут обёрнуты вокруг final и возвращёна результирующая функция.
func (p *Pipeline[T]) Finalize(final T) T {
	next := final
	for i := len(p.steps) - 1; i >= 0; i-- {
		next = p.steps[i](next)
	}
	return next
}
