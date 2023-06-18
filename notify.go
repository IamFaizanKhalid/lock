package lock

func Notify() <-chan Event {
	ch := make(chan Event)

	go routine(ch)

	return ch
}

func HandleEvents(h EventHandler) {
	for e := range Notify() {
		h(e)
	}
}
