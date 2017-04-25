package hret

import "github.com/hzwy23/utils/logs"

type httpPanicFunc func()

// HttpPanic user for stop panic up.
func HttpPanic(f ...httpPanicFunc) {
	if r := recover(); r != nil {
		logs.Error("system generator panic.", r)
		for _, val := range f {
			val()
		}
	}
}
