package system

import "runtime"

func GetRuntimeStack(track *[]byte) {
	*track = (*track)[0:cap(*track)]
	point := runtime.Stack(*track, true)
	for point > len(*track) {
		*track = make([]byte, point)
		runtime.Stack(*track, true)
	}
	*track = (*track)[0:point]
}
