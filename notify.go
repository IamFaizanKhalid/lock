package lock

func Notify() <-chan struct{} {
	ch := make(chan struct{})

	go listen(ch)

	return ch
}
