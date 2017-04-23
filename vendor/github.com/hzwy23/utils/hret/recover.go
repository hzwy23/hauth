package hret

type httpPanicFunc func()

// HttpPanic user for stop panic up.
func HttpPanic(f ...httpPanicFunc) {
	if r := recover(); r != nil {
		for _, val := range f {
			val()
		}
	}
}
